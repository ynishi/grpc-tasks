// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tasks.proto

/*
Package tasks is a generated protocol buffer package.

It is generated from these files:
	tasks.proto

It has these top-level messages:
	TaskResponse
	TaskRequest
	Task
	Tag
*/
package tasks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type TaskResponse_Code int32

const (
	TaskResponse_SUCCEED TaskResponse_Code = 0
	TaskResponse_FAILED  TaskResponse_Code = 1
)

var TaskResponse_Code_name = map[int32]string{
	0: "SUCCEED",
	1: "FAILED",
}
var TaskResponse_Code_value = map[string]int32{
	"SUCCEED": 0,
	"FAILED":  1,
}

func (x TaskResponse_Code) String() string {
	return proto.EnumName(TaskResponse_Code_name, int32(x))
}
func (TaskResponse_Code) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type TaskResponse struct {
	Id   int32                      `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Done *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=done" json:"done,omitempty"`
}

func (m *TaskResponse) Reset()                    { *m = TaskResponse{} }
func (m *TaskResponse) String() string            { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()               {}
func (*TaskResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TaskResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskResponse) GetDone() *google_protobuf.Timestamp {
	if m != nil {
		return m.Done
	}
	return nil
}

type TaskRequest struct {
	Id  int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Tag string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
}

func (m *TaskRequest) Reset()                    { *m = TaskRequest{} }
func (m *TaskRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskRequest) ProtoMessage()               {}
func (*TaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TaskRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type Task struct {
	Name    string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id      int32                      `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Title   string                     `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Tags    []*Tag                     `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	DueTime *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=due_time,json=dueTime" json:"due_time,omitempty"`
	Added   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=added" json:"added,omitempty"`
	Updated *google_protobuf.Timestamp `protobuf:"bytes,6,opt,name=updated" json:"updated,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Task) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Task) GetDueTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.DueTime
	}
	return nil
}

func (m *Task) GetAdded() *google_protobuf.Timestamp {
	if m != nil {
		return m.Added
	}
	return nil
}

func (m *Task) GetUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

type Tag struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id   int32  `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*TaskResponse)(nil), "tasks.TaskResponse")
	proto.RegisterType((*TaskRequest)(nil), "tasks.TaskRequest")
	proto.RegisterType((*Task)(nil), "tasks.Task")
	proto.RegisterType((*Tag)(nil), "tasks.Tag")
	proto.RegisterEnum("tasks.TaskResponse_Code", TaskResponse_Code_name, TaskResponse_Code_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskService service

type TaskServiceClient interface {
	GetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error)
	ListTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (TaskService_ListTaskClient, error)
	AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskResponse, error)
	DoneTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskResponse, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) GetTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := grpc.Invoke(ctx, "/tasks.TaskService/GetTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ListTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (TaskService_ListTaskClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskService_serviceDesc.Streams[0], c.cc, "/tasks.TaskService/ListTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServiceListTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskService_ListTaskClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type taskServiceListTaskClient struct {
	grpc.ClientStream
}

func (x *taskServiceListTaskClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskServiceClient) AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := grpc.Invoke(ctx, "/tasks.TaskService/AddTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) DoneTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := grpc.Invoke(ctx, "/tasks.TaskService/DoneTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceServer interface {
	GetTask(context.Context, *TaskRequest) (*Task, error)
	ListTask(*TaskRequest, TaskService_ListTaskServer) error
	AddTask(context.Context, *Task) (*TaskResponse, error)
	DoneTask(context.Context, *Task) (*TaskResponse, error)
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks.TaskService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).GetTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ListTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskServiceServer).ListTask(m, &taskServiceListTaskServer{stream})
}

type TaskService_ListTaskServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type taskServiceListTaskServer struct {
	grpc.ServerStream
}

func (x *taskServiceListTaskServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskService_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks.TaskService/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).AddTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_DoneTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).DoneTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks.TaskService/DoneTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).DoneTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tasks.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTask",
			Handler:    _TaskService_GetTask_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _TaskService_AddTask_Handler,
		},
		{
			MethodName: "DoneTask",
			Handler:    _TaskService_DoneTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListTask",
			Handler:       _TaskService_ListTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tasks.proto",
}

func init() { proto.RegisterFile("tasks.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0xf1, 0x3f, 0x0c, 0xe3, 0xaa, 0x42, 0xd3, 0x1e, 0x2c, 0x1f, 0x0a, 0xf2, 0x89, 0x4a,
	0xad, 0x8d, 0x68, 0xfb, 0x00, 0x08, 0x68, 0x55, 0x89, 0x93, 0x71, 0xce, 0xd1, 0x92, 0x9d, 0x58,
	0x16, 0xe0, 0x75, 0xd8, 0x75, 0x5e, 0x32, 0x0f, 0x95, 0xc8, 0xbb, 0x80, 0x88, 0x72, 0x80, 0xdb,
	0xcc, 0xec, 0xef, 0x9b, 0x4f, 0xf3, 0x69, 0x21, 0x50, 0x4c, 0x6e, 0x65, 0x52, 0x1f, 0x84, 0x12,
	0xe8, 0xe9, 0x26, 0x1a, 0x16, 0x42, 0x14, 0x3b, 0x4a, 0xf5, 0x70, 0xd3, 0x3c, 0xa6, 0xaa, 0xdc,
	0x93, 0x54, 0x6c, 0x5f, 0x1b, 0x2e, 0x16, 0xf0, 0x29, 0x67, 0x72, 0x9b, 0x91, 0xac, 0x45, 0x25,
	0x09, 0x3f, 0x83, 0x5d, 0xf2, 0xd0, 0x1a, 0x59, 0x63, 0x2f, 0xb3, 0x4b, 0x8e, 0x09, 0xb8, 0x5c,
	0x54, 0x14, 0xda, 0x23, 0x6b, 0x1c, 0x4c, 0xa3, 0xc4, 0xec, 0x4b, 0x4e, 0xfb, 0x92, 0xfc, 0xb4,
	0x2f, 0xd3, 0x5c, 0x3c, 0x04, 0x77, 0x2e, 0x38, 0x61, 0x00, 0xfe, 0xfa, 0x6e, 0x3e, 0x5f, 0x2e,
	0x17, 0x83, 0x0e, 0x02, 0x74, 0xff, 0xce, 0xfe, 0xaf, 0x96, 0x8b, 0x81, 0x15, 0xa7, 0x10, 0x18,
	0xc3, 0xa7, 0x86, 0xa4, 0xfa, 0xe0, 0x37, 0x00, 0x47, 0xb1, 0x42, 0xdb, 0xf5, 0xb3, 0xb6, 0x8c,
	0x5f, 0x2d, 0x70, 0x5b, 0x05, 0x22, 0xb8, 0x15, 0xdb, 0x93, 0x86, 0xfb, 0x99, 0xae, 0x8f, 0x72,
	0xfb, 0x2c, 0xff, 0x0a, 0x9e, 0x2a, 0xd5, 0x8e, 0x42, 0x47, 0x43, 0xa6, 0xc1, 0x6f, 0xe0, 0x2a,
	0x56, 0xc8, 0xd0, 0x1f, 0x39, 0xe3, 0x60, 0x0a, 0x89, 0x09, 0x2a, 0x67, 0x45, 0xa6, 0xe7, 0xf8,
	0x07, 0x7a, 0xbc, 0xa1, 0xfb, 0x36, 0x9b, 0xd0, 0xbd, 0x7a, 0xa8, 0xcf, 0x1b, 0x6a, 0x3b, 0x9c,
	0x80, 0xc7, 0x38, 0x27, 0x1e, 0x7a, 0x57, 0x35, 0x06, 0xc4, 0xdf, 0xe0, 0x37, 0x35, 0x67, 0x8a,
	0x78, 0xd8, 0xbd, 0xee, 0x73, 0x44, 0xe3, 0xef, 0xe0, 0xe4, 0xac, 0xb8, 0xe5, 0xfe, 0xe9, 0x8b,
	0x65, 0xe2, 0x5d, 0xd3, 0xe1, 0xb9, 0x7c, 0x20, 0xfc, 0x01, 0xfe, 0x3f, 0x52, 0x26, 0xbe, 0xf3,
	0xd9, 0xe7, 0xf4, 0xa3, 0xe0, 0x62, 0x16, 0x77, 0x30, 0x85, 0xde, 0xaa, 0x94, 0xb7, 0xe2, 0x13,
	0x0b, 0x7f, 0x82, 0x3f, 0xe3, 0x5c, 0xf3, 0x97, 0x6f, 0xd1, 0x97, 0x77, 0x62, 0xf3, 0xb5, 0xe2,
	0x0e, 0x26, 0xd0, 0x5b, 0x88, 0x8a, 0x6e, 0xe5, 0x37, 0x5d, 0x9d, 0xca, 0xaf, 0xb7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xaf, 0x31, 0x86, 0x10, 0xda, 0x02, 0x00, 0x00,
}
