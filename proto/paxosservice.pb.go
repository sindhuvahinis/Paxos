// Code generated by protoc-gen-go. DO NOT EDIT.
// source: paxosservice.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PrepareRequest struct {
	ProcessId            int64    `protobuf:"varint,1,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	ProposeNumber        string   `protobuf:"bytes,2,opt,name=propose_number,json=proposeNumber,proto3" json:"propose_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrepareRequest) Reset()         { *m = PrepareRequest{} }
func (m *PrepareRequest) String() string { return proto.CompactTextString(m) }
func (*PrepareRequest) ProtoMessage()    {}
func (*PrepareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66e7684f91bf7b3, []int{0}
}

func (m *PrepareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrepareRequest.Unmarshal(m, b)
}
func (m *PrepareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrepareRequest.Marshal(b, m, deterministic)
}
func (m *PrepareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrepareRequest.Merge(m, src)
}
func (m *PrepareRequest) XXX_Size() int {
	return xxx_messageInfo_PrepareRequest.Size(m)
}
func (m *PrepareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrepareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrepareRequest proto.InternalMessageInfo

func (m *PrepareRequest) GetProcessId() int64 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *PrepareRequest) GetProposeNumber() string {
	if m != nil {
		return m.ProposeNumber
	}
	return ""
}

type Promise struct {
	ProposeNumber         string   `protobuf:"bytes,2,opt,name=propose_number,json=proposeNumber,proto3" json:"propose_number,omitempty"`
	AcceptedProposeNumber string   `protobuf:"bytes,3,opt,name=accepted_propose_number,json=acceptedProposeNumber,proto3" json:"accepted_propose_number,omitempty"`
	AcceptedKey           string   `protobuf:"bytes,4,opt,name=accepted_key,json=acceptedKey,proto3" json:"accepted_key,omitempty"`
	AcceptedValue         string   `protobuf:"bytes,5,opt,name=accepted_value,json=acceptedValue,proto3" json:"accepted_value,omitempty"`
	IsPromised            bool     `protobuf:"varint,6,opt,name=is_promised,json=isPromised,proto3" json:"is_promised,omitempty"`
	OperationId           int64    `protobuf:"varint,7,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Promise) Reset()         { *m = Promise{} }
func (m *Promise) String() string { return proto.CompactTextString(m) }
func (*Promise) ProtoMessage()    {}
func (*Promise) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66e7684f91bf7b3, []int{1}
}

func (m *Promise) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Promise.Unmarshal(m, b)
}
func (m *Promise) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Promise.Marshal(b, m, deterministic)
}
func (m *Promise) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Promise.Merge(m, src)
}
func (m *Promise) XXX_Size() int {
	return xxx_messageInfo_Promise.Size(m)
}
func (m *Promise) XXX_DiscardUnknown() {
	xxx_messageInfo_Promise.DiscardUnknown(m)
}

var xxx_messageInfo_Promise proto.InternalMessageInfo

func (m *Promise) GetProposeNumber() string {
	if m != nil {
		return m.ProposeNumber
	}
	return ""
}

func (m *Promise) GetAcceptedProposeNumber() string {
	if m != nil {
		return m.AcceptedProposeNumber
	}
	return ""
}

func (m *Promise) GetAcceptedKey() string {
	if m != nil {
		return m.AcceptedKey
	}
	return ""
}

func (m *Promise) GetAcceptedValue() string {
	if m != nil {
		return m.AcceptedValue
	}
	return ""
}

func (m *Promise) GetIsPromised() bool {
	if m != nil {
		return m.IsPromised
	}
	return false
}

func (m *Promise) GetOperationId() int64 {
	if m != nil {
		return m.OperationId
	}
	return 0
}

type AcceptRequest struct {
	ProcessId            int64    `protobuf:"varint,1,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	ProposeNumber        string   `protobuf:"bytes,2,opt,name=propose_number,json=proposeNumber,proto3" json:"propose_number,omitempty"`
	MaxKey               string   `protobuf:"bytes,3,opt,name=max_key,json=maxKey,proto3" json:"max_key,omitempty"`
	MaxValue             string   `protobuf:"bytes,4,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcceptRequest) Reset()         { *m = AcceptRequest{} }
func (m *AcceptRequest) String() string { return proto.CompactTextString(m) }
func (*AcceptRequest) ProtoMessage()    {}
func (*AcceptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66e7684f91bf7b3, []int{2}
}

func (m *AcceptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcceptRequest.Unmarshal(m, b)
}
func (m *AcceptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcceptRequest.Marshal(b, m, deterministic)
}
func (m *AcceptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcceptRequest.Merge(m, src)
}
func (m *AcceptRequest) XXX_Size() int {
	return xxx_messageInfo_AcceptRequest.Size(m)
}
func (m *AcceptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AcceptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AcceptRequest proto.InternalMessageInfo

func (m *AcceptRequest) GetProcessId() int64 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *AcceptRequest) GetProposeNumber() string {
	if m != nil {
		return m.ProposeNumber
	}
	return ""
}

func (m *AcceptRequest) GetMaxKey() string {
	if m != nil {
		return m.MaxKey
	}
	return ""
}

func (m *AcceptRequest) GetMaxValue() string {
	if m != nil {
		return m.MaxValue
	}
	return ""
}

type Accepted struct {
	ProposeNumber         string   `protobuf:"bytes,1,opt,name=propose_number,json=proposeNumber,proto3" json:"propose_number,omitempty"`
	AcceptedProposeNumber string   `protobuf:"bytes,2,opt,name=accepted_propose_number,json=acceptedProposeNumber,proto3" json:"accepted_propose_number,omitempty"`
	AcceptedKey           string   `protobuf:"bytes,3,opt,name=accepted_key,json=acceptedKey,proto3" json:"accepted_key,omitempty"`
	AcceptedValue         string   `protobuf:"bytes,4,opt,name=accepted_value,json=acceptedValue,proto3" json:"accepted_value,omitempty"`
	IsAccepted            bool     `protobuf:"varint,5,opt,name=isAccepted,proto3" json:"isAccepted,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Accepted) Reset()         { *m = Accepted{} }
func (m *Accepted) String() string { return proto.CompactTextString(m) }
func (*Accepted) ProtoMessage()    {}
func (*Accepted) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66e7684f91bf7b3, []int{3}
}

func (m *Accepted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Accepted.Unmarshal(m, b)
}
func (m *Accepted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Accepted.Marshal(b, m, deterministic)
}
func (m *Accepted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Accepted.Merge(m, src)
}
func (m *Accepted) XXX_Size() int {
	return xxx_messageInfo_Accepted.Size(m)
}
func (m *Accepted) XXX_DiscardUnknown() {
	xxx_messageInfo_Accepted.DiscardUnknown(m)
}

var xxx_messageInfo_Accepted proto.InternalMessageInfo

func (m *Accepted) GetProposeNumber() string {
	if m != nil {
		return m.ProposeNumber
	}
	return ""
}

func (m *Accepted) GetAcceptedProposeNumber() string {
	if m != nil {
		return m.AcceptedProposeNumber
	}
	return ""
}

func (m *Accepted) GetAcceptedKey() string {
	if m != nil {
		return m.AcceptedKey
	}
	return ""
}

func (m *Accepted) GetAcceptedValue() string {
	if m != nil {
		return m.AcceptedValue
	}
	return ""
}

func (m *Accepted) GetIsAccepted() bool {
	if m != nil {
		return m.IsAccepted
	}
	return false
}

type DecidedRequest struct {
	ProcessId            int64    `protobuf:"varint,1,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	ProposeNumber        string   `protobuf:"bytes,2,opt,name=propose_number,json=proposeNumber,proto3" json:"propose_number,omitempty"`
	MaxKey               string   `protobuf:"bytes,3,opt,name=max_key,json=maxKey,proto3" json:"max_key,omitempty"`
	MaxValue             string   `protobuf:"bytes,4,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
	OperationId          int64    `protobuf:"varint,5,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecidedRequest) Reset()         { *m = DecidedRequest{} }
func (m *DecidedRequest) String() string { return proto.CompactTextString(m) }
func (*DecidedRequest) ProtoMessage()    {}
func (*DecidedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66e7684f91bf7b3, []int{4}
}

func (m *DecidedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecidedRequest.Unmarshal(m, b)
}
func (m *DecidedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecidedRequest.Marshal(b, m, deterministic)
}
func (m *DecidedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecidedRequest.Merge(m, src)
}
func (m *DecidedRequest) XXX_Size() int {
	return xxx_messageInfo_DecidedRequest.Size(m)
}
func (m *DecidedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecidedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecidedRequest proto.InternalMessageInfo

func (m *DecidedRequest) GetProcessId() int64 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *DecidedRequest) GetProposeNumber() string {
	if m != nil {
		return m.ProposeNumber
	}
	return ""
}

func (m *DecidedRequest) GetMaxKey() string {
	if m != nil {
		return m.MaxKey
	}
	return ""
}

func (m *DecidedRequest) GetMaxValue() string {
	if m != nil {
		return m.MaxValue
	}
	return ""
}

func (m *DecidedRequest) GetOperationId() int64 {
	if m != nil {
		return m.OperationId
	}
	return 0
}

type DecidedResponse struct {
	IsMarkedDecided      bool     `protobuf:"varint,1,opt,name=isMarkedDecided,proto3" json:"isMarkedDecided,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecidedResponse) Reset()         { *m = DecidedResponse{} }
func (m *DecidedResponse) String() string { return proto.CompactTextString(m) }
func (*DecidedResponse) ProtoMessage()    {}
func (*DecidedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66e7684f91bf7b3, []int{5}
}

func (m *DecidedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecidedResponse.Unmarshal(m, b)
}
func (m *DecidedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecidedResponse.Marshal(b, m, deterministic)
}
func (m *DecidedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecidedResponse.Merge(m, src)
}
func (m *DecidedResponse) XXX_Size() int {
	return xxx_messageInfo_DecidedResponse.Size(m)
}
func (m *DecidedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecidedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecidedResponse proto.InternalMessageInfo

func (m *DecidedResponse) GetIsMarkedDecided() bool {
	if m != nil {
		return m.IsMarkedDecided
	}
	return false
}

func init() {
	proto.RegisterType((*PrepareRequest)(nil), "proto.PrepareRequest")
	proto.RegisterType((*Promise)(nil), "proto.Promise")
	proto.RegisterType((*AcceptRequest)(nil), "proto.AcceptRequest")
	proto.RegisterType((*Accepted)(nil), "proto.Accepted")
	proto.RegisterType((*DecidedRequest)(nil), "proto.DecidedRequest")
	proto.RegisterType((*DecidedResponse)(nil), "proto.DecidedResponse")
}

func init() {
	proto.RegisterFile("paxosservice.proto", fileDescriptor_d66e7684f91bf7b3)
}

var fileDescriptor_d66e7684f91bf7b3 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0xf5, 0xc6, 0xb1, 0x2d, 0x8f, 0x12, 0x19, 0x96, 0xa6, 0x11, 0x2e, 0x6d, 0xdd, 0x85, 0x82,
	0x4e, 0x81, 0xa6, 0xd0, 0x4b, 0xa1, 0x10, 0xe8, 0x25, 0x94, 0x16, 0xa1, 0x42, 0xae, 0x66, 0xa3,
	0x9d, 0xc3, 0x92, 0xca, 0xbb, 0xdd, 0x95, 0x83, 0xf2, 0x0d, 0xfd, 0x98, 0x9e, 0xfa, 0x25, 0xfd,
	0x97, 0x9e, 0xcb, 0xee, 0x4a, 0x02, 0x39, 0x39, 0xc4, 0xd0, 0x43, 0x4f, 0x82, 0x37, 0xb3, 0x9a,
	0xf7, 0xde, 0xcc, 0x03, 0xaa, 0x79, 0xa3, 0xac, 0x45, 0x73, 0x2b, 0x4b, 0x3c, 0xd3, 0x46, 0xd5,
	0x8a, 0x4e, 0xfc, 0x87, 0x5d, 0x41, 0x92, 0x1b, 0xd4, 0xdc, 0x60, 0x81, 0xdf, 0xb7, 0x68, 0x6b,
	0xfa, 0x1c, 0x40, 0x1b, 0x55, 0xa2, 0xb5, 0x6b, 0x29, 0x52, 0xb2, 0x22, 0xd9, 0xb8, 0x98, 0xb7,
	0xc8, 0xa5, 0xa0, 0xaf, 0x21, 0xd1, 0x46, 0x69, 0x65, 0x71, 0xbd, 0xd9, 0x56, 0xd7, 0x68, 0xd2,
	0x83, 0x15, 0xc9, 0xe6, 0xc5, 0x71, 0x8b, 0x7e, 0xf1, 0x20, 0xfb, 0x43, 0x60, 0x96, 0x1b, 0x55,
	0x49, 0x8b, 0x8f, 0x7c, 0x42, 0xdf, 0xc1, 0x29, 0x2f, 0x4b, 0xd4, 0x35, 0x8a, 0xf5, 0x4e, 0xff,
	0xd8, 0xf7, 0x9f, 0x74, 0xe5, 0x7c, 0xf0, 0xee, 0x15, 0x1c, 0xf5, 0xef, 0x6e, 0xf0, 0x2e, 0x3d,
	0xf4, 0xcd, 0x71, 0x87, 0x7d, 0xc2, 0x3b, 0xc7, 0xa0, 0x6f, 0xb9, 0xe5, 0xdf, 0xb6, 0x98, 0x4e,
	0x02, 0x83, 0x0e, 0xbd, 0x72, 0x20, 0x7d, 0x09, 0xb1, 0xb4, 0x6e, 0xb6, 0xa3, 0x2d, 0xd2, 0xe9,
	0x8a, 0x64, 0x51, 0x01, 0xd2, 0xb6, 0x42, 0x84, 0x1b, 0xa5, 0x34, 0x1a, 0x5e, 0x4b, 0xb5, 0x71,
	0xee, 0xcc, 0xbc, 0x3b, 0x71, 0x8f, 0x5d, 0x0a, 0xf6, 0x83, 0xc0, 0xf1, 0x85, 0xff, 0xeb, 0x3f,
	0x35, 0x94, 0x9e, 0xc2, 0xac, 0xe2, 0x8d, 0x17, 0x18, 0xdc, 0x98, 0x56, 0xbc, 0x71, 0xda, 0x9e,
	0xc1, 0xdc, 0x15, 0x82, 0xac, 0xa0, 0x3d, 0xaa, 0x78, 0xe3, 0x15, 0xb1, 0xdf, 0x04, 0xa2, 0x8b,
	0x56, 0xe3, 0x03, 0x93, 0xc8, 0x9e, 0x7b, 0x38, 0xd8, 0x67, 0x0f, 0xe3, 0xc7, 0xec, 0xe1, 0xf0,
	0xa1, 0x3d, 0xbc, 0x00, 0x90, 0xb6, 0xa3, 0xed, 0x57, 0xe5, 0xd7, 0xd0, 0x21, 0xec, 0x27, 0x81,
	0xe4, 0x23, 0x96, 0x52, 0xa0, 0xf8, 0x0f, 0x4c, 0xbe, 0x77, 0x15, 0x93, 0xfb, 0x57, 0xf1, 0x1e,
	0x16, 0x3d, 0x61, 0xab, 0xd5, 0xc6, 0x22, 0xcd, 0x60, 0x21, 0xed, 0x67, 0x6e, 0x6e, 0x50, 0xb4,
	0x25, 0x4f, 0x3b, 0x2a, 0x76, 0xe1, 0xf3, 0x5f, 0x04, 0x8e, 0x72, 0x97, 0xe0, 0xaf, 0x21, 0xc1,
	0xf4, 0xdc, 0x65, 0xcb, 0x87, 0x96, 0x9e, 0x84, 0x38, 0x9f, 0x0d, 0x43, 0xbc, 0x4c, 0x7a, 0xd8,
	0x5f, 0x2e, 0x1b, 0xd1, 0x37, 0x30, 0x0d, 0xfe, 0xd1, 0x27, 0x6d, 0x6d, 0x70, 0xa5, 0xcb, 0xc5,
	0x00, 0x45, 0xc1, 0x46, 0xf4, 0x03, 0xc4, 0x8e, 0x48, 0x4b, 0xa3, 0x1f, 0x35, 0x74, 0x7e, 0xf9,
	0x74, 0x17, 0x0e, 0xfa, 0xd8, 0xe8, 0x7a, 0xea, 0x0b, 0x6f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0xd1, 0xd9, 0x3f, 0x7f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PaxosServiceClient is the client API for PaxosService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PaxosServiceClient interface {
	Prepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*Promise, error)
	Accept(ctx context.Context, in *AcceptRequest, opts ...grpc.CallOption) (*Accepted, error)
	MarkDecided(ctx context.Context, in *DecidedRequest, opts ...grpc.CallOption) (*DecidedResponse, error)
}

type paxosServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaxosServiceClient(cc grpc.ClientConnInterface) PaxosServiceClient {
	return &paxosServiceClient{cc}
}

func (c *paxosServiceClient) Prepare(ctx context.Context, in *PrepareRequest, opts ...grpc.CallOption) (*Promise, error) {
	out := new(Promise)
	err := c.cc.Invoke(ctx, "/proto.PaxosService/Prepare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosServiceClient) Accept(ctx context.Context, in *AcceptRequest, opts ...grpc.CallOption) (*Accepted, error) {
	out := new(Accepted)
	err := c.cc.Invoke(ctx, "/proto.PaxosService/Accept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paxosServiceClient) MarkDecided(ctx context.Context, in *DecidedRequest, opts ...grpc.CallOption) (*DecidedResponse, error) {
	out := new(DecidedResponse)
	err := c.cc.Invoke(ctx, "/proto.PaxosService/MarkDecided", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaxosServiceServer is the server API for PaxosService service.
type PaxosServiceServer interface {
	Prepare(context.Context, *PrepareRequest) (*Promise, error)
	Accept(context.Context, *AcceptRequest) (*Accepted, error)
	MarkDecided(context.Context, *DecidedRequest) (*DecidedResponse, error)
}

// UnimplementedPaxosServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPaxosServiceServer struct {
}

func (*UnimplementedPaxosServiceServer) Prepare(ctx context.Context, req *PrepareRequest) (*Promise, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (*UnimplementedPaxosServiceServer) Accept(ctx context.Context, req *AcceptRequest) (*Accepted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Accept not implemented")
}
func (*UnimplementedPaxosServiceServer) MarkDecided(ctx context.Context, req *DecidedRequest) (*DecidedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkDecided not implemented")
}

func RegisterPaxosServiceServer(s *grpc.Server, srv PaxosServiceServer) {
	s.RegisterService(&_PaxosService_serviceDesc, srv)
}

func _PaxosService_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServiceServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PaxosService/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServiceServer).Prepare(ctx, req.(*PrepareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaxosService_Accept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServiceServer).Accept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PaxosService/Accept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServiceServer).Accept(ctx, req.(*AcceptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaxosService_MarkDecided_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecidedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaxosServiceServer).MarkDecided(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PaxosService/MarkDecided",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaxosServiceServer).MarkDecided(ctx, req.(*DecidedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaxosService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PaxosService",
	HandlerType: (*PaxosServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _PaxosService_Prepare_Handler,
		},
		{
			MethodName: "Accept",
			Handler:    _PaxosService_Accept_Handler,
		},
		{
			MethodName: "MarkDecided",
			Handler:    _PaxosService_MarkDecided_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paxosservice.proto",
}
