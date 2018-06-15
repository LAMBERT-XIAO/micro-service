// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package go_micro_srv_confignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 货轮承运的一批货物
type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_e34f2c6bc7032886, []int{0}
}
func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (dst *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(dst, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

// 单个集装箱
type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_e34f2c6bc7032886, []int{1}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// 托运结果
type Response struct {
	Created              bool         `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_e34f2c6bc7032886, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.confignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.confignment.Container")
	proto.RegisterType((*Response)(nil), "go.micro.srv.confignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingServiceClient(cc *grpc.ClientConn) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/go.micro.srv.confignment.ShippingService/CreateConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
type ShippingServiceServer interface {
	CreateConsignment(context.Context, *Consignment) (*Response, error)
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.confignment.ShippingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, req.(*Consignment))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.confignment.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ShippingService_CreateConsignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consignment/consignment.proto",
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_consignment_e34f2c6bc7032886)
}

var fileDescriptor_consignment_e34f2c6bc7032886 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x24, 0x7d, 0x67, 0x23, 0x81, 0xf0, 0x01, 0x2c, 0x38, 0x10, 0xa5, 0x42, 0xea, 0x29, 0x48,
	0xe5, 0x13, 0x72, 0x40, 0xbd, 0xba, 0x1f, 0x00, 0xc5, 0x5e, 0xd2, 0x95, 0x88, 0x1d, 0xd9, 0x6e,
	0xf9, 0x36, 0xfe, 0x0e, 0xd5, 0x69, 0xc0, 0x12, 0x2a, 0xe2, 0x96, 0xd9, 0x99, 0xc9, 0xcc, 0xda,
	0x86, 0x79, 0x6b, 0x8d, 0x37, 0x0f, 0xd2, 0x68, 0x47, 0xb5, 0x6e, 0x50, 0xfb, 0xf8, 0xbb, 0x0c,
	0x2c, 0xe3, 0xb5, 0x29, 0x1b, 0x92, 0xd6, 0x94, 0xce, 0xee, 0x4b, 0x69, 0xf4, 0xdb, 0x91, 0x2f,
	0x3e, 0x13, 0xc8, 0xaa, 0x1f, 0x3d, 0x3b, 0x87, 0x01, 0x29, 0x9e, 0xe4, 0xc9, 0x22, 0x15, 0x03,
	0x52, 0x2c, 0x87, 0x4c, 0xa1, 0x93, 0x96, 0x5a, 0x4f, 0x46, 0xf3, 0x41, 0x20, 0xe2, 0x11, 0xbb,
	0x82, 0xc9, 0x07, 0x52, 0xbd, 0xf5, 0x7c, 0x98, 0x27, 0x8b, 0xb1, 0x38, 0x22, 0x56, 0x01, 0x48,
	0xa3, 0xfd, 0x86, 0x34, 0x5a, 0xc7, 0x47, 0xf9, 0x70, 0x91, 0x2d, 0xe7, 0xe5, 0xa9, 0x22, 0x65,
	0xd5, 0x6b, 0x45, 0x64, 0x63, 0xb7, 0x90, 0xee, 0xd1, 0x39, 0x7c, 0x7f, 0x26, 0xc5, 0xc7, 0x21,
	0x7c, 0xd6, 0x0d, 0x56, 0xaa, 0x68, 0x20, 0xfd, 0x76, 0xfd, 0x2a, 0x7e, 0x07, 0x99, 0xdc, 0x39,
	0x6f, 0x1a, 0xb4, 0x07, 0x6f, 0x57, 0x1c, 0xfa, 0xd1, 0x4a, 0x1d, 0x7a, 0x1b, 0x4b, 0x35, 0xe9,
	0xd0, 0x3b, 0x15, 0x47, 0xc4, 0xae, 0x61, 0xba, 0x73, 0x9d, 0x69, 0xd4, 0x11, 0x07, 0x18, 0xe2,
	0x66, 0x02, 0x5d, 0x6b, 0xb4, 0x43, 0xc6, 0x61, 0x2a, 0x2d, 0x6e, 0x3c, 0x76, 0x91, 0x33, 0xd1,
	0x43, 0xf6, 0x04, 0x59, 0x74, 0xfe, 0x21, 0x37, 0x5b, 0xde, 0xff, 0xb9, 0x77, 0x2f, 0x16, 0xb1,
	0x73, 0xe9, 0xe0, 0x62, 0xbd, 0xa5, 0xb6, 0x25, 0x5d, 0xaf, 0xd1, 0xee, 0x49, 0x22, 0x7b, 0x81,
	0xcb, 0x2a, 0xc4, 0xc4, 0x37, 0xf6, 0xbf, 0x7f, 0xdf, 0x14, 0xa7, 0x65, 0xfd, 0x56, 0xc5, 0xd9,
	0xeb, 0x24, 0xbc, 0x97, 0xc7, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0xc9, 0xd7, 0x7d, 0x56,
	0x02, 0x00, 0x00,
}
