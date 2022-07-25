package codec

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yddeng/astk/pkg/protocol"
	"reflect"
)

var spaceProto = map[string]*p{}

type p struct {
	cmd2Type map[uint16]reflect.Type
	type2Cmd map[reflect.Type]uint16
}

//根据名字注册实例(注意函数非线程安全，需要在初始化阶段完成所有消息的Register)
func (p *p) register(msg proto.Message, id uint16) {
	if _, ok := p.cmd2Type[id]; ok {
		panic(fmt.Sprintf("id %d id areadly register. ", id))
	}

	tt := reflect.TypeOf(msg)
	p.cmd2Type[id] = tt
	p.type2Cmd[tt] = id
}

func (p *p) marshal(o interface{}) (uint16, []byte, error) {
	tt := reflect.TypeOf(o)
	id, ok := p.type2Cmd[tt]
	if !ok {
		return 0, nil, fmt.Errorf("marshal type: %s undefined. ", reflect.TypeOf(o))
	}
	data, err := proto.Marshal(o.(proto.Message))
	return id, data, err
}

func (p *p) unmarshal(cmd uint16, buff []byte) (interface{}, error) {
	tt, ok := p.cmd2Type[cmd]
	if !ok {
		return nil, fmt.Errorf("unmarshal cmd: %d undefined. ", cmd)
	}

	//反序列化的结构
	msg := reflect.New(tt.Elem()).Interface()
	err := proto.Unmarshal(buff, msg.(proto.Message))
	return msg, err
}

//根据名字注册实例(注意函数非线程安全，需要在初始化阶段完成所有消息的Register)
func Register(namespace string, msg proto.Message, id uint16) {
	var ns *p
	var ok bool

	if ns, ok = spaceProto[namespace]; !ok {
		ns = &p{
			cmd2Type: map[uint16]reflect.Type{},
			type2Cmd: map[reflect.Type]uint16{},
		}
		spaceProto[namespace] = ns
	}

	ns.register(msg, id)
}

func Marshal(namespace string, o interface{}) (uint16, []byte, error) {
	var p *p
	var ok bool
	if p, ok = spaceProto[namespace]; !ok {
		return 0, nil, fmt.Errorf("invaild namespace:%s", namespace)
	}
	return p.marshal(o)
}

func Unmarshal(namespace string, id uint16, buff []byte) (interface{}, error) {
	var p *p
	var ok bool
	if p, ok = spaceProto[namespace]; !ok {
		return nil, fmt.Errorf("invaild namespace:%s", namespace)
	}

	return p.unmarshal(id, buff)
}

const (
	CmdHeartbeat  = 0
	CmdLogin      = 1
	CmdCmdExec    = 2
	CmdProcStart  = 3
	CmdProcSignal = 4
	CmdProcState  = 5
	CmdTailLog    = 6

	CmdNodeState = 101
)

func init() {
	Register("msg", &protocol.Heartbeat{}, CmdHeartbeat)
	Register("msg", &protocol.NodeState{}, CmdNodeState)

	Register("req", &protocol.LoginReq{}, CmdLogin)
	Register("resp", &protocol.LoginResp{}, CmdLogin)

	Register("req", &protocol.CmdExecReq{}, CmdCmdExec)
	Register("resp", &protocol.CmdExecResp{}, CmdCmdExec)

	Register("req", &protocol.ProcessExecReq{}, CmdProcStart)
	Register("resp", &protocol.ProcessExecResp{}, CmdProcStart)

	Register("req", &protocol.ProcessSignalReq{}, CmdProcSignal)
	Register("resp", &protocol.ProcessSignalResp{}, CmdProcSignal)

	Register("req", &protocol.ProcessStateReq{}, CmdProcState)
	Register("resp", &protocol.ProcessStateResp{}, CmdProcState)

	Register("req", &protocol.TailLogReq{}, CmdTailLog)
	Register("resp", &protocol.TailLogResp{}, CmdTailLog)
}
