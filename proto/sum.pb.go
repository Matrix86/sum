// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sum.proto

/*
Package sum is a generated protocol buffer package.

It is generated from these files:
	proto/sum.proto

It has these top-level messages:
	NamedValue
	Record
	Oracle
	EvalRequest
	Query
	Response
	ServerInfo
	Empty
*/
package sum

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

type NamedValue struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *NamedValue) Reset()                    { *m = NamedValue{} }
func (m *NamedValue) String() string            { return proto.CompactTextString(m) }
func (*NamedValue) ProtoMessage()               {}
func (*NamedValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NamedValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NamedValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Record struct {
	Id   string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Data []float32     `protobuf:"fixed32,2,rep,packed,name=data" json:"data,omitempty"`
	Meta []*NamedValue `protobuf:"bytes,3,rep,name=meta" json:"meta,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Record) GetData() []float32 {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Record) GetMeta() []*NamedValue {
	if m != nil {
		return m.Meta
	}
	return nil
}

type Oracle struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Code string `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
}

func (m *Oracle) Reset()                    { *m = Oracle{} }
func (m *Oracle) String() string            { return proto.CompactTextString(m) }
func (*Oracle) ProtoMessage()               {}
func (*Oracle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Oracle) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Oracle) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Oracle) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type EvalRequest struct {
	OracleId string        `protobuf:"bytes,1,opt,name=oracle_id,json=oracleId" json:"oracle_id,omitempty"`
	Args     []*NamedValue `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *EvalRequest) Reset()                    { *m = EvalRequest{} }
func (m *EvalRequest) String() string            { return proto.CompactTextString(m) }
func (*EvalRequest) ProtoMessage()               {}
func (*EvalRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EvalRequest) GetOracleId() string {
	if m != nil {
		return m.OracleId
	}
	return ""
}

func (m *EvalRequest) GetArgs() []*NamedValue {
	if m != nil {
		return m.Args
	}
	return nil
}

type Query struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Response struct {
	Success bool    `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg     string  `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Record  *Record `protobuf:"bytes,3,opt,name=record" json:"record,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type ServerInfo struct {
	Version string   `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Uptime  uint64   `protobuf:"varint,2,opt,name=uptime" json:"uptime,omitempty"`
	Pid     uint64   `protobuf:"varint,3,opt,name=pid" json:"pid,omitempty"`
	Uid     uint64   `protobuf:"varint,4,opt,name=uid" json:"uid,omitempty"`
	Argv    []string `protobuf:"bytes,5,rep,name=argv" json:"argv,omitempty"`
	Records uint64   `protobuf:"varint,6,opt,name=records" json:"records,omitempty"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ServerInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServerInfo) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *ServerInfo) GetPid() uint64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ServerInfo) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ServerInfo) GetArgv() []string {
	if m != nil {
		return m.Argv
	}
	return nil
}

func (m *ServerInfo) GetRecords() uint64 {
	if m != nil {
		return m.Records
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*NamedValue)(nil), "sum.NamedValue")
	proto.RegisterType((*Record)(nil), "sum.Record")
	proto.RegisterType((*Oracle)(nil), "sum.Oracle")
	proto.RegisterType((*EvalRequest)(nil), "sum.EvalRequest")
	proto.RegisterType((*Query)(nil), "sum.Query")
	proto.RegisterType((*Response)(nil), "sum.Response")
	proto.RegisterType((*ServerInfo)(nil), "sum.ServerInfo")
	proto.RegisterType((*Empty)(nil), "sum.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SumService service

type SumServiceClient interface {
	// vectors CRUD
	Create(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Response, error)
	Read(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error)
	// core methods
	// TODO: rpc Eval(EvalRequest) (???) {}
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerInfo, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Create(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/sum.SumService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) Update(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/sum.SumService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) Read(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/sum.SumService/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) Delete(ctx context.Context, in *Query, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/sum.SumService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerInfo, error) {
	out := new(ServerInfo)
	err := grpc.Invoke(ctx, "/sum.SumService/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SumService service

type SumServiceServer interface {
	// vectors CRUD
	Create(context.Context, *Record) (*Response, error)
	Update(context.Context, *Record) (*Response, error)
	Read(context.Context, *Query) (*Response, error)
	Delete(context.Context, *Query) (*Response, error)
	// core methods
	// TODO: rpc Eval(EvalRequest) (???) {}
	Info(context.Context, *Empty) (*ServerInfo, error)
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Create(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Update(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Read(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Delete(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Info(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sum.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SumService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SumService_Update_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _SumService_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SumService_Delete_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _SumService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/sum.proto",
}

func init() { proto.RegisterFile("proto/sum.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6a, 0x14, 0x41,
	0x10, 0xc6, 0xb3, 0x33, 0xb3, 0x93, 0x4d, 0x2d, 0x1a, 0x69, 0x44, 0x9b, 0x78, 0x59, 0x7a, 0x51,
	0xf6, 0x14, 0x61, 0x05, 0xcf, 0x82, 0xe6, 0x90, 0x8b, 0x21, 0x2d, 0x7a, 0xf1, 0x20, 0xed, 0x74,
	0xb9, 0x0c, 0xec, 0x4c, 0x8f, 0xfd, 0x67, 0x20, 0x6f, 0xe1, 0x9b, 0xf9, 0x4a, 0x52, 0xd5, 0xb3,
	0x26, 0x68, 0x20, 0xb7, 0xaf, 0xaa, 0xbe, 0xa9, 0xfa, 0xcd, 0x37, 0x0c, 0x9c, 0x0e, 0xde, 0x45,
	0xf7, 0x3a, 0xa4, 0xee, 0x9c, 0x95, 0x28, 0x43, 0xea, 0xd4, 0x5b, 0x80, 0x8f, 0xa6, 0x43, 0xfb,
	0xc5, 0xec, 0x13, 0x0a, 0x01, 0x55, 0x6f, 0x3a, 0x94, 0xb3, 0xd5, 0x6c, 0x73, 0xa2, 0x59, 0x8b,
	0xa7, 0x30, 0x1f, 0x69, 0x28, 0x0b, 0x6e, 0xe6, 0x42, 0x5d, 0x43, 0xad, 0xb1, 0x71, 0xde, 0x8a,
	0xc7, 0x50, 0xb4, 0x76, 0x7a, 0xa2, 0x68, 0x2d, 0xed, 0xb0, 0x26, 0x1a, 0x59, 0xac, 0xca, 0x4d,
	0xa1, 0x59, 0x8b, 0x35, 0x54, 0x1d, 0x46, 0x23, 0xcb, 0x55, 0xb9, 0x59, 0x6e, 0x4f, 0xcf, 0x09,
	0xe2, 0xf6, 0xac, 0xe6, 0xa1, 0x7a, 0x07, 0xf5, 0x95, 0x37, 0xcd, 0x1e, 0xef, 0x5b, 0xc9, 0x58,
	0xc5, 0x1d, 0x2c, 0x01, 0x55, 0xe3, 0x2c, 0xca, 0x32, 0xf7, 0x48, 0xab, 0x2b, 0x58, 0x5e, 0x8c,
	0x66, 0xaf, 0xf1, 0x67, 0xc2, 0x10, 0xc5, 0x0b, 0x38, 0x71, 0xbc, 0xf0, 0xdb, 0xdf, 0x6d, 0x8b,
	0xdc, 0xb8, 0xb4, 0x84, 0x64, 0xfc, 0x2e, 0x30, 0xe6, 0x7d, 0x48, 0x34, 0x54, 0xcf, 0x61, 0x7e,
	0x9d, 0xd0, 0xdf, 0xfc, 0x4b, 0xa4, 0xbe, 0xc2, 0x42, 0x63, 0x18, 0x5c, 0x1f, 0x50, 0x48, 0x38,
	0x0e, 0xa9, 0x69, 0x30, 0x04, 0x36, 0x2c, 0xf4, 0xa1, 0x14, 0x4f, 0xa0, 0xec, 0xc2, 0x6e, 0xc2,
	0x26, 0x29, 0xd6, 0x50, 0x7b, 0x8e, 0x8d, 0xb9, 0x97, 0xdb, 0x25, 0xdf, 0xcd, 0x49, 0xea, 0x69,
	0xa4, 0x7e, 0xcd, 0x00, 0x3e, 0xa1, 0x1f, 0xd1, 0x5f, 0xf6, 0x3f, 0x1c, 0xed, 0x1f, 0xd1, 0x87,
	0xd6, 0xf5, 0x13, 0xc0, 0xa1, 0x14, 0xcf, 0xa0, 0x4e, 0x43, 0x6c, 0xa7, 0x64, 0x2a, 0x3d, 0x55,
	0x74, 0x77, 0x68, 0xf3, 0x89, 0x4a, 0x93, 0xa4, 0x4e, 0x6a, 0xad, 0xac, 0x72, 0x27, 0xe5, 0x4c,
	0x8d, 0xdf, 0x8d, 0x72, 0xbe, 0x2a, 0x29, 0x3f, 0xd2, 0x74, 0x29, 0x23, 0x04, 0x59, 0xb3, 0xf3,
	0x50, 0xaa, 0x63, 0x98, 0x5f, 0x74, 0x43, 0xbc, 0xd9, 0xfe, 0x26, 0xb6, 0xd4, 0x11, 0x5e, 0xdb,
	0xa0, 0x78, 0x05, 0xf5, 0x7b, 0x8f, 0x26, 0xa2, 0xb8, 0xfb, 0x26, 0x67, 0x8f, 0xa6, 0x22, 0x27,
	0xa4, 0x8e, 0xc8, 0xf7, 0x79, 0xb0, 0x0f, 0xfb, 0xd6, 0x50, 0x69, 0x34, 0x56, 0x00, 0x0f, 0x38,
	0xfb, 0xff, 0x4d, 0x2f, 0xa1, 0xfe, 0x80, 0x7b, 0x8c, 0xf8, 0x90, 0xad, 0xe2, 0xfc, 0xb2, 0x89,
	0xf1, 0xcf, 0xf2, 0x77, 0xbe, 0x0d, 0x57, 0x1d, 0x7d, 0xaf, 0xf9, 0x6f, 0x78, 0xf3, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xf8, 0xea, 0x27, 0xc0, 0x20, 0x03, 0x00, 0x00,
}
