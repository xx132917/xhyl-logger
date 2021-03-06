// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logger.proto

package model

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

//log level
type Level int32

const (
	Level_INFO  Level = 0
	Level_DEBUG Level = 1
	Level_ERROR Level = 2
	Level_WARN  Level = 3
	Level_FATAL Level = 4
	Level_OFF   Level = 5
	Level_TRACE Level = 6
	Level_ALL   Level = 7
)

var Level_name = map[int32]string{
	0: "INFO",
	1: "DEBUG",
	2: "ERROR",
	3: "WARN",
	4: "FATAL",
	5: "OFF",
	6: "TRACE",
	7: "ALL",
}

var Level_value = map[string]int32{
	"INFO":  0,
	"DEBUG": 1,
	"ERROR": 2,
	"WARN":  3,
	"FATAL": 4,
	"OFF":   5,
	"TRACE": 6,
	"ALL":   7,
}

func (x Level) String() string {
	return proto.EnumName(Level_name, int32(x))
}

func (Level) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{0}
}

// The request message containing the user's name.
type LogInfo struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Time                 string   `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Level                string   `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogInfo) Reset()         { *m = LogInfo{} }
func (m *LogInfo) String() string { return proto.CompactTextString(m) }
func (*LogInfo) ProtoMessage()    {}
func (*LogInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{0}
}

func (m *LogInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogInfo.Unmarshal(m, b)
}
func (m *LogInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogInfo.Marshal(b, m, deterministic)
}
func (m *LogInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogInfo.Merge(m, src)
}
func (m *LogInfo) XXX_Size() int {
	return xxx_messageInfo_LogInfo.Size(m)
}
func (m *LogInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LogInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LogInfo proto.InternalMessageInfo

func (m *LogInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *LogInfo) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *LogInfo) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type RequestInfo struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	NodeName             string   `protobuf:"bytes,2,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	Level                Level    `protobuf:"varint,3,opt,name=level,proto3,enum=model.Level" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestInfo) Reset()         { *m = RequestInfo{} }
func (m *RequestInfo) String() string { return proto.CompactTextString(m) }
func (*RequestInfo) ProtoMessage()    {}
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{1}
}

func (m *RequestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestInfo.Unmarshal(m, b)
}
func (m *RequestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestInfo.Marshal(b, m, deterministic)
}
func (m *RequestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestInfo.Merge(m, src)
}
func (m *RequestInfo) XXX_Size() int {
	return xxx_messageInfo_RequestInfo.Size(m)
}
func (m *RequestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RequestInfo proto.InternalMessageInfo

func (m *RequestInfo) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *RequestInfo) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *RequestInfo) GetLevel() Level {
	if m != nil {
		return m.Level
	}
	return Level_INFO
}

type ResposeResult struct {
	Status               int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResposeResult) Reset()         { *m = ResposeResult{} }
func (m *ResposeResult) String() string { return proto.CompactTextString(m) }
func (*ResposeResult) ProtoMessage()    {}
func (*ResposeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d43b7bfc6b6f7b16, []int{2}
}

func (m *ResposeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResposeResult.Unmarshal(m, b)
}
func (m *ResposeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResposeResult.Marshal(b, m, deterministic)
}
func (m *ResposeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResposeResult.Merge(m, src)
}
func (m *ResposeResult) XXX_Size() int {
	return xxx_messageInfo_ResposeResult.Size(m)
}
func (m *ResposeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ResposeResult.DiscardUnknown(m)
}

var xxx_messageInfo_ResposeResult proto.InternalMessageInfo

func (m *ResposeResult) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterEnum("model.Level", Level_name, Level_value)
	proto.RegisterType((*LogInfo)(nil), "model.LogInfo")
	proto.RegisterType((*RequestInfo)(nil), "model.RequestInfo")
	proto.RegisterType((*ResposeResult)(nil), "model.ResposeResult")
}

func init() { proto.RegisterFile("logger.proto", fileDescriptor_d43b7bfc6b6f7b16) }

var fileDescriptor_d43b7bfc6b6f7b16 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xfb, 0x27, 0x4e, 0xda, 0xa3, 0x20, 0xeb, 0x54, 0xa1, 0xa8, 0x13, 0xca, 0x02, 0x62,
	0xe8, 0x50, 0x56, 0x96, 0x14, 0x12, 0xa8, 0x64, 0x12, 0xc9, 0x2a, 0x82, 0x8e, 0x85, 0x1e, 0x11,
	0x52, 0x1a, 0x97, 0xda, 0xe1, 0x03, 0xf1, 0x49, 0x91, 0x4d, 0x1b, 0x65, 0x63, 0xbb, 0xf7, 0xb3,
	0xfc, 0xee, 0xf9, 0x19, 0x46, 0xa5, 0x2a, 0x0a, 0xda, 0x4f, 0x77, 0x7b, 0x65, 0x14, 0xb2, 0xad,
	0xda, 0x50, 0x19, 0x3d, 0x41, 0x20, 0x54, 0xb1, 0xa8, 0x3e, 0x14, 0x86, 0x10, 0xbc, 0xab, 0xca,
	0x50, 0x65, 0xc2, 0xee, 0x45, 0xf7, 0x6a, 0x28, 0x8f, 0x12, 0x11, 0x3c, 0xf3, 0xb9, 0xa5, 0xb0,
	0xe7, 0xb0, 0x9b, 0x71, 0x0c, 0xac, 0xa4, 0x6f, 0x2a, 0xc3, 0xbe, 0x83, 0x7f, 0x22, 0x2a, 0xe0,
	0x44, 0xd2, 0x57, 0x4d, 0xda, 0xfc, 0x63, 0x39, 0x81, 0x41, 0xa5, 0x36, 0x94, 0xad, 0x1b, 0xdb,
	0x46, 0x63, 0xd4, 0xb6, 0x3e, 0x9b, 0x8d, 0xa6, 0x2e, 0xea, 0x54, 0x58, 0x76, 0x5c, 0x74, 0x09,
	0xa7, 0x92, 0xf4, 0x4e, 0x69, 0x92, 0xa4, 0xeb, 0xd2, 0xe0, 0x39, 0xf8, 0xda, 0xac, 0x4d, 0xad,
	0xdd, 0x26, 0x26, 0x0f, 0xea, 0x7a, 0x05, 0xcc, 0x5d, 0xc4, 0x01, 0x78, 0x8b, 0x2c, 0xcd, 0x79,
	0x07, 0x87, 0xc0, 0xee, 0x93, 0xf9, 0xf3, 0x03, 0xef, 0xda, 0x31, 0x91, 0x32, 0x97, 0xbc, 0x67,
	0xcf, 0x5f, 0x62, 0x99, 0xf1, 0xbe, 0x85, 0x69, 0xbc, 0x8c, 0x05, 0xf7, 0x30, 0x80, 0x7e, 0x9e,
	0xa6, 0x9c, 0x59, 0xb6, 0x94, 0xf1, 0x5d, 0xc2, 0x7d, 0xcb, 0x62, 0x21, 0x78, 0x30, 0xbb, 0x05,
	0x5f, 0xb8, 0x4a, 0x71, 0x06, 0x9e, 0x7b, 0x2f, 0x1e, 0xa2, 0xb6, 0x3a, 0x98, 0x8c, 0x1b, 0xd6,
	0x8a, 0x1b, 0x75, 0xe6, 0xc1, 0x4f, 0xcf, 0x7b, 0x7d, 0x5c, 0x89, 0x37, 0xdf, 0x7d, 0xc8, 0xcd,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x6a, 0x75, 0x9c, 0xa0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoggerClient is the client API for Logger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoggerClient interface {
	Info(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*ResposeResult, error)
}

type loggerClient struct {
	cc *grpc.ClientConn
}

func NewLoggerClient(cc *grpc.ClientConn) LoggerClient {
	return &loggerClient{cc}
}

func (c *loggerClient) Info(ctx context.Context, in *RequestInfo, opts ...grpc.CallOption) (*ResposeResult, error) {
	out := new(ResposeResult)
	err := c.cc.Invoke(ctx, "/model.Logger/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServer is the server API for Logger service.
type LoggerServer interface {
	Info(context.Context, *RequestInfo) (*ResposeResult, error)
}

// UnimplementedLoggerServer can be embedded to have forward compatible implementations.
type UnimplementedLoggerServer struct {
}

func (*UnimplementedLoggerServer) Info(ctx context.Context, req *RequestInfo) (*ResposeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

func RegisterLoggerServer(s *grpc.Server, srv LoggerServer) {
	s.RegisterService(&_Logger_serviceDesc, srv)
}

func _Logger_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Logger/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).Info(ctx, req.(*RequestInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Logger_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Logger",
	HandlerType: (*LoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Logger_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logger.proto",
}
