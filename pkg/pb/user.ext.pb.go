// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.ext.proto

package pb

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

type SignInReq struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	DeviceId             int64    `protobuf:"varint,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInReq) Reset()         { *m = SignInReq{} }
func (m *SignInReq) String() string { return proto.CompactTextString(m) }
func (*SignInReq) ProtoMessage()    {}
func (*SignInReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{0}
}

func (m *SignInReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInReq.Unmarshal(m, b)
}
func (m *SignInReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInReq.Marshal(b, m, deterministic)
}
func (m *SignInReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInReq.Merge(m, src)
}
func (m *SignInReq) XXX_Size() int {
	return xxx_messageInfo_SignInReq.Size(m)
}
func (m *SignInReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignInReq proto.InternalMessageInfo

func (m *SignInReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SignInReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SignInReq) GetDeviceId() int64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

type SignInResp struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResp) Reset()         { *m = SignInResp{} }
func (m *SignInResp) String() string { return proto.CompactTextString(m) }
func (*SignInResp) ProtoMessage()    {}
func (*SignInResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{1}
}

func (m *SignInResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResp.Unmarshal(m, b)
}
func (m *SignInResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResp.Marshal(b, m, deterministic)
}
func (m *SignInResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResp.Merge(m, src)
}
func (m *SignInResp) XXX_Size() int {
	return xxx_messageInfo_SignInResp.Size(m)
}
func (m *SignInResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResp proto.InternalMessageInfo

func (m *SignInResp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SignInResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetUserReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserReq) Reset()         { *m = GetUserReq{} }
func (m *GetUserReq) String() string { return proto.CompactTextString(m) }
func (*GetUserReq) ProtoMessage()    {}
func (*GetUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{2}
}

func (m *GetUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserReq.Unmarshal(m, b)
}
func (m *GetUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserReq.Marshal(b, m, deterministic)
}
func (m *GetUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserReq.Merge(m, src)
}
func (m *GetUserReq) XXX_Size() int {
	return xxx_messageInfo_GetUserReq.Size(m)
}
func (m *GetUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserReq proto.InternalMessageInfo

type GetUserResp struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResp) Reset()         { *m = GetUserResp{} }
func (m *GetUserResp) String() string { return proto.CompactTextString(m) }
func (*GetUserResp) ProtoMessage()    {}
func (*GetUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{3}
}

func (m *GetUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResp.Unmarshal(m, b)
}
func (m *GetUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResp.Marshal(b, m, deterministic)
}
func (m *GetUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResp.Merge(m, src)
}
func (m *GetUserResp) XXX_Size() int {
	return xxx_messageInfo_GetUserResp.Size(m)
}
func (m *GetUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResp proto.InternalMessageInfo

func (m *GetUserResp) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*SignInReq)(nil), "pb.SignInReq")
	proto.RegisterType((*SignInResp)(nil), "pb.SignInResp")
	proto.RegisterType((*GetUserReq)(nil), "pb.GetUserReq")
	proto.RegisterType((*GetUserResp)(nil), "pb.GetUserResp")
}

func init() { proto.RegisterFile("user.ext.proto", fileDescriptor_f85764c91792e488) }

var fileDescriptor_f85764c91792e488 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x46, 0x49, 0x5b, 0x93, 0x66, 0x52, 0x2b, 0x0c, 0x82, 0x21, 0x7a, 0xa8, 0xb9, 0x58, 0x14,
	0x72, 0xa8, 0x47, 0xcf, 0x22, 0xb9, 0x78, 0x88, 0x78, 0x35, 0x90, 0xee, 0xa0, 0x41, 0xdc, 0x9d,
	0x66, 0xb7, 0xd2, 0x9f, 0x2f, 0x3b, 0x89, 0xd1, 0xde, 0xf6, 0xbd, 0x99, 0x9d, 0x6f, 0x76, 0x61,
	0xb9, 0xb7, 0xd4, 0x15, 0x74, 0x70, 0x05, 0x77, 0xc6, 0x19, 0x9c, 0x70, 0x93, 0xf5, 0xae, 0xd5,
	0x83, 0xcb, 0x6b, 0x88, 0x5f, 0xda, 0x77, 0x5d, 0xea, 0x8a, 0x76, 0x78, 0x0d, 0x0b, 0xfe, 0x30,
	0x9a, 0x6a, 0xbd, 0xff, 0x6a, 0xa8, 0x4b, 0x83, 0x55, 0xb0, 0x8e, 0xab, 0x44, 0xdc, 0xb3, 0x28,
	0x44, 0x98, 0x6d, 0x8d, 0xa2, 0x74, 0x22, 0x25, 0x39, 0xe3, 0x25, 0xc4, 0x8a, 0xbe, 0xdb, 0x2d,
	0xd5, 0xad, 0x4a, 0xa7, 0xab, 0x60, 0x3d, 0xad, 0xe6, 0xbd, 0x28, 0x55, 0xfe, 0x00, 0xf0, 0x1b,
	0x60, 0x19, 0x2f, 0x20, 0xf2, 0x0b, 0xf8, 0xc6, 0x40, 0x1a, 0x43, 0x8f, 0xa5, 0xc2, 0x73, 0x38,
	0x71, 0xe6, 0x93, 0xf4, 0x30, 0xb8, 0x87, 0x7c, 0x01, 0xf0, 0x44, 0xee, 0xd5, 0x52, 0x57, 0xd1,
	0x2e, 0xbf, 0x83, 0x64, 0x24, 0xcb, 0x78, 0x05, 0x33, 0x7f, 0x59, 0x06, 0x25, 0x9b, 0x79, 0xc1,
	0x4d, 0x21, 0x35, 0xb1, 0x9b, 0x37, 0x88, 0x3c, 0x3d, 0x1e, 0x1c, 0xde, 0x40, 0xd8, 0xaf, 0x80,
	0xa7, 0xbe, 0x69, 0x7c, 0x6f, 0xb6, 0xfc, 0x8f, 0x96, 0xf1, 0x16, 0xa2, 0x21, 0x00, 0xa5, 0xf4,
	0x97, 0x9d, 0x9d, 0x1d, 0xb1, 0xe5, 0x26, 0x94, 0xff, 0xbb, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x93, 0x78, 0xc5, 0x57, 0x65, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserExtClient is the client API for UserExt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserExtClient interface {
	// 登录
	SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error)
	// 获取用户信息
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
}

type userExtClient struct {
	cc *grpc.ClientConn
}

func NewUserExtClient(cc *grpc.ClientConn) UserExtClient {
	return &userExtClient{cc}
}

func (c *userExtClient) SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error) {
	out := new(SignInResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserExtServer is the server API for UserExt service.
type UserExtServer interface {
	// 登录
	SignIn(context.Context, *SignInReq) (*SignInResp, error)
	// 获取用户信息
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
}

// UnimplementedUserExtServer can be embedded to have forward compatible implementations.
type UnimplementedUserExtServer struct {
}

func (*UnimplementedUserExtServer) SignIn(ctx context.Context, req *SignInReq) (*SignInResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedUserExtServer) GetUser(ctx context.Context, req *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func RegisterUserExtServer(s *grpc.Server, srv UserExtServer) {
	s.RegisterService(&_UserExt_serviceDesc, srv)
}

func _UserExt_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).SignIn(ctx, req.(*SignInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserExt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserExt",
	HandlerType: (*UserExtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _UserExt_SignIn_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserExt_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.ext.proto",
}
