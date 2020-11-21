// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

package sessionService

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

type CsrfTokenInput struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	TimeStamp            int64    `protobuf:"varint,2,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsrfTokenInput) Reset()         { *m = CsrfTokenInput{} }
func (m *CsrfTokenInput) String() string { return proto.CompactTextString(m) }
func (*CsrfTokenInput) ProtoMessage()    {}
func (*CsrfTokenInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{2}
}

func (m *CsrfTokenInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsrfTokenInput.Unmarshal(m, b)
}
func (m *CsrfTokenInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsrfTokenInput.Marshal(b, m, deterministic)
}
func (m *CsrfTokenInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrfTokenInput.Merge(m, src)
}
func (m *CsrfTokenInput) XXX_Size() int {
	return xxx_messageInfo_CsrfTokenInput.Size(m)
}
func (m *CsrfTokenInput) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrfTokenInput.DiscardUnknown(m)
}

var xxx_messageInfo_CsrfTokenInput proto.InternalMessageInfo

func (m *CsrfTokenInput) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *CsrfTokenInput) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

type CsrfToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsrfToken) Reset()         { *m = CsrfToken{} }
func (m *CsrfToken) String() string { return proto.CompactTextString(m) }
func (*CsrfToken) ProtoMessage()    {}
func (*CsrfToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{3}
}

func (m *CsrfToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsrfToken.Unmarshal(m, b)
}
func (m *CsrfToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsrfToken.Marshal(b, m, deterministic)
}
func (m *CsrfToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrfToken.Merge(m, src)
}
func (m *CsrfToken) XXX_Size() int {
	return xxx_messageInfo_CsrfToken.Size(m)
}
func (m *CsrfToken) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrfToken.DiscardUnknown(m)
}

var xxx_messageInfo_CsrfToken proto.InternalMessageInfo

func (m *CsrfToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CsrfTokenCheck struct {
	SessionID            string   `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsrfTokenCheck) Reset()         { *m = CsrfTokenCheck{} }
func (m *CsrfTokenCheck) String() string { return proto.CompactTextString(m) }
func (*CsrfTokenCheck) ProtoMessage()    {}
func (*CsrfTokenCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{4}
}

func (m *CsrfTokenCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsrfTokenCheck.Unmarshal(m, b)
}
func (m *CsrfTokenCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsrfTokenCheck.Marshal(b, m, deterministic)
}
func (m *CsrfTokenCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrfTokenCheck.Merge(m, src)
}
func (m *CsrfTokenCheck) XXX_Size() int {
	return xxx_messageInfo_CsrfTokenCheck.Size(m)
}
func (m *CsrfTokenCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrfTokenCheck.DiscardUnknown(m)
}

var xxx_messageInfo_CsrfTokenCheck proto.InternalMessageInfo

func (m *CsrfTokenCheck) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *CsrfTokenCheck) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CheckResult struct {
	Result               bool     `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckResult) Reset()         { *m = CheckResult{} }
func (m *CheckResult) String() string { return proto.CompactTextString(m) }
func (*CheckResult) ProtoMessage()    {}
func (*CheckResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a6be1b361fa6f14, []int{5}
}

func (m *CheckResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResult.Unmarshal(m, b)
}
func (m *CheckResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResult.Marshal(b, m, deterministic)
}
func (m *CheckResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResult.Merge(m, src)
}
func (m *CheckResult) XXX_Size() int {
	return xxx_messageInfo_CheckResult.Size(m)
}
func (m *CheckResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResult.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResult proto.InternalMessageInfo

func (m *CheckResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
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
	return fileDescriptor_3a6be1b361fa6f14, []int{6}
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
	proto.RegisterType((*SessionID)(nil), "sessionService.SessionID")
	proto.RegisterType((*UserID)(nil), "sessionService.UserID")
	proto.RegisterType((*CsrfTokenInput)(nil), "sessionService.CsrfTokenInput")
	proto.RegisterType((*CsrfToken)(nil), "sessionService.CsrfToken")
	proto.RegisterType((*CsrfTokenCheck)(nil), "sessionService.CsrfTokenCheck")
	proto.RegisterType((*CheckResult)(nil), "sessionService.CheckResult")
	proto.RegisterType((*Empty)(nil), "sessionService.Empty")
}

func init() {
	proto.RegisterFile("session.proto", fileDescriptor_3a6be1b361fa6f14)
}

var fileDescriptor_3a6be1b361fa6f14 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xdb, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x7b, 0xa0, 0xd5, 0x8e, 0x24, 0xc2, 0x52, 0x4b, 0x8d, 0x22, 0x75, 0x41, 0xd0, 0x9b,
	0x5c, 0xe8, 0x13, 0x68, 0x52, 0x24, 0x20, 0x0a, 0x49, 0x7d, 0x80, 0x58, 0x46, 0x1a, 0xda, 0x1c,
	0xd8, 0xdd, 0x08, 0xf5, 0x0d, 0x7c, 0x6b, 0x93, 0x4d, 0x9a, 0x4d, 0x53, 0x5a, 0xbd, 0x9b, 0xc3,
	0xbf, 0xdf, 0x64, 0xfe, 0x09, 0x68, 0x1c, 0x39, 0x0f, 0xe2, 0xc8, 0x4c, 0x58, 0x2c, 0x62, 0xa2,
	0x97, 0xa9, 0x87, 0xec, 0x2b, 0x98, 0x23, 0xbd, 0x83, 0x81, 0x57, 0x54, 0x1c, 0x9b, 0x5c, 0xd6,
	0x92, 0x71, 0x7b, 0xd2, 0xbe, 0x1d, 0xb8, 0xaa, 0x40, 0x27, 0xd0, 0x7f, 0xe7, 0xc8, 0x32, 0xdd,
	0x68, 0x13, 0x49, 0x51, 0xd7, 0x2d, 0x33, 0xfa, 0x02, 0xba, 0xc5, 0xd9, 0xe7, 0x2c, 0x5e, 0x62,
	0xe4, 0x44, 0x49, 0x2a, 0x0e, 0x13, 0xf3, 0xee, 0x2c, 0x08, 0xd1, 0x13, 0x7e, 0x98, 0x8c, 0x3b,
	0x12, 0xa5, 0x0a, 0xf4, 0x1a, 0x06, 0x15, 0x8d, 0x0c, 0xa1, 0x27, 0x83, 0x12, 0x52, 0x24, 0xd4,
	0xae, 0x0d, 0xb4, 0x16, 0x38, 0x5f, 0xfe, 0x31, 0xb0, 0xa2, 0x74, 0xea, 0x94, 0x1b, 0x38, 0x91,
	0x8f, 0x5d, 0xe4, 0xe9, 0x4a, 0xe4, 0xdb, 0x15, 0x91, 0x7c, 0x7f, 0xec, 0x96, 0x19, 0x3d, 0x82,
	0xde, 0x34, 0x4c, 0xc4, 0xfa, 0xfe, 0xa7, 0x0b, 0xc3, 0xc7, 0x54, 0x2c, 0x62, 0x16, 0x7c, 0xfb,
	0x42, 0x99, 0x49, 0x6c, 0xd0, 0x2c, 0x86, 0xbe, 0xc0, 0x72, 0x22, 0x19, 0x99, 0xdb, 0x76, 0x9b,
	0x85, 0x51, 0xc6, 0x79, 0xb3, 0xae, 0x5c, 0x6e, 0x91, 0x29, 0xe8, 0xcf, 0x28, 0x1c, 0xfb, 0x69,
	0xbd, 0xc1, 0xec, 0x97, 0x1b, 0x7b, 0x26, 0x64, 0x18, 0x0b, 0x34, 0x1b, 0x57, 0xa8, 0x3e, 0xe6,
	0x00, 0xe5, 0xac, 0xd9, 0x92, 0x8b, 0x66, 0x90, 0x57, 0x38, 0x2d, 0x36, 0x52, 0x97, 0xb8, 0x6a,
	0x6a, 0xb7, 0x4f, 0xbe, 0xbb, 0x5b, 0xd5, 0xcf, 0x78, 0x6f, 0xd9, 0xc1, 0x72, 0xab, 0xff, 0x83,
	0x93, 0x42, 0xe3, 0x62, 0xa7, 0xaf, 0x4e, 0x45, 0x5b, 0x1f, 0x7d, 0xf9, 0x5b, 0x3f, 0xfc, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xae, 0x2e, 0x94, 0xc4, 0xe7, 0x02, 0x00, 0x00,
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
	CreateCsrfToken(ctx context.Context, in *CsrfTokenInput, opts ...grpc.CallOption) (*CsrfToken, error)
	CheckCsrfToken(ctx context.Context, in *CsrfTokenCheck, opts ...grpc.CallOption) (*CheckResult, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) CreateSession(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*SessionID, error) {
	out := new(SessionID)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetIDBySession(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*UserID, error) {
	out := new(UserID)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/GetIDBySession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) DeleteSession(ctx context.Context, in *SessionID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) CreateCsrfToken(ctx context.Context, in *CsrfTokenInput, opts ...grpc.CallOption) (*CsrfToken, error) {
	out := new(CsrfToken)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/CreateCsrfToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) CheckCsrfToken(ctx context.Context, in *CsrfTokenCheck, opts ...grpc.CallOption) (*CheckResult, error) {
	out := new(CheckResult)
	err := c.cc.Invoke(ctx, "/sessionService.AuthorizationService/CheckCsrfToken", in, out, opts...)
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
	CreateCsrfToken(context.Context, *CsrfTokenInput) (*CsrfToken, error)
	CheckCsrfToken(context.Context, *CsrfTokenCheck) (*CheckResult, error)
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
func (*UnimplementedAuthorizationServiceServer) CreateCsrfToken(ctx context.Context, req *CsrfTokenInput) (*CsrfToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCsrfToken not implemented")
}
func (*UnimplementedAuthorizationServiceServer) CheckCsrfToken(ctx context.Context, req *CsrfTokenCheck) (*CheckResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCsrfToken not implemented")
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
		FullMethod: "/sessionService.AuthorizationService/CreateSession",
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
		FullMethod: "/sessionService.AuthorizationService/GetIDBySession",
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
		FullMethod: "/sessionService.AuthorizationService/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).DeleteSession(ctx, req.(*SessionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_CreateCsrfToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CsrfTokenInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateCsrfToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionService.AuthorizationService/CreateCsrfToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateCsrfToken(ctx, req.(*CsrfTokenInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_CheckCsrfToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CsrfTokenCheck)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CheckCsrfToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sessionService.AuthorizationService/CheckCsrfToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CheckCsrfToken(ctx, req.(*CsrfTokenCheck))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorizationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sessionService.AuthorizationService",
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
		{
			MethodName: "CreateCsrfToken",
			Handler:    _AuthorizationService_CreateCsrfToken_Handler,
		},
		{
			MethodName: "CheckCsrfToken",
			Handler:    _AuthorizationService_CheckCsrfToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "session.proto",
}
