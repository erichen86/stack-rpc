// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package stack_rpc_network

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/stack-labs/stack-rpc/router/proto"
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

// Query is passed in a LookupRequest
type Query struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Gateway              string   `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Router               string   `protobuf:"bytes,4,opt,name=router,proto3" json:"router,omitempty"`
	Network              string   `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Query) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Query) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Query) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *Query) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

type ConnectRequest struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type ConnectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectResponse) Reset()         { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()    {}
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2}
}

func (m *ConnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectResponse.Unmarshal(m, b)
}
func (m *ConnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectResponse.Marshal(b, m, deterministic)
}
func (m *ConnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectResponse.Merge(m, src)
}
func (m *ConnectResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectResponse.Size(m)
}
func (m *ConnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectResponse proto.InternalMessageInfo

// PeerRequest requests list of peers
type NodesRequest struct {
	// node topology depth
	Depth                uint32   `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesRequest) Reset()         { *m = NodesRequest{} }
func (m *NodesRequest) String() string { return proto.CompactTextString(m) }
func (*NodesRequest) ProtoMessage()    {}
func (*NodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{3}
}

func (m *NodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesRequest.Unmarshal(m, b)
}
func (m *NodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesRequest.Marshal(b, m, deterministic)
}
func (m *NodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesRequest.Merge(m, src)
}
func (m *NodesRequest) XXX_Size() int {
	return xxx_messageInfo_NodesRequest.Size(m)
}
func (m *NodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodesRequest proto.InternalMessageInfo

func (m *NodesRequest) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

// PeerResponse is returned by ListPeers
type NodesResponse struct {
	// return peer topology
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodesResponse) Reset()         { *m = NodesResponse{} }
func (m *NodesResponse) String() string { return proto.CompactTextString(m) }
func (*NodesResponse) ProtoMessage()    {}
func (*NodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{4}
}

func (m *NodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodesResponse.Unmarshal(m, b)
}
func (m *NodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodesResponse.Marshal(b, m, deterministic)
}
func (m *NodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodesResponse.Merge(m, src)
}
func (m *NodesResponse) XXX_Size() int {
	return xxx_messageInfo_NodesResponse.Size(m)
}
func (m *NodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodesResponse proto.InternalMessageInfo

func (m *NodesResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type GraphRequest struct {
	// node topology depth
	Depth                uint32   `protobuf:"varint,1,opt,name=depth,proto3" json:"depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphRequest) Reset()         { *m = GraphRequest{} }
func (m *GraphRequest) String() string { return proto.CompactTextString(m) }
func (*GraphRequest) ProtoMessage()    {}
func (*GraphRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{5}
}

func (m *GraphRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphRequest.Unmarshal(m, b)
}
func (m *GraphRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphRequest.Marshal(b, m, deterministic)
}
func (m *GraphRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphRequest.Merge(m, src)
}
func (m *GraphRequest) XXX_Size() int {
	return xxx_messageInfo_GraphRequest.Size(m)
}
func (m *GraphRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GraphRequest proto.InternalMessageInfo

func (m *GraphRequest) GetDepth() uint32 {
	if m != nil {
		return m.Depth
	}
	return 0
}

type GraphResponse struct {
	Root                 *Peer    `protobuf:"bytes,1,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphResponse) Reset()         { *m = GraphResponse{} }
func (m *GraphResponse) String() string { return proto.CompactTextString(m) }
func (*GraphResponse) ProtoMessage()    {}
func (*GraphResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{6}
}

func (m *GraphResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphResponse.Unmarshal(m, b)
}
func (m *GraphResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphResponse.Marshal(b, m, deterministic)
}
func (m *GraphResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphResponse.Merge(m, src)
}
func (m *GraphResponse) XXX_Size() int {
	return xxx_messageInfo_GraphResponse.Size(m)
}
func (m *GraphResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GraphResponse proto.InternalMessageInfo

func (m *GraphResponse) GetRoot() *Peer {
	if m != nil {
		return m.Root
	}
	return nil
}

type RoutesRequest struct {
	// filter based on
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutesRequest) Reset()         { *m = RoutesRequest{} }
func (m *RoutesRequest) String() string { return proto.CompactTextString(m) }
func (*RoutesRequest) ProtoMessage()    {}
func (*RoutesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{7}
}

func (m *RoutesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesRequest.Unmarshal(m, b)
}
func (m *RoutesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesRequest.Marshal(b, m, deterministic)
}
func (m *RoutesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesRequest.Merge(m, src)
}
func (m *RoutesRequest) XXX_Size() int {
	return xxx_messageInfo_RoutesRequest.Size(m)
}
func (m *RoutesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesRequest proto.InternalMessageInfo

func (m *RoutesRequest) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type RoutesResponse struct {
	Routes               []*proto1.Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RoutesResponse) Reset()         { *m = RoutesResponse{} }
func (m *RoutesResponse) String() string { return proto.CompactTextString(m) }
func (*RoutesResponse) ProtoMessage()    {}
func (*RoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{8}
}

func (m *RoutesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutesResponse.Unmarshal(m, b)
}
func (m *RoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutesResponse.Marshal(b, m, deterministic)
}
func (m *RoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutesResponse.Merge(m, src)
}
func (m *RoutesResponse) XXX_Size() int {
	return xxx_messageInfo_RoutesResponse.Size(m)
}
func (m *RoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoutesResponse proto.InternalMessageInfo

func (m *RoutesResponse) GetRoutes() []*proto1.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type ServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServicesRequest) Reset()         { *m = ServicesRequest{} }
func (m *ServicesRequest) String() string { return proto.CompactTextString(m) }
func (*ServicesRequest) ProtoMessage()    {}
func (*ServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{9}
}

func (m *ServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServicesRequest.Unmarshal(m, b)
}
func (m *ServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServicesRequest.Marshal(b, m, deterministic)
}
func (m *ServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServicesRequest.Merge(m, src)
}
func (m *ServicesRequest) XXX_Size() int {
	return xxx_messageInfo_ServicesRequest.Size(m)
}
func (m *ServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServicesRequest proto.InternalMessageInfo

type ServicesResponse struct {
	Services             []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServicesResponse) Reset()         { *m = ServicesResponse{} }
func (m *ServicesResponse) String() string { return proto.CompactTextString(m) }
func (*ServicesResponse) ProtoMessage()    {}
func (*ServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{10}
}

func (m *ServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServicesResponse.Unmarshal(m, b)
}
func (m *ServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServicesResponse.Marshal(b, m, deterministic)
}
func (m *ServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServicesResponse.Merge(m, src)
}
func (m *ServicesResponse) XXX_Size() int {
	return xxx_messageInfo_ServicesResponse.Size(m)
}
func (m *ServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServicesResponse proto.InternalMessageInfo

func (m *ServicesResponse) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

// Node is network node
type Node struct {
	// node id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// node address
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// the network
	Network string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	// associated metadata
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{11}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Node) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Connect is sent when the node connects to the network
type Connect struct {
	// network mode
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{12}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Close is sent when the node disconnects from the network
type Close struct {
	// network node
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{13}
}

func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (m *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(m, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func (m *Close) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Peer is used to advertise node peers
type Peer struct {
	// network node
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// node peers
	Peers                []*Peer  `protobuf:"bytes,2,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{14}
}

func (m *Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Peer.Unmarshal(m, b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return xxx_messageInfo_Peer.Size(m)
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Peer) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*Query)(nil), "stack.rpc.network.Query")
	proto.RegisterType((*ConnectRequest)(nil), "stack.rpc.network.ConnectRequest")
	proto.RegisterType((*ConnectResponse)(nil), "stack.rpc.network.ConnectResponse")
	proto.RegisterType((*NodesRequest)(nil), "stack.rpc.network.NodesRequest")
	proto.RegisterType((*NodesResponse)(nil), "stack.rpc.network.NodesResponse")
	proto.RegisterType((*GraphRequest)(nil), "stack.rpc.network.GraphRequest")
	proto.RegisterType((*GraphResponse)(nil), "stack.rpc.network.GraphResponse")
	proto.RegisterType((*RoutesRequest)(nil), "stack.rpc.network.RoutesRequest")
	proto.RegisterType((*RoutesResponse)(nil), "stack.rpc.network.RoutesResponse")
	proto.RegisterType((*ServicesRequest)(nil), "stack.rpc.network.ServicesRequest")
	proto.RegisterType((*ServicesResponse)(nil), "stack.rpc.network.ServicesResponse")
	proto.RegisterType((*Node)(nil), "stack.rpc.network.Node")
	proto.RegisterMapType((map[string]string)(nil), "stack.rpc.network.Node.MetadataEntry")
	proto.RegisterType((*Connect)(nil), "stack.rpc.network.Connect")
	proto.RegisterType((*Close)(nil), "stack.rpc.network.Close")
	proto.RegisterType((*Peer)(nil), "stack.rpc.network.Peer")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x6d, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0xb8, 0x49, 0x87, 0x3a, 0xb4, 0x2b, 0x04, 0x96, 0xff, 0x90, 0x2e, 0x45, 0xaa,
	0x84, 0xea, 0x48, 0x05, 0x21, 0xbe, 0x44, 0x55, 0x55, 0x88, 0x3f, 0x50, 0xc0, 0x88, 0x03, 0x38,
	0xf6, 0xa8, 0x89, 0x92, 0x7a, 0xdd, 0xf5, 0xa6, 0x55, 0x4e, 0xc0, 0x39, 0xb8, 0x0e, 0xa7, 0x42,
	0xbb, 0x3b, 0x4e, 0x63, 0x70, 0x2d, 0xfa, 0xcf, 0x93, 0x79, 0xf3, 0x66, 0x67, 0xe6, 0xbd, 0x80,
	0x9f, 0xa3, 0xba, 0x16, 0x72, 0x16, 0x15, 0x52, 0x28, 0xc1, 0x76, 0x4b, 0x95, 0xa4, 0xb3, 0x48,
	0x16, 0x69, 0x44, 0x89, 0xf0, 0xf5, 0xf9, 0x54, 0x4d, 0x16, 0xe3, 0x28, 0x15, 0x17, 0x23, 0x93,
	0x3d, 0x9c, 0x27, 0xe3, 0x92, 0x3e, 0x65, 0x91, 0x8e, 0xa4, 0x58, 0x28, 0x94, 0x23, 0x43, 0x40,
	0x81, 0x65, 0xe3, 0x3f, 0x1d, 0xf0, 0xbe, 0x2d, 0x50, 0x2e, 0x59, 0x00, 0xbd, 0x12, 0xe5, 0xd5,
	0x34, 0xc5, 0xc0, 0x19, 0x3a, 0x07, 0x5b, 0x71, 0x15, 0xea, 0x4c, 0x92, 0x65, 0x12, 0xcb, 0x32,
	0xe8, 0xd8, 0x0c, 0x85, 0x3a, 0x73, 0x9e, 0x28, 0xbc, 0x4e, 0x96, 0x81, 0x6b, 0x33, 0x14, 0xb2,
	0x87, 0xb0, 0x69, 0xfb, 0x04, 0x5d, 0x93, 0xa0, 0x48, 0x57, 0xd0, 0xab, 0x03, 0xcf, 0x56, 0x50,
	0xc8, 0x8f, 0x61, 0x70, 0x2a, 0xf2, 0x1c, 0x53, 0x15, 0xe3, 0xe5, 0x02, 0x4b, 0xc5, 0x0e, 0xc1,
	0xcb, 0x45, 0x86, 0x65, 0xe0, 0x0c, 0xdd, 0x83, 0x7b, 0x47, 0x8f, 0xa2, 0x7f, 0x26, 0x8f, 0xce,
	0x44, 0x86, 0xb1, 0x45, 0xf1, 0x5d, 0xb8, 0xbf, 0x22, 0x28, 0x0b, 0x91, 0x97, 0xc8, 0xf7, 0x61,
	0x5b, 0x23, 0xca, 0x8a, 0xf1, 0x01, 0x78, 0x19, 0x16, 0x6a, 0x62, 0x26, 0xf4, 0x63, 0x1b, 0xf0,
	0xf7, 0xe0, 0x13, 0xca, 0x96, 0xdd, 0xb5, 0xf1, 0x3e, 0x6c, 0x7f, 0x94, 0x49, 0x31, 0x69, 0xef,
	0xf2, 0x0e, 0x7c, 0x42, 0x51, 0x97, 0x67, 0xd0, 0x95, 0x42, 0x28, 0x83, 0x6a, 0x6e, 0xf2, 0x15,
	0x51, 0xc6, 0x06, 0xc4, 0x8f, 0xc1, 0x8f, 0xf5, 0x06, 0x57, 0xa3, 0x44, 0xe0, 0x5d, 0xea, 0xbb,
	0x51, 0x79, 0xd0, 0x50, 0x6e, 0xee, 0x1a, 0x5b, 0x18, 0x3f, 0x81, 0x41, 0x45, 0x40, 0xfd, 0x47,
	0x74, 0xa2, 0xa6, 0x31, 0x49, 0x23, 0xa6, 0x82, 0x6e, 0x67, 0x16, 0xfc, 0xdd, 0x4a, 0xa2, 0x7a,
	0x05, 0x8f, 0x60, 0xe7, 0xe6, 0x27, 0xe2, 0x0d, 0xa1, 0x4f, 0xca, 0xb1, 0xcc, 0x5b, 0xf1, 0x2a,
	0xe6, 0xbf, 0x1d, 0xe8, 0xea, 0xd5, 0xb1, 0x01, 0x74, 0xa6, 0x19, 0x09, 0xad, 0x33, 0xcd, 0xda,
	0x35, 0x56, 0x29, 0xc6, 0xad, 0x29, 0x86, 0x9d, 0x40, 0xff, 0x02, 0x55, 0x92, 0x25, 0x2a, 0x09,
	0xba, 0x66, 0x84, 0xa7, 0xb7, 0x5c, 0x2a, 0xfa, 0x4c, 0xb8, 0x0f, 0xb9, 0x92, 0xcb, 0x78, 0x55,
	0x16, 0xbe, 0x05, 0xbf, 0x96, 0x62, 0x3b, 0xe0, 0xce, 0x70, 0x49, 0x0f, 0xd3, 0x9f, 0xfa, 0x9a,
	0x57, 0xc9, 0x7c, 0x81, 0xf4, 0x2e, 0x1b, 0xbc, 0xe9, 0xbc, 0x72, 0xf8, 0x4b, 0xe8, 0x91, 0xe0,
	0xf4, 0x2d, 0xb5, 0x16, 0x5a, 0x6e, 0x69, 0x04, 0x63, 0x40, 0xfc, 0x05, 0x78, 0xa7, 0x73, 0x61,
	0x15, 0xf0, 0xff, 0x55, 0x63, 0xe8, 0x6a, 0x3d, 0xdc, 0xa9, 0x48, 0x2b, 0xb9, 0x40, 0x94, 0x7a,
	0xa9, 0x6e, 0x9b, 0xc8, 0x2c, 0xea, 0xe8, 0x97, 0x0b, 0xbd, 0x33, 0xda, 0x6e, 0x7c, 0x33, 0xdd,
	0x5e, 0x43, 0x59, 0xdd, 0xab, 0x21, 0x6f, 0x83, 0x90, 0x1b, 0x37, 0xd8, 0x27, 0xf0, 0x8c, 0x07,
	0xd8, 0xe3, 0x06, 0xf8, 0xba, 0x87, 0xc2, 0xe1, 0xed, 0x80, 0x75, 0x36, 0xe3, 0xdb, 0x46, 0xb6,
	0x75, 0xdf, 0x37, 0xb2, 0xd5, 0x2c, 0xcf, 0x37, 0xd8, 0x17, 0xd8, 0xb4, 0x06, 0x61, 0x4d, 0xe8,
	0x9a, 0xf9, 0xc2, 0xbd, 0x16, 0xc4, 0x8a, 0xf0, 0x07, 0xf4, 0x2b, 0x6f, 0xb0, 0xa6, 0xf5, 0xfc,
	0xe5, 0xa5, 0xf0, 0x49, 0x2b, 0xa6, 0xa2, 0x1d, 0x6f, 0x9a, 0x3f, 0xee, 0xe7, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x80, 0x12, 0xaa, 0x54, 0x17, 0x06, 0x00, 0x00,
}
