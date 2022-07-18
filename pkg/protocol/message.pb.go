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

type LogFileReq struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Payload              int32    `protobuf:"varint,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogFileReq) Reset()         { *m = LogFileReq{} }
func (m *LogFileReq) String() string { return proto.CompactTextString(m) }
func (*LogFileReq) ProtoMessage()    {}
func (*LogFileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{11}
}

func (m *LogFileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogFileReq.Unmarshal(m, b)
}
func (m *LogFileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogFileReq.Marshal(b, m, deterministic)
}
func (m *LogFileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogFileReq.Merge(m, src)
}
func (m *LogFileReq) XXX_Size() int {
	return xxx_messageInfo_LogFileReq.Size(m)
}
func (m *LogFileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LogFileReq.DiscardUnknown(m)
}

var xxx_messageInfo_LogFileReq proto.InternalMessageInfo

func (m *LogFileReq) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *LogFileReq) GetPayload() int32 {
	if m != nil {
		return m.Payload
	}
	return 0
}

type LogFileResp struct {
	Context              []byte   `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogFileResp) Reset()         { *m = LogFileResp{} }
func (m *LogFileResp) String() string { return proto.CompactTextString(m) }
func (*LogFileResp) ProtoMessage()    {}
func (*LogFileResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{12}
}

func (m *LogFileResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogFileResp.Unmarshal(m, b)
}
func (m *LogFileResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogFileResp.Marshal(b, m, deterministic)
}
func (m *LogFileResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogFileResp.Merge(m, src)
}
func (m *LogFileResp) XXX_Size() int {
	return xxx_messageInfo_LogFileResp.Size(m)
}
func (m *LogFileResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LogFileResp.DiscardUnknown(m)
}

var xxx_messageInfo_LogFileResp proto.InternalMessageInfo

func (m *LogFileResp) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
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
	proto.RegisterType((*LogFileReq)(nil), "logFileReq")
	proto.RegisterType((*LogFileResp)(nil), "logFileResp")
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
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0x96, 0xed, 0x38, 0x4d, 0xc6, 0x40, 0x53, 0x17, 0x55, 0x56, 0xaa, 0xaa, 0x68, 0x11, 0x82,
	0x93, 0x0f, 0xa0, 0x96, 0x52, 0xf5, 0x04, 0xa5, 0xe2, 0x42, 0x0f, 0xe6, 0xc6, 0xa9, 0xc6, 0x5e,
	0xcc, 0x2a, 0xb6, 0xd7, 0xf5, 0x6e, 0x2a, 0xf2, 0x1c, 0x7d, 0xab, 0x3e, 0x4c, 0x9f, 0xa1, 0x9a,
	0xfd, 0x31, 0x4e, 0x00, 0xa9, 0xb9, 0x7d, 0x33, 0xfb, 0xcd, 0xdf, 0x37, 0xbb, 0x36, 0x6c, 0x56,
	0x54, 0x88, 0xb4, 0xa0, 0x71, 0xd3, 0x72, 0xc9, 0xc9, 0x35, 0x8c, 0x4a, 0x5e, 0xb0, 0x3a, 0xa1,
	0x3f, 0xc3, 0x10, 0x06, 0x75, 0x5a, 0xd1, 0xc8, 0xd9, 0x71, 0x0e, 0xc6, 0x89, 0xc2, 0xe8, 0x63,
	0x35, 0x95, 0x91, 0xab, 0x7d, 0x88, 0xc3, 0x09, 0x78, 0xe8, 0xf2, 0x94, 0x0b, 0x61, 0xb8, 0x0d,
	0xbe, 0xe4, 0x33, 0x5a, 0x47, 0x03, 0xe5, 0xd3, 0x06, 0x79, 0x0f, 0x63, 0x93, 0x5b, 0x34, 0x98,
	0x28, 0xe3, 0x79, 0x97, 0x1c, 0x31, 0xf9, 0x01, 0x90, 0x55, 0xf9, 0xf9, 0x3d, 0xcd, 0xb0, 0xfc,
	0x04, 0xbc, 0x9c, 0xb5, 0x86, 0x80, 0xb0, 0x6b, 0xc8, 0x5d, 0x6e, 0x28, 0x6d, 0x0b, 0x11, 0x79,
	0x3b, 0x1e, 0xfa, 0x10, 0x87, 0x11, 0xbc, 0x90, 0xac, 0xa2, 0x7c, 0x2e, 0x55, 0x03, 0x7e, 0x62,
	0x4d, 0x72, 0x02, 0x41, 0x57, 0xe1, 0xe9, 0x26, 0xc2, 0x37, 0x30, 0xe4, 0x73, 0x79, 0x25, 0x5b,
	0x53, 0xc6, 0x58, 0xe4, 0x8f, 0x03, 0x5b, 0x4d, 0xcb, 0x33, 0x2a, 0x84, 0xed, 0xf0, 0x08, 0x86,
	0x19, 0xaf, 0x6f, 0x59, 0x11, 0x39, 0x3b, 0xde, 0x41, 0x70, 0xf8, 0x36, 0x5e, 0x26, 0xc4, 0x67,
	0xea, 0xf4, 0xbc, 0x96, 0xed, 0x22, 0x31, 0xd4, 0xff, 0x1e, 0xc2, 0x8c, 0x3f, 0x78, 0x18, 0x7f,
	0x0b, 0x5c, 0x96, 0x47, 0xbe, 0x9a, 0xc8, 0x65, 0xf9, 0xf4, 0x04, 0x82, 0x5e, 0x01, 0x0c, 0x98,
	0xd1, 0x85, 0xd5, 0x6b, 0x46, 0x17, 0xb8, 0x86, 0x5f, 0x69, 0x39, 0xb7, 0xb5, 0xb4, 0xf1, 0xd9,
	0xfd, 0xe4, 0x90, 0x63, 0x78, 0xb9, 0xd4, 0xea, 0x33, 0x5a, 0x4c, 0xc0, 0x6b, 0x58, 0xae, 0xc2,
	0xfd, 0x04, 0x21, 0xf9, 0x02, 0x13, 0x13, 0x78, 0xc5, 0x8a, 0x3a, 0x2d, 0xcd, 0xa2, 0x90, 0xe5,
	0x74, 0x2c, 0xd4, 0x50, 0xa8, 0x63, 0x13, 0x6a, 0x2c, 0xb2, 0x0f, 0xaf, 0x56, 0xa2, 0x9f, 0xb9,
	0x09, 0xbb, 0x5d, 0x7f, 0x57, 0x32, 0x95, 0xd4, 0x54, 0x61, 0xb9, 0x50, 0x4a, 0xfb, 0x09, 0x42,
	0xf2, 0xdb, 0x79, 0x68, 0x46, 0xb3, 0x44, 0x13, 0x7e, 0x80, 0xa1, 0x40, 0x43, 0x98, 0x9d, 0xbc,
	0x8b, 0x57, 0x29, 0xb1, 0x42, 0xc2, 0x6c, 0x45, 0x93, 0xa7, 0x17, 0x10, 0xf4, 0xdc, 0x7d, 0x2d,
	0x7d, 0xad, 0xe5, 0x6e, 0x5f, 0xcb, 0xe0, 0x70, 0x73, 0x39, 0x6d, 0x4f, 0x5a, 0x09, 0x1b, 0xfd,
	0xa3, 0x27, 0xd4, 0xd9, 0x06, 0x5f, 0x55, 0xb5, 0x6b, 0x51, 0x06, 0x5e, 0x5a, 0x7a, 0xcf, 0xe4,
	0xa5, 0x28, 0xcc, 0x4b, 0xb2, 0x26, 0x66, 0xc8, 0x9a, 0xb9, 0xba, 0x09, 0x4e, 0x82, 0x10, 0x3d,
	0x15, 0xad, 0xd4, 0x55, 0x70, 0x12, 0x84, 0xe4, 0x14, 0xa0, 0xe4, 0xc5, 0x37, 0x56, 0x2a, 0xad,
	0xa6, 0x30, 0xba, 0x65, 0x25, 0xed, 0xbd, 0xde, 0xce, 0xc6, 0x3a, 0x4d, 0xba, 0x28, 0x79, 0x6a,
	0xf7, 0x6a, 0x4d, 0xb2, 0x0f, 0x41, 0x97, 0x43, 0x34, 0x48, 0xcc, 0x78, 0x2d, 0xe9, 0xbd, 0x54,
	0x39, 0x36, 0x12, 0x6b, 0x92, 0x00, 0xc6, 0x77, 0x34, 0x6d, 0xe5, 0x0d, 0x4d, 0x25, 0xf9, 0xeb,
	0xc1, 0xb8, 0xe6, 0x39, 0xd5, 0xd3, 0xee, 0xe9, 0x5e, 0xb5, 0xf6, 0xaf, 0xe3, 0xee, 0x20, 0x3e,
	0x6b, 0xe6, 0x5a, 0x71, 0x35, 0xc0, 0x9e, 0x1e, 0xc0, 0x7d, 0x44, 0xbb, 0xa4, 0x95, 0xa1, 0x55,
	0xb4, 0x0a, 0x0f, 0x60, 0x90, 0x33, 0x31, 0x53, 0xef, 0x22, 0x38, 0xdc, 0xee, 0xf1, 0xbe, 0x32,
	0x31, 0xd3, 0x44, 0xc5, 0x40, 0xe6, 0x1d, 0x17, 0xf8, 0xde, 0x57, 0x99, 0x17, 0x5c, 0x48, 0xc3,
	0x44, 0x06, 0x96, 0xc6, 0xaf, 0x95, 0xff, 0xa8, 0xf4, 0x77, 0x6a, 0x78, 0x78, 0x3e, 0xfd, 0x08,
	0x23, 0xdb, 0xf2, 0x3a, 0x2f, 0x0b, 0xe3, 0xec, 0x0c, 0x6b, 0xc5, 0x1d, 0xc3, 0xb8, 0x9b, 0x69,
	0xdd, 0xc0, 0x6e, 0xc4, 0x75, 0x3b, 0xb5, 0x23, 0xaf, 0x13, 0x77, 0x0a, 0xd7, 0x23, 0xf5, 0xaf,
	0xc8, 0x78, 0x79, 0x33, 0x54, 0xe8, 0xe8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xa7, 0xec,
	0x10, 0x46, 0x06, 0x00, 0x00,
}
