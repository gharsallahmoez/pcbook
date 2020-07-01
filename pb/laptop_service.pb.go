// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateLaptopRequest struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopRequest) Reset()         { *m = CreateLaptopRequest{} }
func (m *CreateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopRequest) ProtoMessage()    {}
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{0}
}

func (m *CreateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopRequest.Unmarshal(m, b)
}
func (m *CreateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *CreateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopRequest.Merge(m, src)
}
func (m *CreateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopRequest.Size(m)
}
func (m *CreateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopRequest proto.InternalMessageInfo

func (m *CreateLaptopRequest) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopResponse) Reset()         { *m = CreateLaptopResponse{} }
func (m *CreateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopResponse) ProtoMessage()    {}
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{1}
}

func (m *CreateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopResponse.Unmarshal(m, b)
}
func (m *CreateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *CreateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopResponse.Merge(m, src)
}
func (m *CreateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopResponse.Size(m)
}
func (m *CreateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopResponse proto.InternalMessageInfo

func (m *CreateLaptopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchLaptopRequest struct {
	Filter               *Filter  `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchLaptopRequest) Reset()         { *m = SearchLaptopRequest{} }
func (m *SearchLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*SearchLaptopRequest) ProtoMessage()    {}
func (*SearchLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{2}
}

func (m *SearchLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchLaptopRequest.Unmarshal(m, b)
}
func (m *SearchLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchLaptopRequest.Marshal(b, m, deterministic)
}
func (m *SearchLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchLaptopRequest.Merge(m, src)
}
func (m *SearchLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_SearchLaptopRequest.Size(m)
}
func (m *SearchLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchLaptopRequest proto.InternalMessageInfo

func (m *SearchLaptopRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type SearchLaptopResponse struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchLaptopResponse) Reset()         { *m = SearchLaptopResponse{} }
func (m *SearchLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*SearchLaptopResponse) ProtoMessage()    {}
func (*SearchLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{3}
}

func (m *SearchLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchLaptopResponse.Unmarshal(m, b)
}
func (m *SearchLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchLaptopResponse.Marshal(b, m, deterministic)
}
func (m *SearchLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchLaptopResponse.Merge(m, src)
}
func (m *SearchLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_SearchLaptopResponse.Size(m)
}
func (m *SearchLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchLaptopResponse proto.InternalMessageInfo

func (m *SearchLaptopResponse) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateLaptopRequest)(nil), "techschool.pcbook.CreateLaptopRequest")
	proto.RegisterType((*CreateLaptopResponse)(nil), "techschool.pcbook.CreateLaptopResponse")
	proto.RegisterType((*SearchLaptopRequest)(nil), "techschool.pcbook.SearchLaptopRequest")
	proto.RegisterType((*SearchLaptopResponse)(nil), "techschool.pcbook.SearchLaptopResponse")
}

func init() {
	proto.RegisterFile("laptop_service.proto", fileDescriptor_240c60d9fb227e71)
}

var fileDescriptor_240c60d9fb227e71 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xbb, 0x22, 0x05, 0xc7, 0x2a, 0x18, 0xf7, 0xa0, 0x8b, 0xa0, 0xec, 0xa1, 0xea, 0x25,
	0xd5, 0x7a, 0xf4, 0xa6, 0x20, 0x0a, 0x1e, 0x64, 0x7b, 0xf3, 0x22, 0xd9, 0x38, 0xee, 0x06, 0xb7,
	0x3b, 0x31, 0x89, 0xfe, 0x55, 0xff, 0x8e, 0x75, 0x13, 0xa1, 0xb5, 0x41, 0xe9, 0x2d, 0x4c, 0xde,
	0xbc, 0xef, 0xe5, 0x11, 0x48, 0x1b, 0xa1, 0x1d, 0xe9, 0x27, 0x8b, 0xe6, 0x43, 0x49, 0xe4, 0xda,
	0x90, 0x23, 0xb6, 0xe3, 0x50, 0xd6, 0x56, 0xd6, 0x44, 0x0d, 0xd7, 0xb2, 0x24, 0x7a, 0xcd, 0x7e,
	0x84, 0x53, 0xb4, 0x56, 0x54, 0x41, 0x98, 0xa5, 0x2f, 0xaa, 0x71, 0x68, 0x7e, 0x4d, 0x0f, 0x2a,
	0xa2, 0xaa, 0xc1, 0x91, 0xd0, 0x6a, 0x24, 0xda, 0x96, 0x9c, 0x70, 0x8a, 0x5a, 0xeb, 0x6f, 0xf3,
	0x5b, 0xd8, 0xbd, 0x36, 0x28, 0x1c, 0xde, 0x77, 0x8e, 0x05, 0xbe, 0xbd, 0xa3, 0x75, 0xec, 0x1c,
	0xfa, 0x1e, 0xb1, 0x97, 0x1c, 0x25, 0x27, 0x9b, 0xe3, 0x7d, 0xbe, 0x14, 0x82, 0x87, 0x8d, 0x20,
	0xcc, 0x87, 0x90, 0x2e, 0x3a, 0x59, 0x3d, 0xc3, 0x20, 0xdb, 0x86, 0x35, 0xf5, 0xdc, 0xd9, 0x6c,
	0x14, 0xb3, 0xd3, 0x37, 0x71, 0x82, 0xc2, 0xc8, 0x7a, 0x89, 0xe8, 0xe3, 0xff, 0x41, 0xbc, 0xe9,
	0x04, 0x45, 0x10, 0xe6, 0x77, 0x90, 0x2e, 0x3a, 0x05, 0xe2, 0xea, 0xe1, 0xc7, 0x9f, 0x09, 0x6c,
	0xf9, 0xd1, 0xc4, 0x77, 0xcf, 0x04, 0x0c, 0xe6, 0x9f, 0xc3, 0x86, 0x11, 0x93, 0x48, 0x73, 0xd9,
	0xf1, 0xbf, 0x3a, 0x9f, 0x32, 0xef, 0x31, 0x09, 0x83, 0xf9, 0xfc, 0x51, 0x44, 0xa4, 0xaa, 0x28,
	0x22, 0x56, 0x44, 0xde, 0x3b, 0x4b, 0xae, 0x4e, 0xe1, 0x50, 0xd2, 0x94, 0x57, 0xca, 0x35, 0xa2,
	0x8c, 0x2c, 0xea, 0xf2, 0x21, 0x79, 0x5c, 0xe7, 0x97, 0xba, 0x2c, 0xfb, 0xdd, 0x97, 0xb8, 0xf8,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xe1, 0x6b, 0x87, 0x87, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
	SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (LaptopService_SearchLaptopClient, error)
}

type laptopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaptopServiceClient(cc grpc.ClientConnInterface) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/techschool.pcbook.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laptopServiceClient) SearchLaptop(ctx context.Context, in *SearchLaptopRequest, opts ...grpc.CallOption) (LaptopService_SearchLaptopClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[0], "/techschool.pcbook.LaptopService/SearchLaptop", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceSearchLaptopClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LaptopService_SearchLaptopClient interface {
	Recv() (*SearchLaptopResponse, error)
	grpc.ClientStream
}

type laptopServiceSearchLaptopClient struct {
	grpc.ClientStream
}

func (x *laptopServiceSearchLaptopClient) Recv() (*SearchLaptopResponse, error) {
	m := new(SearchLaptopResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LaptopServiceServer is the server API for LaptopService service.
type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
	SearchLaptop(*SearchLaptopRequest, LaptopService_SearchLaptopServer) error
}

// UnimplementedLaptopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLaptopServiceServer struct {
}

func (*UnimplementedLaptopServiceServer) CreateLaptop(ctx context.Context, req *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}
func (*UnimplementedLaptopServiceServer) SearchLaptop(req *SearchLaptopRequest, srv LaptopService_SearchLaptopServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchLaptop not implemented")
}

func RegisterLaptopServiceServer(s *grpc.Server, srv LaptopServiceServer) {
	s.RegisterService(&_LaptopService_serviceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/techschool.pcbook.LaptopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaptopService_SearchLaptop_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchLaptopRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaptopServiceServer).SearchLaptop(m, &laptopServiceSearchLaptopServer{stream})
}

type LaptopService_SearchLaptopServer interface {
	Send(*SearchLaptopResponse) error
	grpc.ServerStream
}

type laptopServiceSearchLaptopServer struct {
	grpc.ServerStream
}

func (x *laptopServiceSearchLaptopServer) Send(m *SearchLaptopResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _LaptopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "techschool.pcbook.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchLaptop",
			Handler:       _LaptopService_SearchLaptop_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "laptop_service.proto",
}
