package astke

import (
	"bytes"
	"fmt"
	psProc "github.com/shirou/gopsutil/process"
	"github.com/yddeng/astk/pkg/types"
	"github.com/yddeng/astk/pkg/util"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"
	"syscall"
	"time"
)

type Daemon struct {
	// 这个选项是进程启动多少秒之后，此时状态如果是running，则我们认为启动成功了
	// 默认值为1 。。非必须设置
	StartSecs int
	// 当进程启动失败后，最大尝试启动的次数。。当超过预定次数后，将把此进程的状态置为Exited
	// 默认值为3 。。非必须设置。
	StartRetries int
	// 这个是当我们向子进程发送stop信号后，到系统返回信息所等待的最大时间。
	// 超过这个时间会向该子进程发送一个强制kill的信号。
	StopWaitSecs int
	// 运行状态
	Status string
	// Cmd
	Cmd *exec.Cmd
}

type Process struct {
	Pid        int32            `json:"pid"`
	Stderr     string           `json:"stderr"`
	ID         int32            `json:"id"`
	Name       string           `json:"name"`
	Status     types.ProcStatus `json:"status"`
	CreateTime int64            `json:"createTime"`

	tailLog *tailLog
	process *psProc.Process
	mu      sync.Mutex
}

func (this *Process) waitCmd(cmd *exec.Cmd, callback func(process *Process)) {
	go func() {
		err := cmd.Wait()

		this.mu.Lock()
		if err != nil {
			if state, ok := err.(*exec.ExitError); ok {
				// !success
				if state.ProcessState.Exited() {
					// exit
					this.Status = types.ProcStatusExited
				} else {
					// signal 人为操作，视为正常停机
					this.Status = types.ProcStatusStopped
				}
			} else {
				// 异常退出
				this.Status = types.ProcStatusExited
			}
		} else {
			// err == nil && success
			this.Status = types.ProcStatusStopped
		}
		this.mu.Unlock()
		callback(this)
	}()
}

func (this *Process) waitChild(proc *os.Process, callback func(process *Process)) {
	go func() {
		state, err := proc.Wait()
		this.mu.Lock()
		if err != nil {
			// 异常退出
			this.Status = types.ProcStatusExited
		} else {
			if !state.Success() {
				if state.Exited() {
					// exit
					this.Status = types.ProcStatusExited
				} else {
					// signal 人为操作，视为正常停机
					this.Status = types.ProcStatusStopped
				}
			} else {
				// success code=0
				this.Status = types.ProcStatusStopped
			}

		}
		this.mu.Unlock()
		callback(this)

	}()
}

func (this *Process) waitNoChild(callback func(process *Process)) {
	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		for range ticker.C {
			if err := syscall.Kill(int(this.Pid), 0); err != nil {
				if this.Stderr != "" {
					data, err := ioutil.ReadFile(this.Stderr)
					this.mu.Lock()
					if err == nil && len(data) != 0 {
						this.Status = types.ProcStatusExited
					} else {
						this.Status = types.ProcStatusStopped
					}
					this.mu.Unlock()
				} else {
					this.mu.Lock()
					this.Status = types.ProcStatusStopped
					this.mu.Unlock()
				}
				callback(this)
				ticker.Stop()
				return
			}
		}
	}()
}

func (this *Process) wait(callback func(process *Process)) error {
	pp, err := this.process.Parent()
	if err != nil {
		return err
	} else {
		if pp.Pid == int32(os.Getpid()) {
			ppp, err := os.FindProcess(int(this.Pid))
			if err != nil {
				return err
			}
			this.waitChild(ppp, callback)
		} else {
			this.waitNoChild(callback)
		}
	}
	return nil
}

func (this *Process) GetStatus() types.ProcStatus {
	this.mu.Lock()
	defer this.mu.Unlock()
	return this.Status
}

func (this *Process) CPUPercent() float64 {
	if this.process == nil {
		return 0
	}
	percent, err := this.process.Percent(0)
	if err != nil {
		return 0
	}
	return percent
}

func (this *Process) MemoryPercent() float32 {
	if this.process == nil {
		return 0
	}
	percent, err := this.process.MemoryPercent()
	if err != nil {
		return 0
	}
	return percent
}

func NewProcess(pid int32) (*Process, error) {
	p, err := psProc.NewProcess(pid)
	if err != nil {
		return nil, err
	}
	createTime, err := p.CreateTime()
	if err != nil {
		return nil, err
	}

	this := &Process{
		CreateTime: createTime,
		Status:     types.ProcStatusRunning,
		Pid:        pid,
		process:    p,
	}
	return this, nil
}

const tailLogLine = 20

type tailLog struct {
	logs         [][]byte // line -> 仅纪录20行数据
	buff         *bytes.Buffer
	head, length int32
	start, end   int32
}

func newTailLog() *tailLog {
	return &tailLog{
		buff:   &bytes.Buffer{},
		logs:   make([][]byte, 0, tailLogLine),
		length: 0,
		start:  0,
		end:    0,
	}
}

func (this *tailLog) Write(p []byte) (n int, err error) {
	n, err = this.buff.Write(p)
	if err != nil {
		return
	}
	for {
		line, err := this.buff.ReadBytes('\n')
		if err != nil {
			break
		}

		if this.length < tailLogLine {
			this.logs = append(this.logs, line)
			this.length++
			this.end++
		} else {
			this.logs[this.head] = line
			this.head = (this.head + 1) % tailLogLine
			this.start++
			this.end++
		}
	}
	return
}

func (this *tailLog) Read(start int32) ([]byte, int32) {
	if this.length == 0 || start >= this.end {
		return nil, this.end
	}

	buff := &bytes.Buffer{}
	n := this.length
	if start > this.start {
		n = this.end - start
	} else if start < this.start {
		// 前面日志已经丢失
		buff.Write([]byte(fmt.Sprintf("------------------\n 已丢弃%d行日志\n------------------\n", this.start-start)))
	}
	// fmt.Println(this.length, this.start, this.end, this.head, n)

	if n-this.head > 0 {
		buff.Write(bytes.Join(this.logs[this.length-n+this.head:], nil))
		buff.Write(bytes.Join(this.logs[:this.head], nil))
	} else {
		buff.Write(bytes.Join(this.logs[this.head-n:this.head], nil))
	}

	return buff.Bytes(), this.end
}

var (
	waitProcess = map[int32]*Process{}
	processFile string
)

func loadProcess(dataPath string) {
	var processMap map[int32]*Process
	processFile = path.Join(dataPath, "exec_info.json")
	if err := util.DecodeJsonFromFile(&processMap, processFile); err == nil {
		for _, p := range processMap {
			if p.GetStatus() != types.ProcStatusRunning {
				waitProcess[p.ID] = p
				continue
			}

			if proc, err := NewProcess(p.Pid); err != nil {
				log.Printf("loadProcess %s faield %d %v", p.Name, p.Pid, err)
			} else {
				if p.CreateTime != proc.CreateTime {
					log.Printf("loadProcess %s faield %d create time not equal", p.Name, p.Pid)
				} else {
					proc.ID = p.ID
					proc.Name = p.Name
					proc.Stderr = p.Stderr
					waitProcess[proc.ID] = proc
					proc.waitNoChild(func(process *Process) {
						er.Submit(func() {
							if process.GetStatus() == types.ProcStatusStopped {
								delete(waitProcess, process.ID)
							}
							saveProcess()
						})
					})
				}
			}
		}
	}
	saveProcess()
}

func saveProcess() {
	if err := util.EncodeJsonToFile(waitProcess, processFile); err != nil {
		log.Printf("saveProcess faield %v", err)
	}
}
