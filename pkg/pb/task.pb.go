// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type CreateTaskRequest struct {
	JobId                *wrappers.StringValue `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	NodeId               *wrappers.StringValue `protobuf:"bytes,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Target               *wrappers.StringValue `protobuf:"bytes,3,opt,name=target" json:"target,omitempty"`
	TaskAction           *wrappers.StringValue `protobuf:"bytes,4,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Directive            *wrappers.StringValue `protobuf:"bytes,5,opt,name=directive" json:"directive,omitempty"`
	FailureAllowed       *wrappers.BoolValue   `protobuf:"bytes,6,opt,name=failure_allowed,json=failureAllowed" json:"failure_allowed,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_8ceab251caf06fbd, []int{0}
}
func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (dst *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(dst, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *CreateTaskRequest) GetNodeId() *wrappers.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *CreateTaskRequest) GetTarget() *wrappers.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *CreateTaskRequest) GetTaskAction() *wrappers.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *CreateTaskRequest) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *CreateTaskRequest) GetFailureAllowed() *wrappers.BoolValue {
	if m != nil {
		return m.FailureAllowed
	}
	return nil
}

func (m *CreateTaskRequest) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

type CreateTaskResponse struct {
	TaskId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId                *wrappers.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_8ceab251caf06fbd, []int{1}
}
func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (dst *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(dst, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetTaskId() *wrappers.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *CreateTaskResponse) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

type RetryTasksRequest struct {
	TaskId               []string `protobuf:"bytes,1,rep,name=task_id,json=taskId" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryTasksRequest) Reset()         { *m = RetryTasksRequest{} }
func (m *RetryTasksRequest) String() string { return proto.CompactTextString(m) }
func (*RetryTasksRequest) ProtoMessage()    {}
func (*RetryTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_8ceab251caf06fbd, []int{2}
}
func (m *RetryTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryTasksRequest.Unmarshal(m, b)
}
func (m *RetryTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryTasksRequest.Marshal(b, m, deterministic)
}
func (dst *RetryTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryTasksRequest.Merge(dst, src)
}
func (m *RetryTasksRequest) XXX_Size() int {
	return xxx_messageInfo_RetryTasksRequest.Size(m)
}
func (m *RetryTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetryTasksRequest proto.InternalMessageInfo

func (m *RetryTasksRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

type RetryTasksResponse struct {
	TaskSet              []*Task  `protobuf:"bytes,1,rep,name=task_set,json=taskSet" json:"task_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryTasksResponse) Reset()         { *m = RetryTasksResponse{} }
func (m *RetryTasksResponse) String() string { return proto.CompactTextString(m) }
func (*RetryTasksResponse) ProtoMessage()    {}
func (*RetryTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_8ceab251caf06fbd, []int{3}
}
func (m *RetryTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryTasksResponse.Unmarshal(m, b)
}
func (m *RetryTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryTasksResponse.Marshal(b, m, deterministic)
}
func (dst *RetryTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryTasksResponse.Merge(dst, src)
}
func (m *RetryTasksResponse) XXX_Size() int {
	return xxx_messageInfo_RetryTasksResponse.Size(m)
}
func (m *RetryTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RetryTasksResponse proto.InternalMessageInfo

func (m *RetryTasksResponse) GetTaskSet() []*Task {
	if m != nil {
		return m.TaskSet
	}
	return nil
}

type Task struct {
	TaskId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId                *wrappers.StringValue `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	TaskAction           *wrappers.StringValue `protobuf:"bytes,3,opt,name=task_action,json=taskAction" json:"task_action,omitempty"`
	Status               *wrappers.StringValue `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	ErrorCode            *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	Directive            *wrappers.StringValue `protobuf:"bytes,6,opt,name=directive" json:"directive,omitempty"`
	Executor             *wrappers.StringValue `protobuf:"bytes,7,opt,name=executor" json:"executor,omitempty"`
	Owner                *wrappers.StringValue `protobuf:"bytes,8,opt,name=owner" json:"owner,omitempty"`
	Target               *wrappers.StringValue `protobuf:"bytes,9,opt,name=target" json:"target,omitempty"`
	NodeId               *wrappers.StringValue `protobuf:"bytes,10,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	CreateTime           *timestamp.Timestamp  `protobuf:"bytes,11,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime           *timestamp.Timestamp  `protobuf:"bytes,12,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
	FailureAllowed       *wrappers.BoolValue   `protobuf:"bytes,13,opt,name=failure_allowed,json=failureAllowed" json:"failure_allowed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_8ceab251caf06fbd, []int{4}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskId() *wrappers.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *Task) GetJobId() *wrappers.StringValue {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *Task) GetTaskAction() *wrappers.StringValue {
	if m != nil {
		return m.TaskAction
	}
	return nil
}

func (m *Task) GetStatus() *wrappers.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Task) GetErrorCode() *wrappers.UInt32Value {
	if m != nil {
		return m.ErrorCode
	}
	return nil
}

func (m *Task) GetDirective() *wrappers.StringValue {
	if m != nil {
		return m.Directive
	}
	return nil
}

func (m *Task) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *Task) GetOwner() *wrappers.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Task) GetTarget() *wrappers.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Task) GetNodeId() *wrappers.StringValue {
	if m != nil {
		return m.NodeId
	}
	return nil
}

func (m *Task) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetStatusTime() *timestamp.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

func (m *Task) GetFailureAllowed() *wrappers.BoolValue {
	if m != nil {
		return m.FailureAllowed
	}
	return nil
}

type DescribeTasksRequest struct {
	TaskId   []string              `protobuf:"bytes,1,rep,name=task_id,json=taskId" json:"task_id,omitempty"`
	JobId    []string              `protobuf:"bytes,2,rep,name=job_id,json=jobId" json:"job_id,omitempty"`
	Executor *wrappers.StringValue `protobuf:"bytes,3,opt,name=executor" json:"executor,omitempty"`
	Target   *wrappers.StringValue `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	Status   []string              `protobuf:"bytes,5,rep,name=status" json:"status,omitempty"`
	// default is 20, max value is 200
	Limit uint32 `protobuf:"varint,6,opt,name=limit" json:"limit,omitempty"`
	// default is 0
	Offset               uint32                `protobuf:"varint,7,opt,name=offset" json:"offset,omitempty"`
	SearchWord           *wrappers.StringValue `protobuf:"bytes,8,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DescribeTasksRequest) Reset()         { *m = DescribeTasksRequest{} }
func (m *DescribeTasksRequest) String() string { return proto.CompactTextString(m) }
func (*DescribeTasksRequest) ProtoMessage()    {}
func (*DescribeTasksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_8ceab251caf06fbd, []int{5}
}
func (m *DescribeTasksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeTasksRequest.Unmarshal(m, b)
}
func (m *DescribeTasksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeTasksRequest.Marshal(b, m, deterministic)
}
func (dst *DescribeTasksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeTasksRequest.Merge(dst, src)
}
func (m *DescribeTasksRequest) XXX_Size() int {
	return xxx_messageInfo_DescribeTasksRequest.Size(m)
}
func (m *DescribeTasksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeTasksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeTasksRequest proto.InternalMessageInfo

func (m *DescribeTasksRequest) GetTaskId() []string {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *DescribeTasksRequest) GetJobId() []string {
	if m != nil {
		return m.JobId
	}
	return nil
}

func (m *DescribeTasksRequest) GetExecutor() *wrappers.StringValue {
	if m != nil {
		return m.Executor
	}
	return nil
}

func (m *DescribeTasksRequest) GetTarget() *wrappers.StringValue {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *DescribeTasksRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeTasksRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeTasksRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DescribeTasksRequest) GetSearchWord() *wrappers.StringValue {
	if m != nil {
		return m.SearchWord
	}
	return nil
}

type DescribeTasksResponse struct {
	TotalCount           uint32   `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	TaskSet              []*Task  `protobuf:"bytes,2,rep,name=task_set,json=taskSet" json:"task_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescribeTasksResponse) Reset()         { *m = DescribeTasksResponse{} }
func (m *DescribeTasksResponse) String() string { return proto.CompactTextString(m) }
func (*DescribeTasksResponse) ProtoMessage()    {}
func (*DescribeTasksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_task_8ceab251caf06fbd, []int{6}
}
func (m *DescribeTasksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescribeTasksResponse.Unmarshal(m, b)
}
func (m *DescribeTasksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescribeTasksResponse.Marshal(b, m, deterministic)
}
func (dst *DescribeTasksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescribeTasksResponse.Merge(dst, src)
}
func (m *DescribeTasksResponse) XXX_Size() int {
	return xxx_messageInfo_DescribeTasksResponse.Size(m)
}
func (m *DescribeTasksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DescribeTasksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DescribeTasksResponse proto.InternalMessageInfo

func (m *DescribeTasksResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeTasksResponse) GetTaskSet() []*Task {
	if m != nil {
		return m.TaskSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTaskRequest)(nil), "openpitrix.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "openpitrix.CreateTaskResponse")
	proto.RegisterType((*RetryTasksRequest)(nil), "openpitrix.RetryTasksRequest")
	proto.RegisterType((*RetryTasksResponse)(nil), "openpitrix.RetryTasksResponse")
	proto.RegisterType((*Task)(nil), "openpitrix.Task")
	proto.RegisterType((*DescribeTasksRequest)(nil), "openpitrix.DescribeTasksRequest")
	proto.RegisterType((*DescribeTasksResponse)(nil), "openpitrix.DescribeTasksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskManager service

type TaskManagerClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error)
	RetryTasks(ctx context.Context, in *RetryTasksRequest, opts ...grpc.CallOption) (*RetryTasksResponse, error)
}

type taskManagerClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagerClient(cc *grpc.ClientConn) TaskManagerClient {
	return &taskManagerClient{cc}
}

func (c *taskManagerClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/CreateTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) DescribeTasks(ctx context.Context, in *DescribeTasksRequest, opts ...grpc.CallOption) (*DescribeTasksResponse, error) {
	out := new(DescribeTasksResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/DescribeTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagerClient) RetryTasks(ctx context.Context, in *RetryTasksRequest, opts ...grpc.CallOption) (*RetryTasksResponse, error) {
	out := new(RetryTasksResponse)
	err := grpc.Invoke(ctx, "/openpitrix.TaskManager/RetryTasks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskManager service

type TaskManagerServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	DescribeTasks(context.Context, *DescribeTasksRequest) (*DescribeTasksResponse, error)
	RetryTasks(context.Context, *RetryTasksRequest) (*RetryTasksResponse, error)
}

func RegisterTaskManagerServer(s *grpc.Server, srv TaskManagerServer) {
	s.RegisterService(&_TaskManager_serviceDesc, srv)
}

func _TaskManager_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_DescribeTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).DescribeTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/DescribeTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).DescribeTasks(ctx, req.(*DescribeTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManager_RetryTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetryTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagerServer).RetryTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.TaskManager/RetryTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagerServer).RetryTasks(ctx, req.(*RetryTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.TaskManager",
	HandlerType: (*TaskManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskManager_CreateTask_Handler,
		},
		{
			MethodName: "DescribeTasks",
			Handler:    _TaskManager_DescribeTasks_Handler,
		},
		{
			MethodName: "RetryTasks",
			Handler:    _TaskManager_RetryTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}

func init() { proto.RegisterFile("task.proto", fileDescriptor_task_8ceab251caf06fbd) }

var fileDescriptor_task_8ceab251caf06fbd = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcb, 0x4e, 0x1b, 0x49,
	0x14, 0x95, 0x9f, 0xe0, 0xdb, 0xe3, 0x61, 0x28, 0xc1, 0x4c, 0xcb, 0x62, 0xc0, 0xd3, 0x2b, 0xc4,
	0x80, 0x3d, 0x63, 0x66, 0xa4, 0x08, 0x94, 0x85, 0x71, 0x36, 0x56, 0x94, 0x8d, 0x21, 0x89, 0x94,
	0x8d, 0x55, 0xee, 0xbe, 0x6e, 0x1a, 0x9a, 0xae, 0x4e, 0x55, 0x19, 0x93, 0x5d, 0x94, 0x45, 0x3e,
	0x80, 0x7c, 0x46, 0x3e, 0x21, 0x9f, 0x91, 0x6d, 0x96, 0xf9, 0x8a, 0xac, 0xa2, 0xaa, 0x6a, 0xe3,
	0x36, 0x86, 0xd0, 0xde, 0x64, 0x65, 0xb9, 0xee, 0x39, 0x55, 0x75, 0xeb, 0x9c, 0x7b, 0x1a, 0x40,
	0x52, 0x71, 0xde, 0x88, 0x39, 0x93, 0x8c, 0x00, 0x8b, 0x31, 0x8a, 0x03, 0xc9, 0x83, 0xab, 0xda,
	0xa6, 0xcf, 0x98, 0x1f, 0x62, 0x53, 0x57, 0x06, 0xa3, 0x61, 0x73, 0xcc, 0x69, 0x1c, 0x23, 0x17,
	0x06, 0x5b, 0xdb, 0xba, 0x5d, 0x97, 0xc1, 0x05, 0x0a, 0x49, 0x2f, 0xe2, 0x04, 0xb0, 0x91, 0x00,
	0x68, 0x1c, 0x34, 0x69, 0x14, 0x31, 0x49, 0x65, 0xc0, 0xa2, 0x09, 0x7d, 0x57, 0xff, 0xb8, 0x7b,
	0x3e, 0x46, 0x7b, 0x62, 0x4c, 0x7d, 0x1f, 0x79, 0x93, 0xc5, 0x1a, 0x31, 0x8f, 0x76, 0x3e, 0x16,
	0x60, 0xb5, 0xc3, 0x91, 0x4a, 0x3c, 0xa1, 0xe2, 0xbc, 0x87, 0xaf, 0x47, 0x28, 0x24, 0xd9, 0x87,
	0xf2, 0x19, 0x1b, 0xf4, 0x03, 0xcf, 0xce, 0xd5, 0x73, 0xdb, 0x56, 0x6b, 0xa3, 0x61, 0x8e, 0x6c,
	0x4c, 0xee, 0xd4, 0x38, 0x96, 0x3c, 0x88, 0xfc, 0x17, 0x34, 0x1c, 0x61, 0xaf, 0x74, 0xc6, 0x06,
	0x5d, 0x8f, 0xfc, 0x0f, 0x4b, 0x11, 0xf3, 0x50, 0xb1, 0xf2, 0x19, 0x58, 0x65, 0x05, 0xee, 0x7a,
	0xe4, 0x3f, 0x28, 0x4b, 0xca, 0x7d, 0x94, 0x76, 0x21, 0x0b, 0xcb, 0x60, 0xc9, 0x63, 0xb0, 0xd4,
	0xf3, 0xf6, 0xa9, 0xab, 0xba, 0xb1, 0x8b, 0x19, 0xa8, 0x5a, 0x8f, 0xb6, 0xc6, 0x93, 0x03, 0xa8,
	0x78, 0x01, 0x47, 0x57, 0x06, 0x97, 0x68, 0x97, 0x32, 0x90, 0xa7, 0x70, 0xd2, 0x81, 0x95, 0x21,
	0x0d, 0xc2, 0x11, 0xc7, 0x3e, 0x0d, 0x43, 0x36, 0x46, 0xcf, 0x2e, 0xeb, 0x1d, 0x6a, 0x73, 0x3b,
	0x1c, 0x31, 0x16, 0x1a, 0xfe, 0xaf, 0x09, 0xa5, 0x6d, 0x18, 0xaa, 0x6b, 0x21, 0xa9, 0x1c, 0x09,
	0x7b, 0x29, 0x4b, 0xd7, 0x06, 0xeb, 0xbc, 0xcd, 0x01, 0x49, 0xab, 0x25, 0x62, 0x16, 0x09, 0x54,
	0x2f, 0xaf, 0x1f, 0x23, 0xa3, 0x5e, 0x65, 0x05, 0xee, 0x7a, 0x29, 0x95, 0xf3, 0x99, 0x55, 0x76,
	0x76, 0x61, 0xb5, 0x87, 0x92, 0xbf, 0x51, 0x17, 0x10, 0x13, 0xbf, 0xfc, 0x91, 0xbe, 0x40, 0x61,
	0xbb, 0x32, 0x39, 0xc2, 0x69, 0x03, 0x49, 0xa3, 0x93, 0xfb, 0xfe, 0x0d, 0xcb, 0x1a, 0x2e, 0x50,
	0x6a, 0xbc, 0xd5, 0xfa, 0xad, 0x31, 0x1d, 0x90, 0x86, 0xee, 0x4d, 0x6f, 0x78, 0x8c, 0xd2, 0xf9,
	0x56, 0x82, 0xa2, 0x5a, 0xf9, 0x99, 0x5d, 0xde, 0xb6, 0x57, 0x61, 0x41, 0x7b, 0x4d, 0xd5, 0x2d,
	0x66, 0x57, 0x97, 0x1c, 0x02, 0x20, 0xe7, 0x8c, 0xf7, 0x5d, 0xe6, 0xdd, 0xef, 0xca, 0xe7, 0xdd,
	0x48, 0xee, 0xb7, 0x12, 0x57, 0x6a, 0x7c, 0x87, 0x79, 0x38, 0xeb, 0xe8, 0xf2, 0x62, 0x8e, 0x7e,
	0x04, 0xcb, 0x78, 0x85, 0xee, 0x48, 0x32, 0x9e, 0xc9, 0x8e, 0x37, 0x68, 0xd2, 0x82, 0x12, 0x1b,
	0x47, 0xc8, 0xed, 0xe5, 0x2c, 0x6f, 0xab, 0xa1, 0xa9, 0x81, 0xaf, 0x2c, 0x30, 0xf0, 0xa9, 0x74,
	0x81, 0x05, 0xd2, 0xe5, 0x10, 0x2c, 0x57, 0x0f, 0x4c, 0x5f, 0xa5, 0xa8, 0x6d, 0xdd, 0x33, 0xa8,
	0x27, 0x93, 0x88, 0xed, 0x81, 0x81, 0xab, 0x05, 0x45, 0x36, 0xd2, 0x18, 0xf2, 0x2f, 0x0f, 0x93,
	0x0d, 0x5c, 0x93, 0xef, 0x88, 0x89, 0xea, 0xa2, 0x31, 0xe1, 0x7c, 0xca, 0xc3, 0xda, 0x13, 0x14,
	0x2e, 0x0f, 0x06, 0x98, 0x69, 0xe2, 0xc8, 0x7a, 0xca, 0xee, 0x6a, 0x3d, 0x31, 0x74, 0x5a, 0xe2,
	0xc2, 0x42, 0x12, 0x4f, 0xe5, 0x2a, 0x2e, 0x20, 0xd7, 0xef, 0x37, 0x13, 0x50, 0x32, 0xd7, 0x4b,
	0x3c, 0xbe, 0x06, 0xa5, 0x30, 0xb8, 0x08, 0xa4, 0xb6, 0x68, 0xb5, 0x67, 0xfe, 0x28, 0x34, 0x1b,
	0x0e, 0x55, 0x1c, 0x2c, 0xe9, 0xe5, 0xe4, 0x9f, 0x1a, 0x43, 0x81, 0x94, 0xbb, 0xa7, 0xfd, 0x31,
	0xe3, 0x5e, 0x26, 0x93, 0x81, 0x21, 0xbc, 0x64, 0xdc, 0x73, 0x10, 0xd6, 0x6f, 0x3d, 0x5e, 0x12,
	0x40, 0x5b, 0x60, 0x49, 0x26, 0x69, 0xd8, 0x77, 0xd9, 0x28, 0x92, 0x3a, 0x4e, 0xaa, 0x3d, 0xd0,
	0x4b, 0x1d, 0xb5, 0x32, 0x93, 0x50, 0xf9, 0x07, 0x12, 0xaa, 0xf5, 0x25, 0x0f, 0x96, 0x5a, 0x79,
	0x46, 0x23, 0xea, 0x23, 0x27, 0x4f, 0x01, 0xa6, 0x21, 0x4d, 0xfe, 0x4c, 0x13, 0xe7, 0x3e, 0xb5,
	0xb5, 0xcd, 0xfb, 0xca, 0xc9, 0x55, 0xdf, 0xe7, 0xa0, 0x3a, 0xd3, 0x04, 0xa9, 0xa7, 0x19, 0x77,
	0x99, 0xa3, 0xf6, 0xd7, 0x0f, 0x10, 0x66, 0x5b, 0xe7, 0x9f, 0xeb, 0xf6, 0x06, 0xa9, 0x79, 0x49,
	0xad, 0xae, 0x5a, 0x11, 0xf5, 0x71, 0x20, 0x4f, 0xeb, 0xc3, 0x20, 0x94, 0xc8, 0xdf, 0x7d, 0xfe,
	0xfa, 0x21, 0x6f, 0x91, 0x4a, 0xf3, 0xf2, 0xdf, 0xa6, 0x2e, 0x92, 0x31, 0xc0, 0x34, 0xca, 0x67,
	0xbb, 0x9a, 0xfb, 0x20, 0xcc, 0x76, 0x35, 0xff, 0x05, 0x70, 0x76, 0xae, 0xdb, 0x55, 0x62, 0x71,
	0x55, 0x30, 0x67, 0xeb, 0xf3, 0xd6, 0x9c, 0x95, 0x9b, 0xf3, 0x9a, 0xba, 0x78, 0x90, 0xdb, 0x39,
	0x2a, 0xbe, 0xca, 0xc7, 0x83, 0x41, 0x59, 0xab, 0xbd, 0xff, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x59,
	0xc5, 0xf9, 0xc0, 0x56, 0x09, 0x00, 0x00,
}
