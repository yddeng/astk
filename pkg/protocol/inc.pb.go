// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inc.proto

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

type CreateDialerReq struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 string   `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDialerReq) Reset()         { *m = CreateDialerReq{} }
func (m *CreateDialerReq) String() string { return proto.CompactTextString(m) }
func (*CreateDialerReq) ProtoMessage()    {}
func (*CreateDialerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{0}
}

func (m *CreateDialerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDialerReq.Unmarshal(m, b)
}
func (m *CreateDialerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDialerReq.Marshal(b, m, deterministic)
}
func (m *CreateDialerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDialerReq.Merge(m, src)
}
func (m *CreateDialerReq) XXX_Size() int {
	return xxx_messageInfo_CreateDialerReq.Size(m)
}
func (m *CreateDialerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDialerReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDialerReq proto.InternalMessageInfo

func (m *CreateDialerReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateDialerReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreateDialerReq) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *CreateDialerReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateDialerResp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDialerResp) Reset()         { *m = CreateDialerResp{} }
func (m *CreateDialerResp) String() string { return proto.CompactTextString(m) }
func (*CreateDialerResp) ProtoMessage()    {}
func (*CreateDialerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{1}
}

func (m *CreateDialerResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDialerResp.Unmarshal(m, b)
}
func (m *CreateDialerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDialerResp.Marshal(b, m, deterministic)
}
func (m *CreateDialerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDialerResp.Merge(m, src)
}
func (m *CreateDialerResp) XXX_Size() int {
	return xxx_messageInfo_CreateDialerResp.Size(m)
}
func (m *CreateDialerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDialerResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDialerResp proto.InternalMessageInfo

func (m *CreateDialerResp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type OpenConnectionReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OpenID               int32    `protobuf:"varint,2,opt,name=openID,proto3" json:"openID,omitempty"`
	SrcIp                string   `protobuf:"bytes,3,opt,name=srcIp,proto3" json:"srcIp,omitempty"`
	SrcPort              string   `protobuf:"bytes,4,opt,name=srcPort,proto3" json:"srcPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenConnectionReq) Reset()         { *m = OpenConnectionReq{} }
func (m *OpenConnectionReq) String() string { return proto.CompactTextString(m) }
func (*OpenConnectionReq) ProtoMessage()    {}
func (*OpenConnectionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{2}
}

func (m *OpenConnectionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenConnectionReq.Unmarshal(m, b)
}
func (m *OpenConnectionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenConnectionReq.Marshal(b, m, deterministic)
}
func (m *OpenConnectionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenConnectionReq.Merge(m, src)
}
func (m *OpenConnectionReq) XXX_Size() int {
	return xxx_messageInfo_OpenConnectionReq.Size(m)
}
func (m *OpenConnectionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenConnectionReq.DiscardUnknown(m)
}

var xxx_messageInfo_OpenConnectionReq proto.InternalMessageInfo

func (m *OpenConnectionReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OpenConnectionReq) GetOpenID() int32 {
	if m != nil {
		return m.OpenID
	}
	return 0
}

func (m *OpenConnectionReq) GetSrcIp() string {
	if m != nil {
		return m.SrcIp
	}
	return ""
}

func (m *OpenConnectionReq) GetSrcPort() string {
	if m != nil {
		return m.SrcPort
	}
	return ""
}

type OpenConnectionResp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	OpenID               int32    `protobuf:"varint,2,opt,name=openID,proto3" json:"openID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenConnectionResp) Reset()         { *m = OpenConnectionResp{} }
func (m *OpenConnectionResp) String() string { return proto.CompactTextString(m) }
func (*OpenConnectionResp) ProtoMessage()    {}
func (*OpenConnectionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{3}
}

func (m *OpenConnectionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenConnectionResp.Unmarshal(m, b)
}
func (m *OpenConnectionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenConnectionResp.Marshal(b, m, deterministic)
}
func (m *OpenConnectionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenConnectionResp.Merge(m, src)
}
func (m *OpenConnectionResp) XXX_Size() int {
	return xxx_messageInfo_OpenConnectionResp.Size(m)
}
func (m *OpenConnectionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenConnectionResp.DiscardUnknown(m)
}

var xxx_messageInfo_OpenConnectionResp proto.InternalMessageInfo

func (m *OpenConnectionResp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *OpenConnectionResp) GetOpenID() int32 {
	if m != nil {
		return m.OpenID
	}
	return 0
}

type CloseConnectionReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OpenID               int32    `protobuf:"varint,2,opt,name=openID,proto3" json:"openID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseConnectionReq) Reset()         { *m = CloseConnectionReq{} }
func (m *CloseConnectionReq) String() string { return proto.CompactTextString(m) }
func (*CloseConnectionReq) ProtoMessage()    {}
func (*CloseConnectionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{4}
}

func (m *CloseConnectionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseConnectionReq.Unmarshal(m, b)
}
func (m *CloseConnectionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseConnectionReq.Marshal(b, m, deterministic)
}
func (m *CloseConnectionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseConnectionReq.Merge(m, src)
}
func (m *CloseConnectionReq) XXX_Size() int {
	return xxx_messageInfo_CloseConnectionReq.Size(m)
}
func (m *CloseConnectionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseConnectionReq.DiscardUnknown(m)
}

var xxx_messageInfo_CloseConnectionReq proto.InternalMessageInfo

func (m *CloseConnectionReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CloseConnectionReq) GetOpenID() int32 {
	if m != nil {
		return m.OpenID
	}
	return 0
}

type CloseConnectionResp struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseConnectionResp) Reset()         { *m = CloseConnectionResp{} }
func (m *CloseConnectionResp) String() string { return proto.CompactTextString(m) }
func (*CloseConnectionResp) ProtoMessage()    {}
func (*CloseConnectionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{5}
}

func (m *CloseConnectionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseConnectionResp.Unmarshal(m, b)
}
func (m *CloseConnectionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseConnectionResp.Marshal(b, m, deterministic)
}
func (m *CloseConnectionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseConnectionResp.Merge(m, src)
}
func (m *CloseConnectionResp) XXX_Size() int {
	return xxx_messageInfo_CloseConnectionResp.Size(m)
}
func (m *CloseConnectionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseConnectionResp.DiscardUnknown(m)
}

var xxx_messageInfo_CloseConnectionResp proto.InternalMessageInfo

func (m *CloseConnectionResp) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type IncConnMessage struct {
	OpenID               int32    `protobuf:"varint,1,opt,name=openID,proto3" json:"openID,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncConnMessage) Reset()         { *m = IncConnMessage{} }
func (m *IncConnMessage) String() string { return proto.CompactTextString(m) }
func (*IncConnMessage) ProtoMessage()    {}
func (*IncConnMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9ee4de5949f8e83, []int{6}
}

func (m *IncConnMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncConnMessage.Unmarshal(m, b)
}
func (m *IncConnMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncConnMessage.Marshal(b, m, deterministic)
}
func (m *IncConnMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncConnMessage.Merge(m, src)
}
func (m *IncConnMessage) XXX_Size() int {
	return xxx_messageInfo_IncConnMessage.Size(m)
}
func (m *IncConnMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_IncConnMessage.DiscardUnknown(m)
}

var xxx_messageInfo_IncConnMessage proto.InternalMessageInfo

func (m *IncConnMessage) GetOpenID() int32 {
	if m != nil {
		return m.OpenID
	}
	return 0
}

func (m *IncConnMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateDialerReq)(nil), "CreateDialerReq")
	proto.RegisterType((*CreateDialerResp)(nil), "CreateDialerResp")
	proto.RegisterType((*OpenConnectionReq)(nil), "OpenConnectionReq")
	proto.RegisterType((*OpenConnectionResp)(nil), "OpenConnectionResp")
	proto.RegisterType((*CloseConnectionReq)(nil), "CloseConnectionReq")
	proto.RegisterType((*CloseConnectionResp)(nil), "CloseConnectionResp")
	proto.RegisterType((*IncConnMessage)(nil), "IncConnMessage")
}

func init() { proto.RegisterFile("inc.proto", fileDescriptor_a9ee4de5949f8e83) }

var fileDescriptor_a9ee4de5949f8e83 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0xd9, 0x75, 0xef, 0xf4, 0x06, 0x39, 0x35, 0x8a, 0xa4, 0x94, 0x14, 0xa2, 0x8d, 0x8d,
	0xed, 0x15, 0xe2, 0x5e, 0xb3, 0x85, 0x28, 0x8b, 0x8d, 0x76, 0x31, 0x3b, 0x48, 0x70, 0xc9, 0xc4,
	0x24, 0x8d, 0xff, 0x5e, 0x32, 0xbb, 0x27, 0x9e, 0xba, 0x8d, 0xdd, 0x9b, 0xf0, 0xe6, 0xbd, 0x8f,
	0x09, 0x2c, 0xac, 0x33, 0x57, 0x3e, 0x50, 0x22, 0xf5, 0x04, 0x07, 0x75, 0x40, 0x9d, 0x70, 0x6d,
	0x75, 0x8f, 0xa1, 0xc5, 0x77, 0x21, 0xa0, 0x7a, 0xfc, 0xf0, 0x28, 0x8b, 0xb3, 0xe2, 0x62, 0xd1,
	0xb2, 0x16, 0x4b, 0x28, 0xad, 0x97, 0x25, 0xbf, 0x94, 0xd6, 0x67, 0x8f, 0xa7, 0x90, 0xe4, 0xce,
	0xe0, 0xc9, 0x9a, 0x3d, 0x9d, 0xac, 0x46, 0x4f, 0xa7, 0xce, 0xe1, 0x70, 0x3b, 0x3a, 0xf2, 0x9e,
	0xa1, 0xee, 0x2b, 0x3b, 0x6b, 0xf5, 0x06, 0x47, 0xf7, 0x1e, 0x5d, 0x4d, 0xce, 0xa1, 0x49, 0x96,
	0x5c, 0x86, 0x18, 0xc2, 0x8a, 0x4d, 0x98, 0x38, 0x85, 0x39, 0x79, 0x74, 0xcd, 0x9a, 0x21, 0x66,
	0xed, 0x38, 0x89, 0x13, 0x98, 0xc5, 0x60, 0x1a, 0x3f, 0x92, 0x0c, 0x83, 0x90, 0xb0, 0x1b, 0x83,
	0x79, 0xc8, 0x84, 0x03, 0xcf, 0x66, 0x54, 0x37, 0x20, 0x7e, 0x96, 0xfd, 0x8d, 0x35, 0xd5, 0xa8,
	0x56, 0x20, 0xea, 0x9e, 0x22, 0xfe, 0x8b, 0x57, 0x5d, 0xc2, 0xf1, 0xaf, 0xed, 0x89, 0xbb, 0xac,
	0x60, 0xd9, 0x38, 0x93, 0x8d, 0x77, 0x18, 0xa3, 0x7e, 0xfd, 0x8e, 0x54, 0x6c, 0x1d, 0x41, 0x40,
	0xd5, 0xe9, 0xa4, 0xb9, 0x6a, 0xbf, 0x65, 0x7d, 0x0b, 0xcf, 0x7b, 0xfc, 0xc3, 0x86, 0xfa, 0x97,
	0x39, 0xab, 0xeb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x61, 0x35, 0xcb, 0xf8, 0x01, 0x00,
	0x00,
}