// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LoginReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Inet                 string   `protobuf:"bytes,2,opt,name=inet,proto3" json:"inet,omitempty"`
	Net                  string   `protobuf:"bytes,3,opt,name=net,proto3" json:"net,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginReq) GetInet() string {
	if m != nil {
		return m.Inet
	}
	return ""
}

func (m *LoginReq) GetNet() string {
	if m != nil {
		return m.Net
	}
	return ""
}

func (m *LoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LoginResp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResp) Reset()         { *m = LoginResp{} }
func (m *LoginResp) String() string { return proto.CompactTextString(m) }
func (*LoginResp) ProtoMessage()    {}
func (*LoginResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *LoginResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResp.Unmarshal(m, b)
}
func (m *LoginResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResp.Marshal(b, m, deterministic)
}
func (m *LoginResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResp.Merge(m, src)
}
func (m *LoginResp) XXX_Size() int {
	return xxx_messageInfo_LoginResp.Size(m)
}
func (m *LoginResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResp proto.InternalMessageInfo

func (m *LoginResp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type CmdExecReq struct {
	Dir                  string   `protobuf:"bytes,1,opt,name=dir,proto3" json:"dir,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Args                 []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	Timeout              int32    `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdExecReq) Reset()         { *m = CmdExecReq{} }
func (m *CmdExecReq) String() string { return proto.CompactTextString(m) }
func (*CmdExecReq) ProtoMessage()    {}
func (*CmdExecReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *CmdExecReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdExecReq.Unmarshal(m, b)
}
func (m *CmdExecReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdExecReq.Marshal(b, m, deterministic)
}
func (m *CmdExecReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdExecReq.Merge(m, src)
}
func (m *CmdExecReq) XXX_Size() int {
	return xxx_messageInfo_CmdExecReq.Size(m)
}
func (m *CmdExecReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdExecReq.DiscardUnknown(m)
}

var xxx_messageInfo_CmdExecReq proto.InternalMessageInfo

func (m *CmdExecReq) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func (m *CmdExecReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CmdExecReq) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *CmdExecReq) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type CmdExecResp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	OutStr               string   `protobuf:"bytes,2,opt,name=outStr,proto3" json:"outStr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdExecResp) Reset()         { *m = CmdExecResp{} }
func (m *CmdExecResp) String() string { return proto.CompactTextString(m) }
func (*CmdExecResp) ProtoMessage()    {}
func (*CmdExecResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *CmdExecResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdExecResp.Unmarshal(m, b)
}
func (m *CmdExecResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdExecResp.Marshal(b, m, deterministic)
}
func (m *CmdExecResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdExecResp.Merge(m, src)
}
func (m *CmdExecResp) XXX_Size() int {
	return xxx_messageInfo_CmdExecResp.Size(m)
}
func (m *CmdExecResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdExecResp.DiscardUnknown(m)
}

var xxx_messageInfo_CmdExecResp proto.InternalMessageInfo

func (m *CmdExecResp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CmdExecResp) GetOutStr() string {
	if m != nil {
		return m.OutStr
	}
	return ""
}

type ProcessExecReq struct {
	Config               map[string]string `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Name                 string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Args                 []string          `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	Dir                  string            `protobuf:"bytes,4,opt,name=dir,proto3" json:"dir,omitempty"`
	Id                   int32             `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProcessExecReq) Reset()         { *m = ProcessExecReq{} }
func (m *ProcessExecReq) String() string { return proto.CompactTextString(m) }
func (*ProcessExecReq) ProtoMessage()    {}
func (*ProcessExecReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *ProcessExecReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessExecReq.Unmarshal(m, b)
}
func (m *ProcessExecReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessExecReq.Marshal(b, m, deterministic)
}
func (m *ProcessExecReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessExecReq.Merge(m, src)
}
func (m *ProcessExecReq) XXX_Size() int {
	return xxx_messageInfo_ProcessExecReq.Size(m)
}
func (m *ProcessExecReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessExecReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessExecReq proto.InternalMessageInfo

func (m *ProcessExecReq) GetConfig() map[string]string {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ProcessExecReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProcessExecReq) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ProcessExecReq) GetDir() string {
	if m != nil {
		return m.Dir
	}
	return ""
}

func (m *ProcessExecReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ProcessExecResp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Pid                  int32    `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessExecResp) Reset()         { *m = ProcessExecResp{} }
func (m *ProcessExecResp) String() string { return proto.CompactTextString(m) }
func (*ProcessExecResp) ProtoMessage()    {}
func (*ProcessExecResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *ProcessExecResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessExecResp.Unmarshal(m, b)
}
func (m *ProcessExecResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessExecResp.Marshal(b, m, deterministic)
}
func (m *ProcessExecResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessExecResp.Merge(m, src)
}
func (m *ProcessExecResp) XXX_Size() int {
	return xxx_messageInfo_ProcessExecResp.Size(m)
}
func (m *ProcessExecResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessExecResp.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessExecResp proto.InternalMessageInfo

func (m *ProcessExecResp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *ProcessExecResp) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

type ProcessSignalReq struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Signal               int32    `protobuf:"varint,2,opt,name=signal,proto3" json:"signal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessSignalReq) Reset()         { *m = ProcessSignalReq{} }
func (m *ProcessSignalReq) String() string { return proto.CompactTextString(m) }
func (*ProcessSignalReq) ProtoMessage()    {}
func (*ProcessSignalReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *ProcessSignalReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessSignalReq.Unmarshal(m, b)
}
func (m *ProcessSignalReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessSignalReq.Marshal(b, m, deterministic)
}
func (m *ProcessSignalReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessSignalReq.Merge(m, src)
}
func (m *ProcessSignalReq) XXX_Size() int {
	return xxx_messageInfo_ProcessSignalReq.Size(m)
}
func (m *ProcessSignalReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessSignalReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessSignalReq proto.InternalMessageInfo

func (m *ProcessSignalReq) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ProcessSignalReq) GetSignal() int32 {
	if m != nil {
		return m.Signal
	}
	return 0
}

type ProcessSignalResp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessSignalResp) Reset()         { *m = ProcessSignalResp{} }
func (m *ProcessSignalResp) String() string { return proto.CompactTextString(m) }
func (*ProcessSignalResp) ProtoMessage()    {}
func (*ProcessSignalResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{7}
}

func (m *ProcessSignalResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessSignalResp.Unmarshal(m, b)
}
func (m *ProcessSignalResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessSignalResp.Marshal(b, m, deterministic)
}
func (m *ProcessSignalResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessSignalResp.Merge(m, src)
}
func (m *ProcessSignalResp) XXX_Size() int {
	return xxx_messageInfo_ProcessSignalResp.Size(m)
}
func (m *ProcessSignalResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessSignalResp.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessSignalResp proto.InternalMessageInfo

func (m *ProcessSignalResp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ProcessStateReq struct {
	Ids                  []int32  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessStateReq) Reset()         { *m = ProcessStateReq{} }
func (m *ProcessStateReq) String() string { return proto.CompactTextString(m) }
func (*ProcessStateReq) ProtoMessage()    {}
func (*ProcessStateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{8}
}

func (m *ProcessStateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessStateReq.Unmarshal(m, b)
}
func (m *ProcessStateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessStateReq.Marshal(b, m, deterministic)
}
func (m *ProcessStateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessStateReq.Merge(m, src)
}
func (m *ProcessStateReq) XXX_Size() int {
	return xxx_messageInfo_ProcessStateReq.Size(m)
}
func (m *ProcessStateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessStateReq.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessStateReq proto.InternalMessageInfo

func (m *ProcessStateReq) GetIds() []int32 {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ProcessStateResp struct {
	States               map[int32]*ProcessState `protobuf:"bytes,1,rep,name=states,proto3" json:"states,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ProcessStateResp) Reset()         { *m = ProcessStateResp{} }
func (m *ProcessStateResp) String() string { return proto.CompactTextString(m) }
func (*ProcessStateResp) ProtoMessage()    {}
func (*ProcessStateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{9}
}

func (m *ProcessStateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessStateResp.Unmarshal(m, b)
}
func (m *ProcessStateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessStateResp.Marshal(b, m, deterministic)
}
func (m *ProcessStateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessStateResp.Merge(m, src)
}
func (m *ProcessStateResp) XXX_Size() int {
	return xxx_messageInfo_ProcessStateResp.Size(m)
}
func (m *ProcessStateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessStateResp.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessStateResp proto.InternalMessageInfo

func (m *ProcessStateResp) GetStates() map[int32]*ProcessState {
	if m != nil {
		return m.States
	}
	return nil
}

type ProcessState struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	ExitMsg              string   `protobuf:"bytes,3,opt,name=exitMsg,proto3" json:"exitMsg,omitempty"`
	Cpu                  float64  `protobuf:"fixed64,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Mem                  float64  `protobuf:"fixed64,5,opt,name=mem,proto3" json:"mem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessState) Reset()         { *m = ProcessState{} }
func (m *ProcessState) String() string { return proto.CompactTextString(m) }
func (*ProcessState) ProtoMessage()    {}
func (*ProcessState) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{10}
}

func (m *ProcessState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessState.Unmarshal(m, b)
}
func (m *ProcessState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessState.Marshal(b, m, deterministic)
}
func (m *ProcessState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessState.Merge(m, src)
}
func (m *ProcessState) XXX_Size() int {
	return xxx_messageInfo_ProcessState.Size(m)
}
func (m *ProcessState) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessState.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessState proto.InternalMessageInfo

func (m *ProcessState) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ProcessState) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ProcessState) GetExitMsg() string {
	if m != nil {
		return m.ExitMsg
	}
	return ""
}

func (m *ProcessState) GetCpu() float64 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *ProcessState) GetMem() float64 {
	if m != nil {
		return m.Mem
	}
	return 0
}

type TailLogReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Start                int32    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TailLogReq) Reset()         { *m = TailLogReq{} }
func (m *TailLogReq) String() string { return proto.CompactTextString(m) }
func (*TailLogReq) ProtoMessage()    {}
func (*TailLogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{11}
}

func (m *TailLogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TailLogReq.Unmarshal(m, b)
}
func (m *TailLogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TailLogReq.Marshal(b, m, deterministic)
}
func (m *TailLogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TailLogReq.Merge(m, src)
}
func (m *TailLogReq) XXX_Size() int {
	return xxx_messageInfo_TailLogReq.Size(m)
}
func (m *TailLogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TailLogReq.DiscardUnknown(m)
}

var xxx_messageInfo_TailLogReq proto.InternalMessageInfo

func (m *TailLogReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TailLogReq) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

type TailLogResp struct {
	Context              []byte   `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	End                  int32    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TailLogResp) Reset()         { *m = TailLogResp{} }
func (m *TailLogResp) String() string { return proto.CompactTextString(m) }
func (*TailLogResp) ProtoMessage()    {}
func (*TailLogResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{12}
}

func (m *TailLogResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TailLogResp.Unmarshal(m, b)
}
func (m *TailLogResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TailLogResp.Marshal(b, m, deterministic)
}
func (m *TailLogResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TailLogResp.Merge(m, src)
}
func (m *TailLogResp) XXX_Size() int {
	return xxx_messageInfo_TailLogResp.Size(m)
}
func (m *TailLogResp) XXX_DiscardUnknown() {
	xxx_messageInfo_TailLogResp.DiscardUnknown(m)
}

var xxx_messageInfo_TailLogResp proto.InternalMessageInfo

func (m *TailLogResp) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *TailLogResp) GetEnd() int32 {
	if m != nil {
		return m.End
	}
	return 0
}

type Heartbeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{13}
}

func (m *Heartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heartbeat.Unmarshal(m, b)
}
func (m *Heartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heartbeat.Marshal(b, m, deterministic)
}
func (m *Heartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeat.Merge(m, src)
}
func (m *Heartbeat) XXX_Size() int {
	return xxx_messageInfo_Heartbeat.Size(m)
}
func (m *Heartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeat proto.InternalMessageInfo

// 上报物理机状态
type NodeState struct {
	Cpu                  map[string]string `protobuf:"bytes,1,rep,name=cpu,proto3" json:"cpu,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Mem                  map[string]string `protobuf:"bytes,2,rep,name=mem,proto3" json:"mem,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Disk                 map[string]string `protobuf:"bytes,3,rep,name=disk,proto3" json:"disk,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Host                 map[string]string `protobuf:"bytes,4,rep,name=host,proto3" json:"host,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Net                  map[string]string `protobuf:"bytes,5,rep,name=net,proto3" json:"net,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *NodeState) Reset()         { *m = NodeState{} }
func (m *NodeState) String() string { return proto.CompactTextString(m) }
func (*NodeState) ProtoMessage()    {}
func (*NodeState) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{14}
}

func (m *NodeState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeState.Unmarshal(m, b)
}
func (m *NodeState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeState.Marshal(b, m, deterministic)
}
func (m *NodeState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeState.Merge(m, src)
}
func (m *NodeState) XXX_Size() int {
	return xxx_messageInfo_NodeState.Size(m)
}
func (m *NodeState) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeState.DiscardUnknown(m)
}

var xxx_messageInfo_NodeState proto.InternalMessageInfo

func (m *NodeState) GetCpu() map[string]string {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *NodeState) GetMem() map[string]string {
	if m != nil {
		return m.Mem
	}
	return nil
}

func (m *NodeState) GetDisk() map[string]string {
	if m != nil {
		return m.Disk
	}
	return nil
}

func (m *NodeState) GetHost() map[string]string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *NodeState) GetNet() map[string]string {
	if m != nil {
		return m.Net
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "loginReq")
	proto.RegisterType((*LoginResp)(nil), "loginResp")
	proto.RegisterType((*CmdExecReq)(nil), "cmdExecReq")
	proto.RegisterType((*CmdExecResp)(nil), "cmdExecResp")
	proto.RegisterType((*ProcessExecReq)(nil), "processExecReq")
	proto.RegisterMapType((map[string]string)(nil), "processExecReq.ConfigEntry")
	proto.RegisterType((*ProcessExecResp)(nil), "processExecResp")
	proto.RegisterType((*ProcessSignalReq)(nil), "processSignalReq")
	proto.RegisterType((*ProcessSignalResp)(nil), "processSignalResp")
	proto.RegisterType((*ProcessStateReq)(nil), "processStateReq")
	proto.RegisterType((*ProcessStateResp)(nil), "processStateResp")
	proto.RegisterMapType((map[int32]*ProcessState)(nil), "processStateResp.StatesEntry")
	proto.RegisterType((*ProcessState)(nil), "processState")
	proto.RegisterType((*TailLogReq)(nil), "tailLogReq")
	proto.RegisterType((*TailLogResp)(nil), "tailLogResp")
	proto.RegisterType((*Heartbeat)(nil), "heartbeat")
	proto.RegisterType((*NodeState)(nil), "nodeState")
	proto.RegisterMapType((map[string]string)(nil), "nodeState.CpuEntry")
	proto.RegisterMapType((map[string]string)(nil), "nodeState.DiskEntry")
	proto.RegisterMapType((map[string]string)(nil), "nodeState.HostEntry")
	proto.RegisterMapType((map[string]string)(nil), "nodeState.MemEntry")
	proto.RegisterMapType((map[string]string)(nil), "nodeState.NetEntry")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0xed, 0x38, 0x7f, 0x33, 0x6e, 0xfb, 0x07, 0x53, 0x21, 0xab, 0x08, 0x51, 0x6d, 0x55,
	0xd1, 0x93, 0x0f, 0xa9, 0xa0, 0x14, 0x71, 0xa2, 0x54, 0xea, 0x81, 0x72, 0x70, 0x6f, 0x3d, 0xe1,
	0xda, 0x83, 0xbb, 0x4a, 0xec, 0x35, 0xde, 0x0d, 0x4a, 0x9e, 0x83, 0xb7, 0xe2, 0x61, 0x78, 0x06,
	0x34, 0xbb, 0x6b, 0xd7, 0x49, 0x5b, 0x89, 0xdc, 0x66, 0x66, 0xbf, 0xd9, 0xf9, 0xe6, 0x9b, 0x59,
	0x1b, 0x76, 0x4a, 0x94, 0x32, 0x2d, 0x30, 0xae, 0x1b, 0xa1, 0x04, 0xbb, 0x81, 0xad, 0x99, 0x28,
	0x78, 0x95, 0xe0, 0x8f, 0x30, 0x84, 0x41, 0x95, 0x96, 0x18, 0x39, 0x07, 0xce, 0xf1, 0x28, 0xd1,
	0x36, 0xc5, 0x78, 0x85, 0x2a, 0x72, 0x4d, 0x8c, 0xec, 0x70, 0x0c, 0x1e, 0x85, 0x3c, 0x1d, 0x22,
	0x33, 0xdc, 0x03, 0x5f, 0x89, 0x29, 0x56, 0xd1, 0x40, 0xc7, 0x8c, 0xc3, 0x5e, 0xc3, 0xc8, 0xde,
	0x2d, 0x6b, 0xba, 0x28, 0x13, 0x79, 0x77, 0x39, 0xd9, 0xec, 0x1b, 0x40, 0x56, 0xe6, 0x17, 0x0b,
	0xcc, 0xa8, 0xfc, 0x18, 0xbc, 0x9c, 0x37, 0x16, 0x40, 0x66, 0x47, 0xc8, 0x5d, 0x25, 0x94, 0x36,
	0x85, 0x8c, 0xbc, 0x03, 0x8f, 0x62, 0x64, 0x87, 0x11, 0xfc, 0xa7, 0x78, 0x89, 0x62, 0xae, 0x34,
	0x01, 0x3f, 0x69, 0x5d, 0x76, 0x06, 0x41, 0x57, 0xe1, 0x71, 0x12, 0xe1, 0x0b, 0x18, 0x8a, 0xb9,
	0xba, 0x56, 0x8d, 0x2d, 0x63, 0x3d, 0xf6, 0xdb, 0x81, 0xdd, 0xba, 0x11, 0x19, 0x4a, 0xd9, 0x32,
	0x3c, 0x81, 0x61, 0x26, 0xaa, 0xef, 0xbc, 0x88, 0x9c, 0x03, 0xef, 0x38, 0x98, 0xbc, 0x8c, 0x57,
	0x01, 0xf1, 0xb9, 0x3e, 0xbd, 0xa8, 0x54, 0xb3, 0x4c, 0x2c, 0xf4, 0x9f, 0x9b, 0xb0, 0xed, 0x0f,
	0xee, 0xdb, 0xdf, 0x05, 0x97, 0xe7, 0x91, 0xaf, 0x3b, 0x72, 0x79, 0xbe, 0x7f, 0x06, 0x41, 0xaf,
	0x00, 0x25, 0x4c, 0x71, 0xd9, 0xea, 0x35, 0xc5, 0x25, 0x8d, 0xe1, 0x67, 0x3a, 0x9b, 0xb7, 0xb5,
	0x8c, 0xf3, 0xc1, 0x7d, 0xef, 0xb0, 0x53, 0xf8, 0x7f, 0x85, 0xea, 0x13, 0x5a, 0x8c, 0xc1, 0xab,
	0x79, 0xae, 0xd3, 0xfd, 0x84, 0x4c, 0xf6, 0x11, 0xc6, 0x36, 0xf1, 0x9a, 0x17, 0x55, 0x3a, 0xb3,
	0x83, 0x22, 0x94, 0xd3, 0xa1, 0x48, 0x43, 0xa9, 0x8f, 0x6d, 0xaa, 0xf5, 0xd8, 0x1b, 0x78, 0xb6,
	0x96, 0xfd, 0xc4, 0x26, 0x1c, 0x76, 0xfc, 0xae, 0x55, 0xaa, 0xd0, 0x56, 0xe1, 0xb9, 0xd4, 0x4a,
	0xfb, 0x09, 0x99, 0xec, 0x97, 0x73, 0x4f, 0xc6, 0xa0, 0x64, 0x1d, 0xbe, 0x85, 0xa1, 0x24, 0x47,
	0xda, 0x99, 0xbc, 0x8a, 0xd7, 0x21, 0xb1, 0xb6, 0xa4, 0x9d, 0x8a, 0x01, 0xef, 0x5f, 0x42, 0xd0,
	0x0b, 0xf7, 0xb5, 0xf4, 0x8d, 0x96, 0x87, 0x7d, 0x2d, 0x83, 0xc9, 0xce, 0xea, 0xb5, 0x3d, 0x69,
	0x15, 0x6c, 0xf7, 0x8f, 0x1e, 0x51, 0x67, 0x0f, 0x7c, 0x5d, 0xb5, 0x1d, 0x8b, 0x76, 0x68, 0x69,
	0x71, 0xc1, 0xd5, 0x95, 0x2c, 0xec, 0x4b, 0x6a, 0x5d, 0xba, 0x21, 0xab, 0xe7, 0x7a, 0x13, 0x9c,
	0x84, 0x4c, 0x8a, 0x94, 0x58, 0xea, 0x55, 0x70, 0x12, 0x32, 0xd9, 0x04, 0x40, 0xa5, 0x7c, 0xf6,
	0x45, 0x14, 0xa4, 0x95, 0xd9, 0x14, 0xa7, 0xdd, 0x14, 0x5b, 0xb1, 0x51, 0x76, 0x1c, 0xc6, 0xa1,
	0xc7, 0xd0, 0xe5, 0xc8, 0x9a, 0x08, 0x64, 0xa2, 0x52, 0xb8, 0x50, 0x3a, 0x73, 0x3b, 0x69, 0x5d,
	0x2a, 0x87, 0x55, 0xb7, 0x06, 0x58, 0xe5, 0x2c, 0x80, 0xd1, 0x1d, 0xa6, 0x8d, 0xba, 0xc5, 0x54,
	0xb1, 0x3f, 0x1e, 0x8c, 0x2a, 0x91, 0xa3, 0xe9, 0xf7, 0xc8, 0xb0, 0x35, 0xea, 0x3f, 0x8f, 0xbb,
	0x83, 0xf8, 0xbc, 0x9e, 0x1b, 0xcd, 0x75, 0x0b, 0x47, 0xa6, 0x05, 0xf7, 0x01, 0xec, 0x0a, 0x4b,
	0x0b, 0x2b, 0xb1, 0x0c, 0x8f, 0x61, 0x90, 0x73, 0x39, 0xd5, 0x2f, 0x23, 0x98, 0xec, 0xf5, 0x70,
	0x9f, 0xb9, 0x9c, 0x1a, 0xa0, 0x46, 0x10, 0xf2, 0x4e, 0x48, 0x7a, 0xf1, 0xeb, 0xc8, 0x4b, 0x21,
	0x95, 0x45, 0x12, 0x82, 0x4a, 0xd3, 0xf7, 0xca, 0x7f, 0x50, 0xfa, 0x2b, 0x5a, 0x1c, 0x9d, 0xef,
	0xbf, 0x83, 0xad, 0x96, 0xf2, 0x26, 0x6f, 0x8b, 0xf2, 0xda, 0x1e, 0x36, 0xca, 0x3b, 0x85, 0x51,
	0xd7, 0xd3, 0xa6, 0x89, 0x5d, 0x8b, 0x9b, 0x32, 0x6d, 0x5b, 0xde, 0x24, 0xef, 0x13, 0xdc, 0x6c,
	0xe9, 0xbf, 0x45, 0x26, 0x66, 0xb7, 0x43, 0x6d, 0x9d, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x82, 0x2d, 0xe3, 0x48, 0x06, 0x00, 0x00,
}
