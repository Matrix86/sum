// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/sum.proto

/*
Package sum is a generated protocol buffer package.

It is generated from these files:
	proto/sum.proto

It has these top-level messages:
	NamedValue
	Record
	RecordResponse
	Oracle
	OracleResponse
	Call
	CallResponse
	ById
	ByName
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

type RecordResponse struct {
	Success bool    `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg     string  `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Record  *Record `protobuf:"bytes,3,opt,name=record" json:"record,omitempty"`
}

func (m *RecordResponse) Reset()                    { *m = RecordResponse{} }
func (m *RecordResponse) String() string            { return proto.CompactTextString(m) }
func (*RecordResponse) ProtoMessage()               {}
func (*RecordResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RecordResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RecordResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RecordResponse) GetRecord() *Record {
	if m != nil {
		return m.Record
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
func (*Oracle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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

type OracleResponse struct {
	Success bool      `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg     string    `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Oracles []*Oracle `protobuf:"bytes,3,rep,name=oracles" json:"oracles,omitempty"`
}

func (m *OracleResponse) Reset()                    { *m = OracleResponse{} }
func (m *OracleResponse) String() string            { return proto.CompactTextString(m) }
func (*OracleResponse) ProtoMessage()               {}
func (*OracleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OracleResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *OracleResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *OracleResponse) GetOracles() []*Oracle {
	if m != nil {
		return m.Oracles
	}
	return nil
}

type Call struct {
	OracleId string   `protobuf:"bytes,1,opt,name=oracle_id,json=oracleId" json:"oracle_id,omitempty"`
	Args     []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *Call) Reset()                    { *m = Call{} }
func (m *Call) String() string            { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()               {}
func (*Call) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Call) GetOracleId() string {
	if m != nil {
		return m.OracleId
	}
	return ""
}

func (m *Call) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type CallResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Json    string `protobuf:"bytes,3,opt,name=json" json:"json,omitempty"`
}

func (m *CallResponse) Reset()                    { *m = CallResponse{} }
func (m *CallResponse) String() string            { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()               {}
func (*CallResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CallResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CallResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CallResponse) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

type ById struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ById) Reset()                    { *m = ById{} }
func (m *ById) String() string            { return proto.CompactTextString(m) }
func (*ById) ProtoMessage()               {}
func (*ById) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ById) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ByName struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ByName) Reset()                    { *m = ByName{} }
func (m *ByName) String() string            { return proto.CompactTextString(m) }
func (*ByName) ProtoMessage()               {}
func (*ByName) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ByName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ServerInfo struct {
	Version string   `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Uptime  uint64   `protobuf:"varint,2,opt,name=uptime" json:"uptime,omitempty"`
	Pid     uint64   `protobuf:"varint,3,opt,name=pid" json:"pid,omitempty"`
	Uid     uint64   `protobuf:"varint,4,opt,name=uid" json:"uid,omitempty"`
	Argv    []string `protobuf:"bytes,5,rep,name=argv" json:"argv,omitempty"`
	Records uint64   `protobuf:"varint,6,opt,name=records" json:"records,omitempty"`
	Oracles uint64   `protobuf:"varint,7,opt,name=oracles" json:"oracles,omitempty"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

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

func (m *ServerInfo) GetOracles() uint64 {
	if m != nil {
		return m.Oracles
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*NamedValue)(nil), "sum.NamedValue")
	proto.RegisterType((*Record)(nil), "sum.Record")
	proto.RegisterType((*RecordResponse)(nil), "sum.RecordResponse")
	proto.RegisterType((*Oracle)(nil), "sum.Oracle")
	proto.RegisterType((*OracleResponse)(nil), "sum.OracleResponse")
	proto.RegisterType((*Call)(nil), "sum.Call")
	proto.RegisterType((*CallResponse)(nil), "sum.CallResponse")
	proto.RegisterType((*ById)(nil), "sum.ById")
	proto.RegisterType((*ByName)(nil), "sum.ByName")
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
	CreateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*RecordResponse, error)
	UpdateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*RecordResponse, error)
	ReadRecord(ctx context.Context, in *ById, opts ...grpc.CallOption) (*RecordResponse, error)
	DeleteRecord(ctx context.Context, in *ById, opts ...grpc.CallOption) (*RecordResponse, error)
	// oracles CRUD
	CreateOracle(ctx context.Context, in *Oracle, opts ...grpc.CallOption) (*OracleResponse, error)
	UpdateOracle(ctx context.Context, in *Oracle, opts ...grpc.CallOption) (*OracleResponse, error)
	ReadOracle(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OracleResponse, error)
	FindOracle(ctx context.Context, in *ByName, opts ...grpc.CallOption) (*OracleResponse, error)
	DeleteOracle(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OracleResponse, error)
	// execute a call to a oracle given its id
	Run(ctx context.Context, in *Call, opts ...grpc.CallOption) (*CallResponse, error)
	// get info about the service
	Info(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServerInfo, error)
}

type sumServiceClient struct {
	cc *grpc.ClientConn
}

func NewSumServiceClient(cc *grpc.ClientConn) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) CreateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/CreateRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) UpdateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/UpdateRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) ReadRecord(ctx context.Context, in *ById, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/ReadRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) DeleteRecord(ctx context.Context, in *ById, opts ...grpc.CallOption) (*RecordResponse, error) {
	out := new(RecordResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/DeleteRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) CreateOracle(ctx context.Context, in *Oracle, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/CreateOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) UpdateOracle(ctx context.Context, in *Oracle, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/UpdateOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) ReadOracle(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/ReadOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) FindOracle(ctx context.Context, in *ByName, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/FindOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) DeleteOracle(ctx context.Context, in *ById, opts ...grpc.CallOption) (*OracleResponse, error) {
	out := new(OracleResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/DeleteOracle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) Run(ctx context.Context, in *Call, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := grpc.Invoke(ctx, "/sum.SumService/Run", in, out, c.cc, opts...)
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
	CreateRecord(context.Context, *Record) (*RecordResponse, error)
	UpdateRecord(context.Context, *Record) (*RecordResponse, error)
	ReadRecord(context.Context, *ById) (*RecordResponse, error)
	DeleteRecord(context.Context, *ById) (*RecordResponse, error)
	// oracles CRUD
	CreateOracle(context.Context, *Oracle) (*OracleResponse, error)
	UpdateOracle(context.Context, *Oracle) (*OracleResponse, error)
	ReadOracle(context.Context, *ById) (*OracleResponse, error)
	FindOracle(context.Context, *ByName) (*OracleResponse, error)
	DeleteOracle(context.Context, *ById) (*OracleResponse, error)
	// execute a call to a oracle given its id
	Run(context.Context, *Call) (*CallResponse, error)
	// get info about the service
	Info(context.Context, *Empty) (*ServerInfo, error)
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).CreateRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/UpdateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).UpdateRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_ReadRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).ReadRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/ReadRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).ReadRecord(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).DeleteRecord(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_CreateOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Oracle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).CreateOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/CreateOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).CreateOracle(ctx, req.(*Oracle))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_UpdateOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Oracle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).UpdateOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/UpdateOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).UpdateOracle(ctx, req.(*Oracle))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_ReadOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).ReadOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/ReadOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).ReadOracle(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_FindOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).FindOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/FindOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).FindOracle(ctx, req.(*ByName))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_DeleteOracle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).DeleteOracle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/DeleteOracle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).DeleteOracle(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Call)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sum.SumService/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Run(ctx, req.(*Call))
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
			MethodName: "CreateRecord",
			Handler:    _SumService_CreateRecord_Handler,
		},
		{
			MethodName: "UpdateRecord",
			Handler:    _SumService_UpdateRecord_Handler,
		},
		{
			MethodName: "ReadRecord",
			Handler:    _SumService_ReadRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _SumService_DeleteRecord_Handler,
		},
		{
			MethodName: "CreateOracle",
			Handler:    _SumService_CreateOracle_Handler,
		},
		{
			MethodName: "UpdateOracle",
			Handler:    _SumService_UpdateOracle_Handler,
		},
		{
			MethodName: "ReadOracle",
			Handler:    _SumService_ReadOracle_Handler,
		},
		{
			MethodName: "FindOracle",
			Handler:    _SumService_FindOracle_Handler,
		},
		{
			MethodName: "DeleteOracle",
			Handler:    _SumService_DeleteOracle_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _SumService_Run_Handler,
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
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x6d, 0xec, 0x8d, 0xd3, 0x4c, 0xa2, 0x14, 0x16, 0x54, 0x59, 0x81, 0x43, 0xb4, 0x55, 0xa5,
	0x9c, 0x42, 0x15, 0x24, 0xb8, 0xa2, 0x16, 0x90, 0x72, 0x09, 0x62, 0x2b, 0xb8, 0xa2, 0xc5, 0xbb,
	0x54, 0x46, 0xfe, 0x92, 0xd7, 0xb6, 0x94, 0x7f, 0xc3, 0x99, 0x5f, 0x89, 0x76, 0x76, 0xed, 0x3a,
	0x84, 0x8a, 0xb6, 0xb7, 0xf9, 0x7a, 0x33, 0x6f, 0x66, 0x9f, 0x0d, 0x27, 0x45, 0x99, 0x57, 0xf9,
	0x2b, 0x5d, 0xa7, 0x2b, 0xb4, 0xa8, 0xaf, 0xeb, 0x94, 0xbd, 0x01, 0xd8, 0x8a, 0x54, 0xc9, 0xaf,
	0x22, 0xa9, 0x15, 0xa5, 0x40, 0x32, 0x91, 0xaa, 0x70, 0xb0, 0x18, 0x2c, 0xc7, 0x1c, 0x6d, 0xfa,
	0x1c, 0x86, 0x8d, 0x49, 0x86, 0x1e, 0x06, 0xad, 0xc3, 0x3e, 0x43, 0xc0, 0x55, 0x94, 0x97, 0x92,
	0xce, 0xc0, 0x8b, 0xa5, 0x43, 0x78, 0xb1, 0x34, 0x3d, 0xa4, 0xa8, 0x44, 0xe8, 0x2d, 0xfc, 0xa5,
	0xc7, 0xd1, 0xa6, 0x67, 0x40, 0x52, 0x55, 0x89, 0xd0, 0x5f, 0xf8, 0xcb, 0xc9, 0xfa, 0x64, 0x65,
	0x48, 0xdc, 0x8e, 0xe5, 0x98, 0x64, 0x02, 0x66, 0xb6, 0x25, 0x57, 0xba, 0xc8, 0x33, 0xad, 0x68,
	0x08, 0x23, 0x5d, 0x47, 0x91, 0xd2, 0x1a, 0xfb, 0x1f, 0xf3, 0xd6, 0xa5, 0x4f, 0xc0, 0x4f, 0xf5,
	0x8d, 0xa3, 0x64, 0x4c, 0x7a, 0x06, 0x41, 0x89, 0xe8, 0xd0, 0x5f, 0x0c, 0x96, 0x93, 0xf5, 0x04,
	0x87, 0xb8, 0x86, 0x2e, 0xc5, 0xde, 0x41, 0xf0, 0xa9, 0x14, 0x51, 0xa2, 0xfe, 0xc5, 0x1a, 0x37,
	0xf7, 0x7a, 0x9b, 0x53, 0x20, 0x51, 0x2e, 0x15, 0x36, 0x1c, 0x73, 0xb4, 0x59, 0x04, 0x33, 0xdb,
	0xe1, 0x51, 0x24, 0xcf, 0x61, 0x94, 0x23, 0x5a, 0xbb, 0x53, 0x58, 0x96, 0xae, 0x63, 0x9b, 0x63,
	0x6f, 0x81, 0x5c, 0x89, 0x24, 0xa1, 0x2f, 0x60, 0x6c, 0x43, 0xdf, 0x3a, 0xae, 0xc7, 0x36, 0xb0,
	0x41, 0xc6, 0xa2, 0xbc, 0xd1, 0x78, 0xe7, 0x31, 0x47, 0x9b, 0x6d, 0x61, 0x6a, 0x80, 0x8f, 0xe2,
	0x46, 0x81, 0xfc, 0xd4, 0x79, 0xd6, 0x6e, 0x6b, 0x6c, 0x76, 0x0a, 0xe4, 0x72, 0xb7, 0x39, 0x78,
	0x63, 0xf6, 0x12, 0x82, 0xcb, 0xdd, 0xd6, 0xdd, 0xe8, 0x6f, 0xc5, 0xb0, 0xdf, 0x03, 0x80, 0x6b,
	0x55, 0x36, 0xaa, 0xdc, 0x64, 0x3f, 0x72, 0x43, 0xa2, 0x51, 0xa5, 0x8e, 0xf3, 0xcc, 0x55, 0xb5,
	0x2e, 0x3d, 0x85, 0xa0, 0x2e, 0xaa, 0xd8, 0x9d, 0x9d, 0x70, 0xe7, 0x19, 0x72, 0x45, 0x6c, 0x1f,
	0x92, 0x70, 0x63, 0x9a, 0x48, 0x1d, 0xcb, 0x90, 0xd8, 0x48, 0x1d, 0xb7, 0xeb, 0x37, 0xe1, 0xb0,
	0x5b, 0xbf, 0x31, 0x93, 0xec, 0x43, 0xeb, 0x30, 0xc0, 0xca, 0xd6, 0x35, 0x99, 0xf6, 0xf0, 0x23,
	0x9b, 0x69, 0x6f, 0x3d, 0x82, 0xe1, 0x87, 0xb4, 0xa8, 0x76, 0xeb, 0x5f, 0x04, 0xe0, 0xba, 0x4e,
	0x0d, 0xf1, 0x38, 0x52, 0x74, 0x0d, 0xd3, 0xab, 0x52, 0x89, 0x4a, 0x39, 0x99, 0xf7, 0xf5, 0x34,
	0x7f, 0xd6, 0x17, 0x97, 0x3b, 0x36, 0x3b, 0x32, 0x98, 0x2f, 0x85, 0x7c, 0x18, 0x66, 0x05, 0xc0,
	0x95, 0x90, 0x0e, 0x31, 0xc6, 0x22, 0x73, 0xf3, 0xbb, 0xea, 0x2f, 0x60, 0xfa, 0x5e, 0x25, 0xaa,
	0x9b, 0xf1, 0x7f, 0x44, 0xb7, 0x89, 0x93, 0x7e, 0x5f, 0x73, 0x0e, 0xb3, 0x2f, 0xe9, 0xfe, 0x26,
	0x0f, 0xc0, 0xb8, 0x4d, 0x1c, 0xe2, 0x80, 0xd7, 0x41, 0xfd, 0x05, 0xc0, 0xc7, 0x38, 0x93, 0x7b,
	0x13, 0xac, 0xaa, 0xee, 0x46, 0xb8, 0xdd, 0xef, 0x3d, 0xe3, 0x1c, 0x7c, 0x5e, 0x67, 0xae, 0xd0,
	0x7c, 0x1a, 0xf3, 0xa7, 0x9d, 0xb9, 0x57, 0x46, 0x50, 0xaa, 0x80, 0x49, 0xd4, 0xc3, 0xdc, 0xfe,
	0xa5, 0x6e, 0x75, 0xcc, 0x8e, 0xbe, 0x07, 0xf8, 0xe3, 0x7c, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x07, 0x43, 0xa6, 0x19, 0x4b, 0x05, 0x00, 0x00,
}
