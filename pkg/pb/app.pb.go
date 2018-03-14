// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	app.proto
	job.proto
	pilot.proto
	repo.proto
	repo_indexer.proto
	runtime_env.proto
	task.proto

It has these top-level messages:
	CreateAppRequest
	CreateAppResponse
	ModifyAppRequest
	ModifyAppResponse
	DeleteAppRequest
	DeleteAppResponse
	App
	DescribeAppsRequest
	DescribeAppsResponse
	CreateJobRequest
	CreateJobResponse
	Job
	DescribeJobsRequest
	DescribeJobsResponse
	HandleSubtaskRequest
	HandleSubtaskResponse
	GetSubtaskStatusRequest
	SubtaskStatus
	GetSubtaskStatusResponse
	CreateRepoRequest
	CreateRepoResponse
	ModifyRepoRequest
	ModifyRepoResponse
	DeleteRepoRequest
	DeleteRepoResponse
	Repo
	DescribeReposRequest
	DescribeReposResponse
	IndexRepoRequest
	IndexRepoResponse
	RepoTask
	DescribeRepoTasksRequest
	DescribeRepoTasksResponse
	RuntimeEnv
	CreateRuntimeEnvRequest
	CreateRuntimeEnvResponse
	DescribeRuntimeEnvsRequest
	DescribeRuntimeEnvsResponse
	ModifyRuntimeEnvRequest
	ModifyRuntimeEnvResponse
	DeleteRuntimeEnvRequest
	DeleteRuntimeEnvResponse
	RuntimeEnvCredential
	CreateRuntimeEnvCredentialRequset
	CreateRuntimeEnvCredentialResponse
	DescribeRuntimeEnvCredentialsRequset
	DescribeRuntimeEnvCredentialsResponse
	ModifyRuntimeEnvCredentialRequest
	ModifyRuntimeEnvCredentialResponse
	DeleteRuntimeEnvCredentialRequset
	DeleteRuntimeEnvCredentialResponse
	AttachCredentialToRuntimeEnvRequset
	AttachCredentialToRuntimeEnvResponse
	DetachCredentialFromRuntimeEnvRequset
	DetachCredentialFromRuntimeEnvResponse
	CreateTaskRequest
	CreateTaskResponse
	Task
	DescribeTasksRequest
	DescribeTasksResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

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

type CreateAppRequest struct {
	X           *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=_" json:"_,omitempty"`
	Name        *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RepoId      *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Owner       *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
	ChartName   *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=chart_name,json=chartName" json:"chart_name,omitempty"`
	Description *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Home        *google_protobuf.StringValue `protobuf:"bytes,8,opt,name=home" json:"home,omitempty"`
	Icon        *google_protobuf.StringValue `protobuf:"bytes,9,opt,name=icon" json:"icon,omitempty"`
	Screenshots *google_protobuf.StringValue `protobuf:"bytes,10,opt,name=screenshots" json:"screenshots,omitempty"`
	Maintainers *google_protobuf.StringValue `protobuf:"bytes,11,opt,name=maintainers" json:"maintainers,omitempty"`
	Sources     *google_protobuf.StringValue `protobuf:"bytes,12,opt,name=sources" json:"sources,omitempty"`
	Readme      *google_protobuf.StringValue `protobuf:"bytes,13,opt,name=readme" json:"readme,omitempty"`
}

func (m *CreateAppRequest) Reset()                    { *m = CreateAppRequest{} }
func (m *CreateAppRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateAppRequest) ProtoMessage()               {}
func (*CreateAppRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateAppRequest) GetX() *google_protobuf.StringValue {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *CreateAppRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateAppRequest) GetRepoId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *CreateAppRequest) GetOwner() *google_protobuf.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *CreateAppRequest) GetChartName() *google_protobuf.StringValue {
	if m != nil {
		return m.ChartName
	}
	return nil
}

func (m *CreateAppRequest) GetDescription() *google_protobuf.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CreateAppRequest) GetHome() *google_protobuf.StringValue {
	if m != nil {
		return m.Home
	}
	return nil
}

func (m *CreateAppRequest) GetIcon() *google_protobuf.StringValue {
	if m != nil {
		return m.Icon
	}
	return nil
}

func (m *CreateAppRequest) GetScreenshots() *google_protobuf.StringValue {
	if m != nil {
		return m.Screenshots
	}
	return nil
}

func (m *CreateAppRequest) GetMaintainers() *google_protobuf.StringValue {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

func (m *CreateAppRequest) GetSources() *google_protobuf.StringValue {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *CreateAppRequest) GetReadme() *google_protobuf.StringValue {
	if m != nil {
		return m.Readme
	}
	return nil
}

type CreateAppResponse struct {
	App *App `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
}

func (m *CreateAppResponse) Reset()                    { *m = CreateAppResponse{} }
func (m *CreateAppResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateAppResponse) ProtoMessage()               {}
func (*CreateAppResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateAppResponse) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

type ModifyAppRequest struct {
	AppId       *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Name        *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RepoId      *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Owner       *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=owner" json:"owner,omitempty"`
	ChartName   *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=chart_name,json=chartName" json:"chart_name,omitempty"`
	Description *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	Home        *google_protobuf.StringValue `protobuf:"bytes,8,opt,name=home" json:"home,omitempty"`
	Icon        *google_protobuf.StringValue `protobuf:"bytes,9,opt,name=icon" json:"icon,omitempty"`
	Screenshots *google_protobuf.StringValue `protobuf:"bytes,10,opt,name=screenshots" json:"screenshots,omitempty"`
	Maintainers *google_protobuf.StringValue `protobuf:"bytes,11,opt,name=maintainers" json:"maintainers,omitempty"`
	Sources     *google_protobuf.StringValue `protobuf:"bytes,12,opt,name=sources" json:"sources,omitempty"`
	Readme      *google_protobuf.StringValue `protobuf:"bytes,13,opt,name=readme" json:"readme,omitempty"`
}

func (m *ModifyAppRequest) Reset()                    { *m = ModifyAppRequest{} }
func (m *ModifyAppRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyAppRequest) ProtoMessage()               {}
func (*ModifyAppRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ModifyAppRequest) GetAppId() *google_protobuf.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *ModifyAppRequest) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ModifyAppRequest) GetRepoId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *ModifyAppRequest) GetOwner() *google_protobuf.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ModifyAppRequest) GetChartName() *google_protobuf.StringValue {
	if m != nil {
		return m.ChartName
	}
	return nil
}

func (m *ModifyAppRequest) GetDescription() *google_protobuf.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *ModifyAppRequest) GetHome() *google_protobuf.StringValue {
	if m != nil {
		return m.Home
	}
	return nil
}

func (m *ModifyAppRequest) GetIcon() *google_protobuf.StringValue {
	if m != nil {
		return m.Icon
	}
	return nil
}

func (m *ModifyAppRequest) GetScreenshots() *google_protobuf.StringValue {
	if m != nil {
		return m.Screenshots
	}
	return nil
}

func (m *ModifyAppRequest) GetMaintainers() *google_protobuf.StringValue {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

func (m *ModifyAppRequest) GetSources() *google_protobuf.StringValue {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *ModifyAppRequest) GetReadme() *google_protobuf.StringValue {
	if m != nil {
		return m.Readme
	}
	return nil
}

type ModifyAppResponse struct {
	App *App `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
}

func (m *ModifyAppResponse) Reset()                    { *m = ModifyAppResponse{} }
func (m *ModifyAppResponse) String() string            { return proto.CompactTextString(m) }
func (*ModifyAppResponse) ProtoMessage()               {}
func (*ModifyAppResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ModifyAppResponse) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

type DeleteAppRequest struct {
	AppId *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *DeleteAppRequest) Reset()                    { *m = DeleteAppRequest{} }
func (m *DeleteAppRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteAppRequest) ProtoMessage()               {}
func (*DeleteAppRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteAppRequest) GetAppId() *google_protobuf.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

type DeleteAppResponse struct {
	App *App `protobuf:"bytes,1,opt,name=app" json:"app,omitempty"`
}

func (m *DeleteAppResponse) Reset()                    { *m = DeleteAppResponse{} }
func (m *DeleteAppResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteAppResponse) ProtoMessage()               {}
func (*DeleteAppResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteAppResponse) GetApp() *App {
	if m != nil {
		return m.App
	}
	return nil
}

type App struct {
	AppId       *google_protobuf.StringValue `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	Name        *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	RepoId      *google_protobuf.StringValue `protobuf:"bytes,3,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Description *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Status      *google_protobuf.StringValue `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	Home        *google_protobuf.StringValue `protobuf:"bytes,6,opt,name=home" json:"home,omitempty"`
	Icon        *google_protobuf.StringValue `protobuf:"bytes,7,opt,name=icon" json:"icon,omitempty"`
	Screenshots *google_protobuf.StringValue `protobuf:"bytes,8,opt,name=screenshots" json:"screenshots,omitempty"`
	Maintainers *google_protobuf.StringValue `protobuf:"bytes,9,opt,name=maintainers" json:"maintainers,omitempty"`
	Sources     *google_protobuf.StringValue `protobuf:"bytes,10,opt,name=sources" json:"sources,omitempty"`
	Readme      *google_protobuf.StringValue `protobuf:"bytes,11,opt,name=readme" json:"readme,omitempty"`
	ChartName   *google_protobuf.StringValue `protobuf:"bytes,12,opt,name=chart_name,json=chartName" json:"chart_name,omitempty"`
	Owner       *google_protobuf.StringValue `protobuf:"bytes,13,opt,name=owner" json:"owner,omitempty"`
	CreateTime  *google_protobuf1.Timestamp  `protobuf:"bytes,14,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime  *google_protobuf1.Timestamp  `protobuf:"bytes,15,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *App) Reset()                    { *m = App{} }
func (m *App) String() string            { return proto.CompactTextString(m) }
func (*App) ProtoMessage()               {}
func (*App) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *App) GetAppId() *google_protobuf.StringValue {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *App) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *App) GetRepoId() *google_protobuf.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *App) GetDescription() *google_protobuf.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *App) GetStatus() *google_protobuf.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *App) GetHome() *google_protobuf.StringValue {
	if m != nil {
		return m.Home
	}
	return nil
}

func (m *App) GetIcon() *google_protobuf.StringValue {
	if m != nil {
		return m.Icon
	}
	return nil
}

func (m *App) GetScreenshots() *google_protobuf.StringValue {
	if m != nil {
		return m.Screenshots
	}
	return nil
}

func (m *App) GetMaintainers() *google_protobuf.StringValue {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

func (m *App) GetSources() *google_protobuf.StringValue {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *App) GetReadme() *google_protobuf.StringValue {
	if m != nil {
		return m.Readme
	}
	return nil
}

func (m *App) GetChartName() *google_protobuf.StringValue {
	if m != nil {
		return m.ChartName
	}
	return nil
}

func (m *App) GetOwner() *google_protobuf.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *App) GetCreateTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *App) GetStatusTime() *google_protobuf1.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeAppsRequest struct {
	AppId      []string                     `protobuf:"bytes,1,rep,name=app_id,json=appId" json:"app_id,omitempty"`
	Name       []string                     `protobuf:"bytes,2,rep,name=name" json:"name,omitempty"`
	RepoId     []string                     `protobuf:"bytes,3,rep,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Status     []string                     `protobuf:"bytes,4,rep,name=status" json:"status,omitempty"`
	Owner      []string                     `protobuf:"bytes,5,rep,name=owner" json:"owner,omitempty"`
	ChartName  []string                     `protobuf:"bytes,6,rep,name=chart_name,json=chartName" json:"chart_name,omitempty"`
	SearchWord *google_protobuf.StringValue `protobuf:"bytes,7,opt,name=search_word,json=searchWord" json:"search_word,omitempty"`
	Limit      uint32                       `protobuf:"varint,8,opt,name=limit" json:"limit,omitempty"`
	Offset     uint32                       `protobuf:"varint,9,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeAppsRequest) Reset()                    { *m = DescribeAppsRequest{} }
func (m *DescribeAppsRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeAppsRequest) ProtoMessage()               {}
func (*DescribeAppsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DescribeAppsRequest) GetAppId() []string {
	if m != nil {
		return m.AppId
	}
	return nil
}

func (m *DescribeAppsRequest) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *DescribeAppsRequest) GetRepoId() []string {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *DescribeAppsRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeAppsRequest) GetOwner() []string {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *DescribeAppsRequest) GetChartName() []string {
	if m != nil {
		return m.ChartName
	}
	return nil
}

func (m *DescribeAppsRequest) GetSearchWord() *google_protobuf.StringValue {
	if m != nil {
		return m.SearchWord
	}
	return nil
}

func (m *DescribeAppsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeAppsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeAppsResponse struct {
	TotalCount uint32 `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	AppSet     []*App `protobuf:"bytes,2,rep,name=app_set,json=appSet" json:"app_set,omitempty"`
}

func (m *DescribeAppsResponse) Reset()                    { *m = DescribeAppsResponse{} }
func (m *DescribeAppsResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeAppsResponse) ProtoMessage()               {}
func (*DescribeAppsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DescribeAppsResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeAppsResponse) GetAppSet() []*App {
	if m != nil {
		return m.AppSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateAppRequest)(nil), "openpitrix.CreateAppRequest")
	proto.RegisterType((*CreateAppResponse)(nil), "openpitrix.CreateAppResponse")
	proto.RegisterType((*ModifyAppRequest)(nil), "openpitrix.ModifyAppRequest")
	proto.RegisterType((*ModifyAppResponse)(nil), "openpitrix.ModifyAppResponse")
	proto.RegisterType((*DeleteAppRequest)(nil), "openpitrix.DeleteAppRequest")
	proto.RegisterType((*DeleteAppResponse)(nil), "openpitrix.DeleteAppResponse")
	proto.RegisterType((*App)(nil), "openpitrix.App")
	proto.RegisterType((*DescribeAppsRequest)(nil), "openpitrix.DescribeAppsRequest")
	proto.RegisterType((*DescribeAppsResponse)(nil), "openpitrix.DescribeAppsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AppManager service

type AppManagerClient interface {
	CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error)
	DescribeApps(ctx context.Context, in *DescribeAppsRequest, opts ...grpc.CallOption) (*DescribeAppsResponse, error)
	ModifyApp(ctx context.Context, in *ModifyAppRequest, opts ...grpc.CallOption) (*ModifyAppResponse, error)
	DeleteApp(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*DeleteAppResponse, error)
}

type appManagerClient struct {
	cc *grpc.ClientConn
}

func NewAppManagerClient(cc *grpc.ClientConn) AppManagerClient {
	return &appManagerClient{cc}
}

func (c *appManagerClient) CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*CreateAppResponse, error) {
	out := new(CreateAppResponse)
	err := grpc.Invoke(ctx, "/openpitrix.AppManager/CreateApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerClient) DescribeApps(ctx context.Context, in *DescribeAppsRequest, opts ...grpc.CallOption) (*DescribeAppsResponse, error) {
	out := new(DescribeAppsResponse)
	err := grpc.Invoke(ctx, "/openpitrix.AppManager/DescribeApps", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerClient) ModifyApp(ctx context.Context, in *ModifyAppRequest, opts ...grpc.CallOption) (*ModifyAppResponse, error) {
	out := new(ModifyAppResponse)
	err := grpc.Invoke(ctx, "/openpitrix.AppManager/ModifyApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appManagerClient) DeleteApp(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*DeleteAppResponse, error) {
	out := new(DeleteAppResponse)
	err := grpc.Invoke(ctx, "/openpitrix.AppManager/DeleteApp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppManager service

type AppManagerServer interface {
	CreateApp(context.Context, *CreateAppRequest) (*CreateAppResponse, error)
	DescribeApps(context.Context, *DescribeAppsRequest) (*DescribeAppsResponse, error)
	ModifyApp(context.Context, *ModifyAppRequest) (*ModifyAppResponse, error)
	DeleteApp(context.Context, *DeleteAppRequest) (*DeleteAppResponse, error)
}

func RegisterAppManagerServer(s *grpc.Server, srv AppManagerServer) {
	s.RegisterService(&_AppManager_serviceDesc, srv)
}

func _AppManager_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.AppManager/CreateApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).CreateApp(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppManager_DescribeApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).DescribeApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.AppManager/DescribeApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).DescribeApps(ctx, req.(*DescribeAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppManager_ModifyApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).ModifyApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.AppManager/ModifyApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).ModifyApp(ctx, req.(*ModifyAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppManager_DeleteApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppManagerServer).DeleteApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.AppManager/DeleteApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppManagerServer).DeleteApp(ctx, req.(*DeleteAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.AppManager",
	HandlerType: (*AppManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApp",
			Handler:    _AppManager_CreateApp_Handler,
		},
		{
			MethodName: "DescribeApps",
			Handler:    _AppManager_DescribeApps_Handler,
		},
		{
			MethodName: "ModifyApp",
			Handler:    _AppManager_ModifyApp_Handler,
		},
		{
			MethodName: "DeleteApp",
			Handler:    _AppManager_DeleteApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

func init() { proto.RegisterFile("app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 911 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x4f, 0x8f, 0x1b, 0x35,
	0x14, 0x27, 0x9b, 0xcd, 0xec, 0xe6, 0x65, 0x43, 0x17, 0xd3, 0xc2, 0x10, 0xb6, 0x6c, 0x98, 0xd3,
	0xb2, 0xd0, 0x99, 0x6d, 0x5a, 0x7a, 0x58, 0x04, 0x52, 0x68, 0x25, 0xe8, 0xa1, 0x50, 0xa5, 0x08,
	0x24, 0x2e, 0x91, 0x77, 0xc6, 0x49, 0x5c, 0x4d, 0x6c, 0xd7, 0x76, 0x48, 0x39, 0xc2, 0x47, 0xd8,
	0x7e, 0x0a, 0xce, 0x7c, 0x14, 0x24, 0x0e, 0x9c, 0xf9, 0x20, 0xc8, 0x9e, 0x3f, 0x3b, 0x93, 0x6c,
	0x85, 0xb3, 0x95, 0x10, 0x87, 0x9e, 0xa2, 0x79, 0xfe, 0xfd, 0xfc, 0x9e, 0x9f, 0x7e, 0xbf, 0x67,
	0x07, 0xda, 0x58, 0x88, 0x50, 0x48, 0xae, 0x39, 0x02, 0x2e, 0x08, 0x13, 0x54, 0x4b, 0xfa, 0xbc,
	0xf7, 0xc1, 0x94, 0xf3, 0x69, 0x4a, 0x22, 0xbb, 0x72, 0xb6, 0x98, 0x44, 0x4b, 0x89, 0x85, 0x20,
	0x52, 0x65, 0xd8, 0xde, 0xe1, 0xea, 0xba, 0xa6, 0x73, 0xa2, 0x34, 0x9e, 0xe7, 0x9b, 0xf5, 0x0e,
	0x72, 0x00, 0x16, 0x34, 0xc2, 0x8c, 0x71, 0x8d, 0x35, 0xe5, 0xac, 0xa0, 0x7f, 0x62, 0x7f, 0xe2,
	0x5b, 0x53, 0xc2, 0x6e, 0xa9, 0x25, 0x9e, 0x4e, 0x89, 0x8c, 0xb8, 0xb0, 0x88, 0x75, 0x74, 0xf0,
	0x5b, 0x0b, 0xf6, 0xef, 0x4b, 0x82, 0x35, 0x19, 0x0a, 0x31, 0x22, 0xcf, 0x16, 0x44, 0x69, 0xf4,
	0x11, 0x34, 0xc6, 0x7e, 0xa3, 0xdf, 0x38, 0xea, 0x0c, 0x0e, 0xc2, 0x2c, 0x59, 0x58, 0x54, 0x13,
	0x3e, 0xd1, 0x92, 0xb2, 0xe9, 0xf7, 0x38, 0x5d, 0x90, 0xd1, 0x1b, 0xe8, 0x04, 0xb6, 0x19, 0x9e,
	0x13, 0x7f, 0xcb, 0x01, 0x6d, 0x91, 0xe8, 0x53, 0xd8, 0x91, 0x44, 0xf0, 0x31, 0x4d, 0xfc, 0xa6,
	0x03, 0xc9, 0x33, 0xe0, 0x87, 0x09, 0x1a, 0x40, 0x8b, 0x2f, 0x19, 0x91, 0xfe, 0xb6, 0x03, 0x29,
	0x83, 0xa2, 0xcf, 0x00, 0xe2, 0x19, 0x96, 0x7a, 0x6c, 0x4b, 0x6c, 0x39, 0x10, 0xdb, 0x16, 0xff,
	0x8d, 0xa9, 0xf3, 0x0b, 0xe8, 0x24, 0x44, 0xc5, 0x92, 0xda, 0xde, 0xf9, 0x9e, 0x03, 0xbb, 0x4a,
	0x30, 0x9d, 0x99, 0xf1, 0x39, 0xf1, 0x77, 0x5d, 0x3a, 0x63, 0x90, 0x86, 0x41, 0x63, 0xce, 0xfc,
	0xb6, 0x0b, 0xc3, 0x20, 0x4d, 0x8d, 0x2a, 0x96, 0x84, 0x30, 0x35, 0xe3, 0x5a, 0xf9, 0xe0, 0x52,
	0x63, 0x85, 0x60, 0xf8, 0x73, 0x4c, 0x99, 0xc6, 0x94, 0x11, 0xa9, 0xfc, 0x8e, 0x0b, 0xbf, 0x42,
	0x40, 0xf7, 0x60, 0x47, 0xf1, 0x85, 0x8c, 0x89, 0xf2, 0xf7, 0x1c, 0xb8, 0x05, 0x18, 0xdd, 0x05,
	0x4f, 0x12, 0x9c, 0xcc, 0x89, 0xdf, 0x75, 0x93, 0x80, 0xc1, 0x06, 0xf7, 0xe0, 0xad, 0x8a, 0x54,
	0x95, 0xe0, 0x4c, 0x11, 0xf4, 0x21, 0x34, 0xb1, 0x10, 0xb9, 0x5a, 0xaf, 0x85, 0x17, 0x3e, 0x0b,
	0x0d, 0xca, 0xac, 0x05, 0xbf, 0xb7, 0x60, 0xff, 0x11, 0x4f, 0xe8, 0xe4, 0xe7, 0x8a, 0xc6, 0xef,
	0x80, 0x87, 0x85, 0x30, 0x2a, 0x74, 0x11, 0x7a, 0x0b, 0x0b, 0xf1, 0x30, 0x79, 0xad, 0xf6, 0xd7,
	0x6a, 0xff, 0x2f, 0xd5, 0x5e, 0x11, 0xad, 0xbb, 0xda, 0xbf, 0x82, 0xfd, 0x07, 0x24, 0x25, 0xb5,
	0x81, 0x7e, 0x15, 0xb1, 0x9b, 0x02, 0x2a, 0x1b, 0xb9, 0x17, 0xf0, 0xa7, 0x07, 0xcd, 0xa1, 0x10,
	0xff, 0x73, 0x87, 0xad, 0x08, 0x7e, 0x7b, 0x53, 0xc1, 0xdf, 0x05, 0x4f, 0x69, 0xac, 0x17, 0xca,
	0xc9, 0x69, 0x39, 0xb6, 0xb4, 0x89, 0xb7, 0xb1, 0x4d, 0x76, 0xae, 0x6a, 0x93, 0xdd, 0x57, 0xb4,
	0x49, 0xfb, 0x15, 0x6c, 0x02, 0x57, 0xb3, 0x49, 0xc7, 0xdd, 0x26, 0x2b, 0x53, 0x6f, 0x6f, 0xb3,
	0xa9, 0x57, 0x8e, 0xd9, 0xee, 0x26, 0x63, 0xb6, 0x13, 0xdb, 0x5b, 0x68, 0x6c, 0xde, 0x65, 0xfe,
	0x9b, 0x96, 0xd9, 0x5b, 0x63, 0x7e, 0x57, 0x3c, 0xda, 0x46, 0x90, 0xc1, 0x4d, 0xc0, 0x90, 0x33,
	0x25, 0x64, 0xe4, 0x6b, 0xff, 0x4e, 0xce, 0xe0, 0x26, 0x10, 0xbc, 0xd8, 0x82, 0xb7, 0x1f, 0x58,
	0x09, 0x9e, 0x19, 0x4f, 0xaa, 0xc2, 0xdd, 0x37, 0x2a, 0x46, 0x6b, 0x1e, 0xb5, 0x0b, 0x2b, 0xa1,
	0xd2, 0x4a, 0x26, 0x98, 0x99, 0xe5, 0xdd, 0xaa, 0x59, 0x4c, 0xb8, 0xb0, 0xc3, 0x3b, 0xa5, 0x9c,
	0xb7, 0xb3, 0x78, 0x2e, 0xd8, 0xeb, 0x45, 0x87, 0x5a, 0xd9, 0xd6, 0x59, 0x0f, 0x6e, 0xd6, 0x9a,
	0xee, 0xd9, 0xa5, 0x4a, 0x5b, 0x3f, 0x87, 0x8e, 0x22, 0x58, 0xc6, 0xb3, 0xf1, 0x92, 0xcb, 0xc4,
	0x49, 0xba, 0x90, 0x11, 0x7e, 0xe0, 0x32, 0x31, 0x39, 0x53, 0x3a, 0xa7, 0xda, 0x4a, 0xb7, 0x3b,
	0xca, 0x3e, 0x4c, 0x85, 0x7c, 0x32, 0x51, 0x44, 0x5b, 0x45, 0x76, 0x47, 0xf9, 0x57, 0x80, 0xe1,
	0x7a, 0xbd, 0x29, 0xf9, 0xa4, 0x3a, 0x84, 0x8e, 0xe6, 0x1a, 0xa7, 0xe3, 0x98, 0x2f, 0x98, 0xb6,
	0x33, 0xa8, 0x3b, 0x02, 0x1b, 0xba, 0x6f, 0x22, 0xe8, 0x08, 0x76, 0x4c, 0xdb, 0xcc, 0x8e, 0xa6,
	0x45, 0x97, 0x8c, 0x33, 0xd3, 0xd6, 0x27, 0x44, 0x0f, 0xfe, 0x6a, 0x02, 0x0c, 0x85, 0x78, 0x84,
	0x19, 0x9e, 0x12, 0x89, 0x52, 0x68, 0x97, 0xef, 0x10, 0x74, 0x50, 0x25, 0xad, 0xbe, 0xa4, 0x7b,
	0x37, 0x5f, 0xb2, 0x9a, 0xd5, 0x18, 0x04, 0xe7, 0xc3, 0x3d, 0x94, 0xeb, 0xa3, 0x8f, 0x85, 0xf8,
	0xf5, 0x8f, 0xbf, 0x5f, 0x6c, 0x75, 0x83, 0xdd, 0xe8, 0xa7, 0xdb, 0x11, 0x16, 0x42, 0x9d, 0x36,
	0x8e, 0xd1, 0x2f, 0x0d, 0xd8, 0xab, 0x1e, 0x10, 0x1d, 0x56, 0xf7, 0xbc, 0x44, 0x0f, 0xbd, 0xfe,
	0xcb, 0x01, 0x79, 0xde, 0xf0, 0x7c, 0xf8, 0x3e, 0x7a, 0x2f, 0xc9, 0x97, 0x4c, 0x66, 0xd5, 0x5f,
	0x52, 0x3d, 0xeb, 0x4f, 0x68, 0xaa, 0x89, 0xb4, 0x65, 0x00, 0x2a, 0xcb, 0x30, 0x27, 0x2e, 0xef,
	0xa2, 0xfa, 0x89, 0x57, 0xdf, 0x55, 0xf5, 0x13, 0xaf, 0x5d, 0x60, 0xf9, 0x89, 0xe7, 0x36, 0x7e,
	0x71, 0xe2, 0x41, 0xed, 0xc4, 0x29, 0xb4, 0xcb, 0x8b, 0xa7, 0x9e, 0x6d, 0xf5, 0x62, 0xab, 0x67,
	0x5b, 0xbb, 0xad, 0xf2, 0x6c, 0x89, 0x8d, 0x5f, 0x64, 0x3b, 0xae, 0x66, 0xfb, 0xf2, 0xf9, 0xf9,
	0xf0, 0x19, 0xfa, 0x1a, 0xd0, 0xb7, 0x82, 0xb0, 0xc7, 0x76, 0xa3, 0xfe, 0x63, 0xc9, 0x9f, 0x92,
	0x58, 0x07, 0x1f, 0x5f, 0x16, 0x45, 0x37, 0x66, 0x5a, 0x0b, 0x75, 0x1a, 0x45, 0x95, 0xd4, 0x94,
	0x0f, 0x5a, 0x27, 0xe1, 0x49, 0x78, 0xfb, 0xb8, 0xd1, 0x18, 0xec, 0x63, 0x21, 0x52, 0x1a, 0xdb,
	0xff, 0x59, 0xd1, 0x53, 0xc5, 0xd9, 0xe9, 0x5a, 0xe4, 0xc7, 0x2d, 0x71, 0x76, 0xe6, 0x59, 0x27,
	0xdc, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x28, 0x31, 0xb6, 0x1f, 0x28, 0x0e, 0x00, 0x00,
}
