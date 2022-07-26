package astke

import (
	"fmt"
	"os/exec"
	"sync"
	"testing"
)

func TestNewProcess(t *testing.T) {

	cmd := exec.Command("./test/test")
	if err := cmd.Start(); err != nil {
		t.Error(err)
		return
	}

	wg := sync.WaitGroup{}

	t.Log("process start", cmd.Process.Pid)
	p, err := NewProcess(int32(cmd.Process.Pid))
	if err != nil {
		t.Error(err)
		return
	}

	wg.Add(1)
	p.waitCmd(cmd, func(process *Process) {
		t.Log(process.Pid, process.GetState())
		wg.Done()
	})

	wg.Wait()
}

func TestProcessWait(t *testing.T) {

	p, err := NewProcess(96543)
	if err != nil {
		t.Error(err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	p.waitNoChild(func(process *Process) {
		t.Log(process.Pid, process.GetState())
		wg.Done()
	})

	wg.Wait()
}

func TestTailLog(t *testing.T) {
	tl := newTailLog()

	for i := int32(1); i <= 16; i++ {
		tl.Write([]byte(fmt.Sprintf("line %d\n", i)))
		ctx, end := tl.Read(0)
		fmt.Println(string(ctx), end)
	}

	for i := int32(1); i <= 16; i++ {
		ctx, end := tl.Read(i)
		fmt.Println(string(ctx), end)
	}
}

func TestTest(t *testing.T) {
	type Test struct {
		name string
		age  int
	}

	yddeng := Test{
		name: "yddeng",
		age:  25,
	}
	fmt.Println(yddeng)

	yddeng = Test{
		age: 26,
	}
	fmt.Println(yddeng)
}
