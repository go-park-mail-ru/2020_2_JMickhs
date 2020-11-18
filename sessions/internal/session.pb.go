// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

package session

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SessionID struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionID) Reset()         { *m = SessionID{} }
func (m *SessionID) String() string { return proto.CompactTextString(m) }
func (*SessionID) ProtoMessage()    {}
func (*SessionID) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{0}
}

func (m *SessionID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionID.Unmarshal(m, b)
}
func (m *SessionID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionID.Marshal(b, m, deterministic)
}
func (m *SessionID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionID.Merge(m, src)
}
func (m *SessionID) XXX_Size() int {
	return xxx_messageInfo_SessionID.Size(m)
}
func (m *SessionID) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionID.DiscardUnknown(m)
}

var xxx_messageInfo_SessionID proto.InternalMessageInfo

func (m *SessionID) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

type UserID struct {
	UserID               int64    `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{1}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{2}
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

func init() {
	proto.RegisterType((*SessionID)(nil), "session.SessionID")
	proto.RegisterType((*UserID)(nil), "session.UserID")
	proto.RegisterType((*Empty)(nil), "session.Empty")
}

func init() {
	proto.RegisterFile("session.proto", fileDescriptor_3a6be1b361fa6f14)
}

var fileDescriptor_3a6be1b361fa6f14 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2e,
	0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x34, 0xb9,
	0x38, 0x83, 0x21, 0x4c, 0x4f, 0x17, 0x21, 0x19, 0x24, 0x8e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x42, 0x40, 0x49, 0x81, 0x8b, 0x2d, 0xb4, 0x38, 0xb5, 0x08, 0xa8, 0x4e, 0x0c, 0xc6, 0x02,
	0x2b, 0x62, 0x0e, 0x82, 0xf2, 0x94, 0xd8, 0xb9, 0x58, 0x5d, 0x73, 0x0b, 0x4a, 0x2a, 0x8d, 0xf6,
	0x31, 0x72, 0x89, 0x38, 0x96, 0x96, 0x64, 0xe4, 0x17, 0x65, 0x56, 0x25, 0x96, 0x00, 0xb5, 0x07,
	0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x99, 0x71, 0xf1, 0x3a, 0x17, 0xa5, 0x26, 0x96, 0xa4,
	0x42, 0x8d, 0x15, 0xe2, 0xd7, 0x83, 0x39, 0x0c, 0x62, 0x86, 0x94, 0x10, 0x5c, 0x00, 0x61, 0x33,
	0x83, 0x90, 0x39, 0x17, 0x9f, 0x7b, 0x6a, 0x89, 0xa7, 0x8b, 0x53, 0x25, 0x4c, 0x23, 0x16, 0x75,
	0x52, 0xe8, 0x86, 0x01, 0x35, 0x9a, 0x72, 0xf1, 0xba, 0xa4, 0xe6, 0xa4, 0x22, 0x2c, 0xc4, 0xa6,
	0x8f, 0x0f, 0x2e, 0x06, 0x76, 0xbe, 0x12, 0x43, 0x12, 0x1b, 0x38, 0x98, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x80, 0x20, 0x59, 0x87, 0x37, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	CreateSession(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*SessionID, error)
	GetIDBySession(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*UserID, error)
	DeleteSession(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Empty, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) CreateSession(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*SessionID, error) {
	out := new(SessionID)
	err := c.cc.Invoke(ctx, "/session.AuthorizationService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetIDBySession(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/session.AuthorizationService/GetIDBySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) DeleteSession(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/session.AuthorizationService/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
type AuthorizationServiceServer interface {
	CreateSession(context.Context, *UserID) (*SessionID, error)
	GetIDBySession(context.Context, *SessionID) (*UserID, error)
	DeleteSession(context.Context, *SessionID) (*Empty, error)
}

// UnimplementedAuthorizationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (*UnimplementedAuthorizationServiceServer) CreateSession(ctx context.Context, req *UserID) (*SessionID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (*UnimplementedAuthorizationServiceServer) GetIDBySession(ctx context.Context, req *SessionID) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIDBySession not implemented")
}
func (*UnimplementedAuthorizationServiceServer) DeleteSession(ctx context.Context, req *SessionID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}

func RegisterAuthorizationServiceServer(s *grpc.Server, srv AuthorizationServiceServer) {
	s.RegisterService(&_AuthorizationService_serviceDesc, srv)
}

func _AuthorizationService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthorizationService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateSession(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetIDBySession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetIDBySession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthorizationService/GetIDBySession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetIDBySession(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/session.AuthorizationService/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).DeleteSession(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorizationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "session.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _AuthorizationService_CreateSession_Handler,
		},
		{
			MethodName: "GetIDBySession",
			Handler:    _AuthorizationService_GetIDBySession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _AuthorizationService_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
