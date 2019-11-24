// Code generated by protoc-gen-go. DO NOT EDIT.
// source: groupbirthday.proto

package groupbirthday

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

type GetGroupsRequest struct {
	MemberId             int32    `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupsRequest) Reset()         { *m = GetGroupsRequest{} }
func (m *GetGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupsRequest) ProtoMessage()    {}
func (*GetGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_203a0210451df990, []int{0}
}

func (m *GetGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupsRequest.Unmarshal(m, b)
}
func (m *GetGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupsRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupsRequest.Merge(m, src)
}
func (m *GetGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupsRequest.Size(m)
}
func (m *GetGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupsRequest proto.InternalMessageInfo

func (m *GetGroupsRequest) GetMemberId() int32 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

type GetGroupsReply struct {
	Groups               []string `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupsReply) Reset()         { *m = GetGroupsReply{} }
func (m *GetGroupsReply) String() string { return proto.CompactTextString(m) }
func (*GetGroupsReply) ProtoMessage()    {}
func (*GetGroupsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_203a0210451df990, []int{1}
}

func (m *GetGroupsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupsReply.Unmarshal(m, b)
}
func (m *GetGroupsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupsReply.Marshal(b, m, deterministic)
}
func (m *GetGroupsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupsReply.Merge(m, src)
}
func (m *GetGroupsReply) XXX_Size() int {
	return xxx_messageInfo_GetGroupsReply.Size(m)
}
func (m *GetGroupsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupsReply proto.InternalMessageInfo

func (m *GetGroupsReply) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

type GetMemberIdRequest struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	TelegramUsername     string   `protobuf:"bytes,3,opt,name=telegramUsername,proto3" json:"telegramUsername,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemberIdRequest) Reset()         { *m = GetMemberIdRequest{} }
func (m *GetMemberIdRequest) String() string { return proto.CompactTextString(m) }
func (*GetMemberIdRequest) ProtoMessage()    {}
func (*GetMemberIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_203a0210451df990, []int{2}
}

func (m *GetMemberIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemberIdRequest.Unmarshal(m, b)
}
func (m *GetMemberIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemberIdRequest.Marshal(b, m, deterministic)
}
func (m *GetMemberIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberIdRequest.Merge(m, src)
}
func (m *GetMemberIdRequest) XXX_Size() int {
	return xxx_messageInfo_GetMemberIdRequest.Size(m)
}
func (m *GetMemberIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberIdRequest proto.InternalMessageInfo

func (m *GetMemberIdRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *GetMemberIdRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *GetMemberIdRequest) GetTelegramUsername() string {
	if m != nil {
		return m.TelegramUsername
	}
	return ""
}

type GetMemberIdReply struct {
	MemberId             int32    `protobuf:"varint,1,opt,name=memberId,proto3" json:"memberId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemberIdReply) Reset()         { *m = GetMemberIdReply{} }
func (m *GetMemberIdReply) String() string { return proto.CompactTextString(m) }
func (*GetMemberIdReply) ProtoMessage()    {}
func (*GetMemberIdReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_203a0210451df990, []int{3}
}

func (m *GetMemberIdReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemberIdReply.Unmarshal(m, b)
}
func (m *GetMemberIdReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemberIdReply.Marshal(b, m, deterministic)
}
func (m *GetMemberIdReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberIdReply.Merge(m, src)
}
func (m *GetMemberIdReply) XXX_Size() int {
	return xxx_messageInfo_GetMemberIdReply.Size(m)
}
func (m *GetMemberIdReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberIdReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberIdReply proto.InternalMessageInfo

func (m *GetMemberIdReply) GetMemberId() int32 {
	if m != nil {
		return m.MemberId
	}
	return 0
}

type GetMemberBirthdaysRequest struct {
	GroupName            string   `protobuf:"bytes,1,opt,name=groupName,proto3" json:"groupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMemberBirthdaysRequest) Reset()         { *m = GetMemberBirthdaysRequest{} }
func (m *GetMemberBirthdaysRequest) String() string { return proto.CompactTextString(m) }
func (*GetMemberBirthdaysRequest) ProtoMessage()    {}
func (*GetMemberBirthdaysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_203a0210451df990, []int{4}
}

func (m *GetMemberBirthdaysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemberBirthdaysRequest.Unmarshal(m, b)
}
func (m *GetMemberBirthdaysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemberBirthdaysRequest.Marshal(b, m, deterministic)
}
func (m *GetMemberBirthdaysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberBirthdaysRequest.Merge(m, src)
}
func (m *GetMemberBirthdaysRequest) XXX_Size() int {
	return xxx_messageInfo_GetMemberBirthdaysRequest.Size(m)
}
func (m *GetMemberBirthdaysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberBirthdaysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberBirthdaysRequest proto.InternalMessageInfo

func (m *GetMemberBirthdaysRequest) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type GetMemberBirthdaysReply struct {
	MemberBirthdays      []*MemberBirthday `protobuf:"bytes,1,rep,name=memberBirthdays,proto3" json:"memberBirthdays,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetMemberBirthdaysReply) Reset()         { *m = GetMemberBirthdaysReply{} }
func (m *GetMemberBirthdaysReply) String() string { return proto.CompactTextString(m) }
func (*GetMemberBirthdaysReply) ProtoMessage()    {}
func (*GetMemberBirthdaysReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_203a0210451df990, []int{5}
}

func (m *GetMemberBirthdaysReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMemberBirthdaysReply.Unmarshal(m, b)
}
func (m *GetMemberBirthdaysReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMemberBirthdaysReply.Marshal(b, m, deterministic)
}
func (m *GetMemberBirthdaysReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMemberBirthdaysReply.Merge(m, src)
}
func (m *GetMemberBirthdaysReply) XXX_Size() int {
	return xxx_messageInfo_GetMemberBirthdaysReply.Size(m)
}
func (m *GetMemberBirthdaysReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMemberBirthdaysReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMemberBirthdaysReply proto.InternalMessageInfo

func (m *GetMemberBirthdaysReply) GetMemberBirthdays() []*MemberBirthday {
	if m != nil {
		return m.MemberBirthdays
	}
	return nil
}

type MemberBirthday struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Day                  int32    `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	Month                int32    `protobuf:"varint,4,opt,name=month,proto3" json:"month,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberBirthday) Reset()         { *m = MemberBirthday{} }
func (m *MemberBirthday) String() string { return proto.CompactTextString(m) }
func (*MemberBirthday) ProtoMessage()    {}
func (*MemberBirthday) Descriptor() ([]byte, []int) {
	return fileDescriptor_203a0210451df990, []int{6}
}

func (m *MemberBirthday) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberBirthday.Unmarshal(m, b)
}
func (m *MemberBirthday) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberBirthday.Marshal(b, m, deterministic)
}
func (m *MemberBirthday) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberBirthday.Merge(m, src)
}
func (m *MemberBirthday) XXX_Size() int {
	return xxx_messageInfo_MemberBirthday.Size(m)
}
func (m *MemberBirthday) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberBirthday.DiscardUnknown(m)
}

var xxx_messageInfo_MemberBirthday proto.InternalMessageInfo

func (m *MemberBirthday) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *MemberBirthday) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *MemberBirthday) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *MemberBirthday) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func init() {
	proto.RegisterType((*GetGroupsRequest)(nil), "GetGroupsRequest")
	proto.RegisterType((*GetGroupsReply)(nil), "GetGroupsReply")
	proto.RegisterType((*GetMemberIdRequest)(nil), "GetMemberIdRequest")
	proto.RegisterType((*GetMemberIdReply)(nil), "GetMemberIdReply")
	proto.RegisterType((*GetMemberBirthdaysRequest)(nil), "GetMemberBirthdaysRequest")
	proto.RegisterType((*GetMemberBirthdaysReply)(nil), "GetMemberBirthdaysReply")
	proto.RegisterType((*MemberBirthday)(nil), "MemberBirthday")
}

func init() { proto.RegisterFile("groupbirthday.proto", fileDescriptor_203a0210451df990) }

var fileDescriptor_203a0210451df990 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x6d, 0x28, 0xad, 0xc8, 0x55, 0x34, 0xa9, 0x8b, 0x20, 0x44, 0x0c, 0x95, 0xa7, 0x88, 0xc1,
	0x12, 0x65, 0x40, 0x5d, 0x59, 0x22, 0x24, 0x60, 0xb0, 0xe0, 0x07, 0x24, 0x8a, 0x69, 0x2b, 0xc5,
	0x4d, 0xb0, 0xdd, 0x21, 0xfc, 0x43, 0xfe, 0x15, 0xb2, 0x09, 0x6e, 0x92, 0x52, 0x16, 0x36, 0xdf,
	0xbb, 0xe7, 0xfb, 0x78, 0xf7, 0x60, 0xba, 0x14, 0xc5, 0xb6, 0x4c, 0xd7, 0x42, 0xad, 0xb2, 0xa4,
	0x22, 0xa5, 0x28, 0x54, 0x81, 0x09, 0xf8, 0x31, 0x53, 0xb1, 0xce, 0x48, 0xca, 0xde, 0xb7, 0x4c,
	0x2a, 0x14, 0xc2, 0x09, 0x67, 0x3c, 0x65, 0xe2, 0x21, 0x0b, 0x9c, 0x99, 0x13, 0x0d, 0xa8, 0x8d,
	0x71, 0x04, 0xe3, 0x06, 0xbf, 0xcc, 0x2b, 0x74, 0x0e, 0x43, 0x53, 0x58, 0x06, 0xce, 0xac, 0x1f,
	0xb9, 0xb4, 0x8e, 0xf0, 0x07, 0xa0, 0x98, 0xa9, 0xa7, 0xfa, 0xe3, 0x4f, 0xed, 0x2b, 0x70, 0xdf,
	0xd6, 0x42, 0xaa, 0xe7, 0x84, 0x33, 0x53, 0xdc, 0xa5, 0x3b, 0x40, 0x77, 0xce, 0x93, 0x3a, 0x79,
	0x64, 0x92, 0x36, 0x46, 0xd7, 0xe0, 0x2b, 0x96, 0xb3, 0xa5, 0x48, 0xf8, 0xab, 0x64, 0x62, 0xa3,
	0x39, 0x7d, 0xc3, 0xd9, 0xc3, 0xeb, 0xad, 0x76, 0xbd, 0xf5, 0x9c, 0x7f, 0x6d, 0xb5, 0x80, 0x4b,
	0xcb, 0xbf, 0xaf, 0x05, 0x92, 0x8d, 0x91, 0xcd, 0x4a, 0xcd, 0x91, 0x2d, 0x80, 0x5f, 0xe0, 0xe2,
	0xb7, 0xaf, 0xba, 0xe3, 0x02, 0x3c, 0xde, 0xc6, 0x8d, 0x44, 0xa3, 0xb9, 0x47, 0xda, 0x7c, 0xda,
	0xe5, 0x61, 0x01, 0xe3, 0x36, 0xe5, 0x1f, 0xc2, 0xf9, 0xd0, 0xcf, 0x92, 0xca, 0x68, 0x35, 0xa0,
	0xfa, 0x89, 0xce, 0x60, 0xc0, 0x8b, 0x8d, 0x5a, 0x05, 0xc7, 0x06, 0xfb, 0x0e, 0xe6, 0x9f, 0x0e,
	0x9c, 0x9a, 0xc3, 0xda, 0x9e, 0x37, 0xe0, 0xda, 0x63, 0xa3, 0x09, 0xe9, 0x1a, 0x25, 0xf4, 0x48,
	0xdb, 0x0b, 0xb8, 0x87, 0xee, 0x60, 0xd4, 0x50, 0x1e, 0x4d, 0xc9, 0xbe, 0x07, 0xc2, 0x09, 0xe9,
	0x1e, 0x07, 0xf7, 0xd0, 0x63, 0xc3, 0x2e, 0x56, 0x07, 0x14, 0x92, 0x83, 0x77, 0x09, 0x03, 0x72,
	0x40, 0x78, 0xdc, 0x4b, 0x87, 0xc6, 0xdd, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x0c,
	0xed, 0x28, 0xf4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GroupBirthdayClient is the client API for GroupBirthday service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupBirthdayClient interface {
	GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsReply, error)
	GetMemberId(ctx context.Context, in *GetMemberIdRequest, opts ...grpc.CallOption) (*GetMemberIdReply, error)
	GetMemberBirthdays(ctx context.Context, in *GetMemberBirthdaysRequest, opts ...grpc.CallOption) (*GetMemberBirthdaysReply, error)
}

type groupBirthdayClient struct {
	cc *grpc.ClientConn
}

func NewGroupBirthdayClient(cc *grpc.ClientConn) GroupBirthdayClient {
	return &groupBirthdayClient{cc}
}

func (c *groupBirthdayClient) GetGroups(ctx context.Context, in *GetGroupsRequest, opts ...grpc.CallOption) (*GetGroupsReply, error) {
	out := new(GetGroupsReply)
	err := c.cc.Invoke(ctx, "/GroupBirthday/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupBirthdayClient) GetMemberId(ctx context.Context, in *GetMemberIdRequest, opts ...grpc.CallOption) (*GetMemberIdReply, error) {
	out := new(GetMemberIdReply)
	err := c.cc.Invoke(ctx, "/GroupBirthday/GetMemberId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupBirthdayClient) GetMemberBirthdays(ctx context.Context, in *GetMemberBirthdaysRequest, opts ...grpc.CallOption) (*GetMemberBirthdaysReply, error) {
	out := new(GetMemberBirthdaysReply)
	err := c.cc.Invoke(ctx, "/GroupBirthday/GetMemberBirthdays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupBirthdayServer is the server API for GroupBirthday service.
type GroupBirthdayServer interface {
	GetGroups(context.Context, *GetGroupsRequest) (*GetGroupsReply, error)
	GetMemberId(context.Context, *GetMemberIdRequest) (*GetMemberIdReply, error)
	GetMemberBirthdays(context.Context, *GetMemberBirthdaysRequest) (*GetMemberBirthdaysReply, error)
}

// UnimplementedGroupBirthdayServer can be embedded to have forward compatible implementations.
type UnimplementedGroupBirthdayServer struct {
}

func (*UnimplementedGroupBirthdayServer) GetGroups(ctx context.Context, req *GetGroupsRequest) (*GetGroupsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (*UnimplementedGroupBirthdayServer) GetMemberId(ctx context.Context, req *GetMemberIdRequest) (*GetMemberIdReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberId not implemented")
}
func (*UnimplementedGroupBirthdayServer) GetMemberBirthdays(ctx context.Context, req *GetMemberBirthdaysRequest) (*GetMemberBirthdaysReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemberBirthdays not implemented")
}

func RegisterGroupBirthdayServer(s *grpc.Server, srv GroupBirthdayServer) {
	s.RegisterService(&_GroupBirthday_serviceDesc, srv)
}

func _GroupBirthday_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupBirthdayServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupBirthday/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupBirthdayServer).GetGroups(ctx, req.(*GetGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupBirthday_GetMemberId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupBirthdayServer).GetMemberId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupBirthday/GetMemberId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupBirthdayServer).GetMemberId(ctx, req.(*GetMemberIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupBirthday_GetMemberBirthdays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberBirthdaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupBirthdayServer).GetMemberBirthdays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GroupBirthday/GetMemberBirthdays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupBirthdayServer).GetMemberBirthdays(ctx, req.(*GetMemberBirthdaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupBirthday_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GroupBirthday",
	HandlerType: (*GroupBirthdayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroups",
			Handler:    _GroupBirthday_GetGroups_Handler,
		},
		{
			MethodName: "GetMemberId",
			Handler:    _GroupBirthday_GetMemberId_Handler,
		},
		{
			MethodName: "GetMemberBirthdays",
			Handler:    _GroupBirthday_GetMemberBirthdays_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "groupbirthday.proto",
}
