// Code generated by protoc-gen-go. DO NOT EDIT.
// source: broker.proto

package stack_rpc_broker

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type PublishRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{1}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type SubscribeRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Queue                string   `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{2}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SubscribeRequest) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

type Message struct {
	Header               map[string]string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{3}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "stack.rpc.broker.Empty")
	proto.RegisterType((*PublishRequest)(nil), "stack.rpc.broker.PublishRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "stack.rpc.broker.SubscribeRequest")
	proto.RegisterType((*Message)(nil), "stack.rpc.broker.Message")
	proto.RegisterMapType((map[string]string)(nil), "stack.rpc.broker.Message.HeaderEntry")
}

func init() { proto.RegisterFile("broker.proto", fileDescriptor_f209535e190f2bed) }

var fileDescriptor_f209535e190f2bed = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x97, 0xed, 0xdb, 0x96, 0xbe, 0x1d, 0x5f, 0x4a, 0x18, 0x58, 0x7b, 0x2a, 0x01, 0xa1,
	0xa7, 0x20, 0xdd, 0x45, 0x05, 0x3d, 0x08, 0x13, 0x2f, 0x8a, 0xc4, 0xa3, 0xa7, 0xa6, 0x0b, 0xae,
	0x74, 0xb3, 0x5d, 0x92, 0x0a, 0xfd, 0x4b, 0xbc, 0xfa, 0xa7, 0xca, 0x92, 0xfa, 0x6b, 0x65, 0xde,
	0xde, 0x37, 0x7c, 0xf2, 0xbc, 0x0f, 0xcf, 0x03, 0x53, 0x2e, 0xeb, 0x4a, 0x48, 0xda, 0xc8, 0x5a,
	0xd7, 0x38, 0x54, 0x3a, 0x2f, 0x2a, 0x2a, 0x9b, 0x82, 0xda, 0x77, 0xe2, 0x81, 0xb3, 0xd8, 0x34,
	0xba, 0x23, 0x4f, 0xf0, 0xff, 0xa1, 0xe5, 0xeb, 0x52, 0xad, 0x98, 0xd8, 0xb6, 0x42, 0x69, 0x3c,
	0x03, 0x47, 0xd7, 0x4d, 0x59, 0x44, 0x28, 0x41, 0xa9, 0xcf, 0xec, 0x82, 0xe7, 0xe0, 0x6d, 0x84,
	0x52, 0xf9, 0xb3, 0x88, 0xc6, 0x09, 0x4a, 0x83, 0xec, 0x98, 0xee, 0x8b, 0xd2, 0x3b, 0x0b, 0xb0,
	0x4f, 0x92, 0x5c, 0x41, 0xf8, 0xd8, 0x72, 0x55, 0xc8, 0x92, 0x8b, 0xbf, 0xe5, 0x67, 0xe0, 0x6c,
	0x5b, 0xd1, 0x5a, 0x71, 0x9f, 0xd9, 0x85, 0xbc, 0x21, 0xf0, 0x7a, 0x51, 0x7c, 0x09, 0xee, 0x4a,
	0xe4, 0x4b, 0x21, 0x23, 0x94, 0x4c, 0xd2, 0x20, 0x3b, 0x39, 0x78, 0x9f, 0xde, 0x1a, 0x6e, 0xf1,
	0xa2, 0x65, 0xc7, 0xfa, 0x4f, 0x18, 0xc3, 0x3f, 0x5e, 0x2f, 0x3b, 0xa3, 0x3f, 0x65, 0x66, 0x8e,
	0xcf, 0x21, 0xf8, 0x81, 0xe2, 0x10, 0x26, 0x95, 0xe8, 0x7a, 0x5f, 0xbb, 0x71, 0xe7, 0xea, 0x35,
	0x5f, 0x7f, 0xbb, 0x32, 0xcb, 0xc5, 0xf8, 0x0c, 0x65, 0xef, 0x08, 0xdc, 0x6b, 0x73, 0x15, 0xdf,
	0x80, 0xd7, 0x27, 0x88, 0x93, 0xa1, 0xa7, 0xdf, 0xe1, 0xc6, 0x47, 0x43, 0xc2, 0xf6, 0x30, 0xc2,
	0xf7, 0xe0, 0x7f, 0x85, 0x85, 0xc9, 0x90, 0xdb, 0x4f, 0x32, 0x3e, 0xdc, 0x00, 0x19, 0x9d, 0x22,
	0xee, 0x9a, 0xee, 0xe7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x44, 0x26, 0x93, 0xec, 0x0b, 0x02,
	0x00, 0x00,
}
