// Code generated by protoc-gen-go.
// source: bot.proto
// DO NOT EDIT!

/*
Package bot is a generated protocol buffer package.

It is generated from these files:
	bot.proto

It has these top-level messages:
	BotRequest
	BotResponse
	UserInfo
	UserSession
	Slot
	Intent
	UserIntent
*/
package bot

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

type RequestType int32

const (
	RequestType_ENQUIRE RequestType = 0
	RequestType_CONFIRM RequestType = 1
)

var RequestType_name = map[int32]string{
	0: "ENQUIRE",
	1: "CONFIRM",
}
var RequestType_value = map[string]int32{
	"ENQUIRE": 0,
	"CONFIRM": 1,
}

func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}
func (RequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ReturnCode int32

const (
	ReturnCode_SUCCESS  ReturnCode = 0
	ReturnCode_SYSERR   ReturnCode = 1
	ReturnCode_NULLDATA ReturnCode = 2
)

var ReturnCode_name = map[int32]string{
	0: "SUCCESS",
	1: "SYSERR",
	2: "NULLDATA",
}
var ReturnCode_value = map[string]int32{
	"SUCCESS":  0,
	"SYSERR":   1,
	"NULLDATA": 2,
}

func (x ReturnCode) String() string {
	return proto.EnumName(ReturnCode_name, int32(x))
}
func (ReturnCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BotRequest struct {
	// 基本信息
	BotName     string      `protobuf:"bytes,1,opt,name=bot_name,json=botName" json:"bot_name,omitempty"`
	LogId       string      `protobuf:"bytes,2,opt,name=log_id,json=logId" json:"log_id,omitempty"`
	UserId      string      `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	RequestType RequestType `protobuf:"varint,4,opt,name=request_type,json=requestType,enum=bot.RequestType" json:"request_type,omitempty"`
	Query       string      `protobuf:"bytes,5,opt,name=query" json:"query,omitempty"`
	// 用户个性化信息
	UserInfo *UserInfo `protobuf:"bytes,100,opt,name=user_info,json=userInfo" json:"user_info,omitempty"`
	// 多个session
	Session []*UserSession `protobuf:"bytes,101,rep,name=session" json:"session,omitempty"`
	// 用户的潜在意图
	// 中控把当前bot需要理解的intent以及一些公共intent下发给下游bot
	Intent *UserIntent `protobuf:"bytes,102,opt,name=intent" json:"intent,omitempty"`
}

func (m *BotRequest) Reset()                    { *m = BotRequest{} }
func (m *BotRequest) String() string            { return proto.CompactTextString(m) }
func (*BotRequest) ProtoMessage()               {}
func (*BotRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BotRequest) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *BotRequest) GetSession() []*UserSession {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *BotRequest) GetIntent() *UserIntent {
	if m != nil {
		return m.Intent
	}
	return nil
}

type BotResponse struct {
	Status ReturnCode `protobuf:"varint,1,opt,name=status,enum=bot.ReturnCode" json:"status,omitempty"`
	Score  int32      `protobuf:"varint,2,opt,name=score" json:"score,omitempty"`
	// 多个返回结果，多张内容卡片
	ResItem []*BotResponse_ResultItem `protobuf:"bytes,10,rep,name=res_item,json=resItem" json:"res_item,omitempty"`
	// 多个引导项
	Guide []*BotResponse_Guide `protobuf:"bytes,11,rep,name=guide" json:"guide,omitempty"`
	// 多个session，回写
	Session []*UserSession `protobuf:"bytes,101,rep,name=session" json:"session,omitempty"`
}

func (m *BotResponse) Reset()                    { *m = BotResponse{} }
func (m *BotResponse) String() string            { return proto.CompactTextString(m) }
func (*BotResponse) ProtoMessage()               {}
func (*BotResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BotResponse) GetResItem() []*BotResponse_ResultItem {
	if m != nil {
		return m.ResItem
	}
	return nil
}

func (m *BotResponse) GetGuide() []*BotResponse_Guide {
	if m != nil {
		return m.Guide
	}
	return nil
}

func (m *BotResponse) GetSession() []*UserSession {
	if m != nil {
		return m.Session
	}
	return nil
}

type BotResponse_ResultItem struct {
	Card string `protobuf:"bytes,1,opt,name=card" json:"card,omitempty"`
}

func (m *BotResponse_ResultItem) Reset()                    { *m = BotResponse_ResultItem{} }
func (m *BotResponse_ResultItem) String() string            { return proto.CompactTextString(m) }
func (*BotResponse_ResultItem) ProtoMessage()               {}
func (*BotResponse_ResultItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type BotResponse_Guide struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *BotResponse_Guide) Reset()                    { *m = BotResponse_Guide{} }
func (m *BotResponse_Guide) String() string            { return proto.CompactTextString(m) }
func (*BotResponse_Guide) ProtoMessage()               {}
func (*BotResponse_Guide) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

// 用户个性化信息
type UserInfo struct {
	Gender int32 `protobuf:"varint,1,opt,name=gender" json:"gender,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 用户Session信息
type UserSession struct {
	TableName    string `protobuf:"bytes,1,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	SessionKey   string `protobuf:"bytes,2,opt,name=session_key,json=sessionKey" json:"session_key,omitempty"`
	SessionValue string `protobuf:"bytes,3,opt,name=session_value,json=sessionValue" json:"session_value,omitempty"`
}

func (m *UserSession) Reset()                    { *m = UserSession{} }
func (m *UserSession) String() string            { return proto.CompactTextString(m) }
func (*UserSession) ProtoMessage()               {}
func (*UserSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Slot struct {
	// 槽位名
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// 槽位值
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Slot) Reset()                    { *m = Slot{} }
func (m *Slot) String() string            { return proto.CompactTextString(m) }
func (*Slot) ProtoMessage()               {}
func (*Slot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Intent struct {
	// 意图识别的置信度，目前固定是1分
	Score int32 `protobuf:"varint,1,opt,name=score" json:"score,omitempty"`
	// 该意图下可提取到的多个槽位信息
	Slot []*Slot `protobuf:"bytes,2,rep,name=slot" json:"slot,omitempty"`
	// 意图的附加信息，需求的细分类从这个字段解析出来
	ExtraAttr string `protobuf:"bytes,3,opt,name=extra_attr,json=extraAttr" json:"extra_attr,omitempty"`
}

func (m *Intent) Reset()                    { *m = Intent{} }
func (m *Intent) String() string            { return proto.CompactTextString(m) }
func (*Intent) ProtoMessage()               {}
func (*Intent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Intent) GetSlot() []*Slot {
	if m != nil {
		return m.Slot
	}
	return nil
}

// User意图
type UserIntent struct {
	// 识别出多个不同的意图
	Intent []*Intent `protobuf:"bytes,1,rep,name=intent" json:"intent,omitempty"`
}

func (m *UserIntent) Reset()                    { *m = UserIntent{} }
func (m *UserIntent) String() string            { return proto.CompactTextString(m) }
func (*UserIntent) ProtoMessage()               {}
func (*UserIntent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UserIntent) GetIntent() []*Intent {
	if m != nil {
		return m.Intent
	}
	return nil
}

func init() {
	proto.RegisterType((*BotRequest)(nil), "bot.BotRequest")
	proto.RegisterType((*BotResponse)(nil), "bot.BotResponse")
	proto.RegisterType((*BotResponse_ResultItem)(nil), "bot.BotResponse.ResultItem")
	proto.RegisterType((*BotResponse_Guide)(nil), "bot.BotResponse.Guide")
	proto.RegisterType((*UserInfo)(nil), "bot.UserInfo")
	proto.RegisterType((*UserSession)(nil), "bot.UserSession")
	proto.RegisterType((*Slot)(nil), "bot.Slot")
	proto.RegisterType((*Intent)(nil), "bot.Intent")
	proto.RegisterType((*UserIntent)(nil), "bot.UserIntent")
	proto.RegisterEnum("bot.RequestType", RequestType_name, RequestType_value)
	proto.RegisterEnum("bot.ReturnCode", ReturnCode_name, ReturnCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for BotServer service

type BotServerClient interface {
	Work(ctx context.Context, in *BotRequest, opts ...grpc.CallOption) (*BotResponse, error)
}

type botServerClient struct {
	cc *grpc.ClientConn
}

func NewBotServerClient(cc *grpc.ClientConn) BotServerClient {
	return &botServerClient{cc}
}

func (c *botServerClient) Work(ctx context.Context, in *BotRequest, opts ...grpc.CallOption) (*BotResponse, error) {
	out := new(BotResponse)
	err := grpc.Invoke(ctx, "/bot.BotServer/Work", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BotServer service

type BotServerServer interface {
	Work(context.Context, *BotRequest) (*BotResponse, error)
}

func RegisterBotServerServer(s *grpc.Server, srv BotServerServer) {
	s.RegisterService(&_BotServer_serviceDesc, srv)
}

func _BotServer_Work_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServerServer).Work(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bot.BotServer/Work",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServerServer).Work(ctx, req.(*BotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BotServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bot.BotServer",
	HandlerType: (*BotServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Work",
			Handler:    _BotServer_Work_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("bot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x6d, 0xd2, 0xd8, 0x89, 0xc7, 0xfd, 0xb1, 0x56, 0xdf, 0x57, 0x4c, 0xab, 0x8a, 0xca, 0x5c,
	0x14, 0x15, 0x54, 0x41, 0x2b, 0x21, 0x6e, 0x4b, 0x08, 0x28, 0xa2, 0x04, 0xb1, 0x6e, 0x40, 0x48,
	0x48, 0x96, 0x5b, 0x4f, 0xab, 0xa8, 0xa9, 0xb7, 0xec, 0xae, 0x2b, 0xf2, 0x30, 0x5c, 0xf1, 0xa2,
	0xec, 0x8e, 0xd7, 0xc4, 0xc0, 0x15, 0x77, 0x73, 0xce, 0xfc, 0x9c, 0xf5, 0xcc, 0x49, 0x20, 0x38,
	0x17, 0xfa, 0xf0, 0x56, 0x0a, 0x2d, 0xd8, 0xaa, 0x09, 0x93, 0x1f, 0x5d, 0x80, 0x97, 0x42, 0x73,
	0xfc, 0x5a, 0xa1, 0xd2, 0xec, 0x3e, 0x0c, 0x0c, 0x9b, 0x95, 0xf9, 0x0d, 0xc6, 0x9d, 0xbd, 0xce,
	0xa3, 0x80, 0xf7, 0x0d, 0x9e, 0x18, 0xc8, 0xfe, 0x07, 0x7f, 0x2e, 0xae, 0xb2, 0x59, 0x11, 0x77,
	0x29, 0xe1, 0x19, 0x34, 0x2e, 0xd8, 0x3d, 0xe8, 0x57, 0x0a, 0xa5, 0xe5, 0x57, 0x89, 0xf7, 0x2d,
	0x34, 0x89, 0x63, 0x58, 0x93, 0xf5, 0xd4, 0x4c, 0x2f, 0x6e, 0x31, 0xee, 0x99, 0xec, 0xc6, 0x51,
	0x74, 0x68, 0x1f, 0xe0, 0xe4, 0xce, 0x0c, 0xcf, 0x43, 0xb9, 0x04, 0xec, 0x3f, 0xf0, 0x0c, 0x90,
	0x8b, 0xd8, 0xab, 0x35, 0x08, 0xb0, 0x03, 0x08, 0x6a, 0x8d, 0xf2, 0x52, 0xc4, 0x85, 0xc9, 0x84,
	0x47, 0xeb, 0x34, 0x67, 0x6a, 0xa5, 0x0c, 0xc9, 0x07, 0x95, 0x8b, 0x4c, 0x6d, 0x5f, 0xa1, 0x52,
	0x33, 0x51, 0xc6, 0xb8, 0xb7, 0x6a, 0x2a, 0xa3, 0x5f, 0x95, 0x69, 0xcd, 0xf3, 0xa6, 0x80, 0xed,
	0x83, 0x3f, 0x2b, 0x35, 0x96, 0x3a, 0xbe, 0xa4, 0xa1, 0x9b, 0xad, 0xa1, 0x96, 0xe6, 0x2e, 0x9d,
	0x7c, 0xef, 0x42, 0x48, 0x5b, 0x52, 0xb7, 0xa2, 0x54, 0x68, 0x1b, 0x95, 0xce, 0x75, 0xa5, 0x68,
	0x49, 0x1b, 0xae, 0x91, 0xa3, 0xae, 0x64, 0x39, 0x14, 0x05, 0x72, 0x97, 0xb6, 0xdf, 0xa3, 0x2e,
	0x84, 0x44, 0xda, 0x99, 0xc7, 0x6b, 0xc0, 0x9e, 0xc3, 0x40, 0xa2, 0xca, 0x66, 0x1a, 0x6f, 0x62,
	0xa0, 0x47, 0xee, 0xd0, 0x80, 0x96, 0x84, 0x19, 0xa6, 0xaa, 0xb9, 0x1e, 0x9b, 0x12, 0xde, 0x37,
	0xc5, 0x36, 0x60, 0x4f, 0xc0, 0xbb, 0xaa, 0x66, 0x05, 0xc6, 0x21, 0x35, 0x6d, 0xfd, 0xd5, 0xf4,
	0xc6, 0x66, 0x79, 0x5d, 0xf4, 0x2f, 0x9b, 0xd8, 0xde, 0x03, 0x58, 0x0a, 0x32, 0x06, 0xbd, 0x8b,
	0x5c, 0x16, 0xce, 0x01, 0x14, 0x6f, 0xef, 0x80, 0x47, 0xd3, 0x6d, 0x52, 0xe3, 0x37, 0xdd, 0x24,
	0x6d, 0x9c, 0x24, 0x30, 0x68, 0x4e, 0xc1, 0xb6, 0xc0, 0xbf, 0xc2, 0xb2, 0x40, 0x49, 0x15, 0x1e,
	0x77, 0x28, 0x91, 0x10, 0xb6, 0xa4, 0xd9, 0x2e, 0x80, 0xce, 0xcf, 0xe7, 0xd8, 0xf6, 0x5a, 0x40,
	0x0c, 0xb9, 0xed, 0x01, 0x84, 0xee, 0x6d, 0xd9, 0x35, 0x2e, 0x9c, 0xe5, 0xc0, 0x51, 0x6f, 0x71,
	0xc1, 0x1e, 0xc2, 0x7a, 0x53, 0x70, 0x97, 0xcf, 0x2b, 0x74, 0xee, 0x5b, 0x73, 0xe4, 0x47, 0xcb,
	0x25, 0x4f, 0xa1, 0x97, 0xce, 0x85, 0xb6, 0x6f, 0x6e, 0xc9, 0x50, 0x6c, 0x4f, 0x53, 0x37, 0x3a,
	0x3b, 0x13, 0x48, 0xbe, 0x80, 0x5f, 0xdf, 0x7e, 0x79, 0xba, 0x4e, 0xfb, 0x74, 0xbb, 0xd0, 0x53,
	0x66, 0xa2, 0x69, 0xb2, 0x1b, 0x0d, 0x68, 0xa3, 0x56, 0x82, 0x13, 0x6d, 0xbf, 0xca, 0xec, 0x43,
	0xe6, 0x59, 0xae, 0xb5, 0x74, 0x4f, 0x0a, 0x88, 0x39, 0x31, 0x44, 0xf2, 0x0c, 0x60, 0xe9, 0x2e,
	0xf3, 0x09, 0x8d, 0xfd, 0x3a, 0x34, 0x2d, 0xa4, 0x69, 0xbf, 0x5b, 0xef, 0x60, 0x1f, 0xc2, 0xd6,
	0xaf, 0x85, 0x85, 0xd0, 0x1f, 0x4d, 0x3e, 0x4c, 0xc7, 0x7c, 0x14, 0xad, 0x58, 0x30, 0x7c, 0x3f,
	0x79, 0x3d, 0xe6, 0xef, 0xa2, 0xce, 0xc1, 0xb1, 0x3d, 0x61, 0x63, 0x40, 0x9b, 0x4a, 0xa7, 0xc3,
	0xe1, 0x28, 0x4d, 0x4d, 0x1d, 0x80, 0x9f, 0x7e, 0x4e, 0x47, 0x9c, 0x47, 0x1d, 0xb6, 0x06, 0x83,
	0xc9, 0xf4, 0xf4, 0xf4, 0xd5, 0xc9, 0xd9, 0x49, 0xd4, 0x3d, 0x7a, 0x01, 0x81, 0xf1, 0x4f, 0x8a,
	0xf2, 0x0e, 0x25, 0x7b, 0x0c, 0xbd, 0x4f, 0x42, 0x5e, 0xb3, 0xcd, 0xa5, 0xaf, 0x48, 0x78, 0x3b,
	0xfa, 0xd3, 0x68, 0xc9, 0xca, 0xb9, 0x4f, 0x7f, 0x22, 0xc7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x23, 0xc0, 0x4c, 0x4f, 0x51, 0x04, 0x00, 0x00,
}
