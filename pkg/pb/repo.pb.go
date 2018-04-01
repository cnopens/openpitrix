// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repo.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/wrappers"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"
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

type CreateRepoRequest struct {
	X           *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=_" json:"_,omitempty"`
	Name        *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description *google_protobuf2.StringValue `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Type        *google_protobuf2.StringValue `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Url         *google_protobuf2.StringValue `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	Credential  *google_protobuf2.StringValue `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Visibility  *google_protobuf2.StringValue `protobuf:"bytes,7,opt,name=visibility" json:"visibility,omitempty"`
	Provider    []string                      `protobuf:"bytes,8,rep,name=provider" json:"provider,omitempty"`
	Label       *google_protobuf2.StringValue `protobuf:"bytes,9,opt,name=label" json:"label,omitempty"`
	Selector    *google_protobuf2.StringValue `protobuf:"bytes,10,opt,name=selector" json:"selector,omitempty"`
}

func (m *CreateRepoRequest) Reset()                    { *m = CreateRepoRequest{} }
func (m *CreateRepoRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRepoRequest) ProtoMessage()               {}
func (*CreateRepoRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *CreateRepoRequest) GetX() *google_protobuf2.StringValue {
	if m != nil {
		return m.X
	}
	return nil
}

func (m *CreateRepoRequest) GetName() *google_protobuf2.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateRepoRequest) GetDescription() *google_protobuf2.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CreateRepoRequest) GetType() *google_protobuf2.StringValue {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *CreateRepoRequest) GetUrl() *google_protobuf2.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *CreateRepoRequest) GetCredential() *google_protobuf2.StringValue {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *CreateRepoRequest) GetVisibility() *google_protobuf2.StringValue {
	if m != nil {
		return m.Visibility
	}
	return nil
}

func (m *CreateRepoRequest) GetProvider() []string {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *CreateRepoRequest) GetLabel() *google_protobuf2.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *CreateRepoRequest) GetSelector() *google_protobuf2.StringValue {
	if m != nil {
		return m.Selector
	}
	return nil
}

type CreateRepoResponse struct {
	Repo *Repo `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
}

func (m *CreateRepoResponse) Reset()                    { *m = CreateRepoResponse{} }
func (m *CreateRepoResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateRepoResponse) ProtoMessage()               {}
func (*CreateRepoResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *CreateRepoResponse) GetRepo() *Repo {
	if m != nil {
		return m.Repo
	}
	return nil
}

type ModifyRepoRequest struct {
	RepoId      *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Name        *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description *google_protobuf2.StringValue `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Type        *google_protobuf2.StringValue `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Url         *google_protobuf2.StringValue `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	Credential  *google_protobuf2.StringValue `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Visibility  *google_protobuf2.StringValue `protobuf:"bytes,7,opt,name=visibility" json:"visibility,omitempty"`
	Label       *google_protobuf2.StringValue `protobuf:"bytes,8,opt,name=label" json:"label,omitempty"`
	Selector    *google_protobuf2.StringValue `protobuf:"bytes,9,opt,name=selector" json:"selector,omitempty"`
}

func (m *ModifyRepoRequest) Reset()                    { *m = ModifyRepoRequest{} }
func (m *ModifyRepoRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyRepoRequest) ProtoMessage()               {}
func (*ModifyRepoRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *ModifyRepoRequest) GetRepoId() *google_protobuf2.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *ModifyRepoRequest) GetName() *google_protobuf2.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ModifyRepoRequest) GetDescription() *google_protobuf2.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *ModifyRepoRequest) GetType() *google_protobuf2.StringValue {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ModifyRepoRequest) GetUrl() *google_protobuf2.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *ModifyRepoRequest) GetCredential() *google_protobuf2.StringValue {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *ModifyRepoRequest) GetVisibility() *google_protobuf2.StringValue {
	if m != nil {
		return m.Visibility
	}
	return nil
}

func (m *ModifyRepoRequest) GetLabel() *google_protobuf2.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *ModifyRepoRequest) GetSelector() *google_protobuf2.StringValue {
	if m != nil {
		return m.Selector
	}
	return nil
}

type ModifyRepoResponse struct {
	Repo *Repo `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
}

func (m *ModifyRepoResponse) Reset()                    { *m = ModifyRepoResponse{} }
func (m *ModifyRepoResponse) String() string            { return proto.CompactTextString(m) }
func (*ModifyRepoResponse) ProtoMessage()               {}
func (*ModifyRepoResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *ModifyRepoResponse) GetRepo() *Repo {
	if m != nil {
		return m.Repo
	}
	return nil
}

type DeleteRepoRequest struct {
	RepoId *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
}

func (m *DeleteRepoRequest) Reset()                    { *m = DeleteRepoRequest{} }
func (m *DeleteRepoRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRepoRequest) ProtoMessage()               {}
func (*DeleteRepoRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *DeleteRepoRequest) GetRepoId() *google_protobuf2.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

type DeleteRepoResponse struct {
	Repo *Repo `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
}

func (m *DeleteRepoResponse) Reset()                    { *m = DeleteRepoResponse{} }
func (m *DeleteRepoResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteRepoResponse) ProtoMessage()               {}
func (*DeleteRepoResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *DeleteRepoResponse) GetRepo() *Repo {
	if m != nil {
		return m.Repo
	}
	return nil
}

type RepoLabel struct {
	RepoLabelId *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=repo_label_id,json=repoLabelId" json:"repo_label_id,omitempty"`
	RepoId      *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	LabelKey    *google_protobuf2.StringValue `protobuf:"bytes,3,opt,name=label_key,json=labelKey" json:"label_key,omitempty"`
	LabelValue  *google_protobuf2.StringValue `protobuf:"bytes,4,opt,name=label_value,json=labelValue" json:"label_value,omitempty"`
	CreateTime  *google_protobuf3.Timestamp   `protobuf:"bytes,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
}

func (m *RepoLabel) Reset()                    { *m = RepoLabel{} }
func (m *RepoLabel) String() string            { return proto.CompactTextString(m) }
func (*RepoLabel) ProtoMessage()               {}
func (*RepoLabel) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *RepoLabel) GetRepoLabelId() *google_protobuf2.StringValue {
	if m != nil {
		return m.RepoLabelId
	}
	return nil
}

func (m *RepoLabel) GetRepoId() *google_protobuf2.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *RepoLabel) GetLabelKey() *google_protobuf2.StringValue {
	if m != nil {
		return m.LabelKey
	}
	return nil
}

func (m *RepoLabel) GetLabelValue() *google_protobuf2.StringValue {
	if m != nil {
		return m.LabelValue
	}
	return nil
}

func (m *RepoLabel) GetCreateTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type RepoSelector struct {
	RepoSelectorId *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=repo_selector_id,json=repoSelectorId" json:"repo_selector_id,omitempty"`
	RepoId         *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	SelectorKey    *google_protobuf2.StringValue `protobuf:"bytes,3,opt,name=selector_key,json=selectorKey" json:"selector_key,omitempty"`
	SelectorValue  *google_protobuf2.StringValue `protobuf:"bytes,4,opt,name=selector_value,json=selectorValue" json:"selector_value,omitempty"`
	CreateTime     *google_protobuf3.Timestamp   `protobuf:"bytes,5,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
}

func (m *RepoSelector) Reset()                    { *m = RepoSelector{} }
func (m *RepoSelector) String() string            { return proto.CompactTextString(m) }
func (*RepoSelector) ProtoMessage()               {}
func (*RepoSelector) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *RepoSelector) GetRepoSelectorId() *google_protobuf2.StringValue {
	if m != nil {
		return m.RepoSelectorId
	}
	return nil
}

func (m *RepoSelector) GetRepoId() *google_protobuf2.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *RepoSelector) GetSelectorKey() *google_protobuf2.StringValue {
	if m != nil {
		return m.SelectorKey
	}
	return nil
}

func (m *RepoSelector) GetSelectorValue() *google_protobuf2.StringValue {
	if m != nil {
		return m.SelectorValue
	}
	return nil
}

func (m *RepoSelector) GetCreateTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

type Repo struct {
	RepoId      *google_protobuf2.StringValue `protobuf:"bytes,1,opt,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Name        *google_protobuf2.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description *google_protobuf2.StringValue `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Type        *google_protobuf2.StringValue `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Url         *google_protobuf2.StringValue `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	Credential  *google_protobuf2.StringValue `protobuf:"bytes,6,opt,name=credential" json:"credential,omitempty"`
	Visibility  *google_protobuf2.StringValue `protobuf:"bytes,7,opt,name=visibility" json:"visibility,omitempty"`
	Owner       *google_protobuf2.StringValue `protobuf:"bytes,8,opt,name=owner" json:"owner,omitempty"`
	Providers   []string                      `protobuf:"bytes,9,rep,name=providers" json:"providers,omitempty"`
	Labels      []*RepoLabel                  `protobuf:"bytes,10,rep,name=labels" json:"labels,omitempty"`
	Selectors   []*RepoSelector               `protobuf:"bytes,11,rep,name=selectors" json:"selectors,omitempty"`
	Status      *google_protobuf2.StringValue `protobuf:"bytes,12,opt,name=status" json:"status,omitempty"`
	CreateTime  *google_protobuf3.Timestamp   `protobuf:"bytes,13,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	StatusTime  *google_protobuf3.Timestamp   `protobuf:"bytes,14,opt,name=status_time,json=statusTime" json:"status_time,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *Repo) GetRepoId() *google_protobuf2.StringValue {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *Repo) GetName() *google_protobuf2.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Repo) GetDescription() *google_protobuf2.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *Repo) GetType() *google_protobuf2.StringValue {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Repo) GetUrl() *google_protobuf2.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *Repo) GetCredential() *google_protobuf2.StringValue {
	if m != nil {
		return m.Credential
	}
	return nil
}

func (m *Repo) GetVisibility() *google_protobuf2.StringValue {
	if m != nil {
		return m.Visibility
	}
	return nil
}

func (m *Repo) GetOwner() *google_protobuf2.StringValue {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Repo) GetProviders() []string {
	if m != nil {
		return m.Providers
	}
	return nil
}

func (m *Repo) GetLabels() []*RepoLabel {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Repo) GetSelectors() []*RepoSelector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *Repo) GetStatus() *google_protobuf2.StringValue {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Repo) GetCreateTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Repo) GetStatusTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.StatusTime
	}
	return nil
}

type DescribeReposRequest struct {
	RepoId     []string                      `protobuf:"bytes,1,rep,name=repo_id,json=repoId" json:"repo_id,omitempty"`
	Name       []string                      `protobuf:"bytes,2,rep,name=name" json:"name,omitempty"`
	Type       []string                      `protobuf:"bytes,3,rep,name=type" json:"type,omitempty"`
	Visibility []string                      `protobuf:"bytes,4,rep,name=visibility" json:"visibility,omitempty"`
	Status     []string                      `protobuf:"bytes,5,rep,name=status" json:"status,omitempty"`
	Provider   []string                      `protobuf:"bytes,6,rep,name=provider" json:"provider,omitempty"`
	Label      *google_protobuf2.StringValue `protobuf:"bytes,7,opt,name=label" json:"label,omitempty"`
	Selector   *google_protobuf2.StringValue `protobuf:"bytes,8,opt,name=selector" json:"selector,omitempty"`
	Limit      uint32                        `protobuf:"varint,9,opt,name=limit" json:"limit,omitempty"`
	Offset     uint32                        `protobuf:"varint,10,opt,name=offset" json:"offset,omitempty"`
}

func (m *DescribeReposRequest) Reset()                    { *m = DescribeReposRequest{} }
func (m *DescribeReposRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeReposRequest) ProtoMessage()               {}
func (*DescribeReposRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

func (m *DescribeReposRequest) GetRepoId() []string {
	if m != nil {
		return m.RepoId
	}
	return nil
}

func (m *DescribeReposRequest) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *DescribeReposRequest) GetType() []string {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DescribeReposRequest) GetVisibility() []string {
	if m != nil {
		return m.Visibility
	}
	return nil
}

func (m *DescribeReposRequest) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeReposRequest) GetProvider() []string {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *DescribeReposRequest) GetLabel() *google_protobuf2.StringValue {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *DescribeReposRequest) GetSelector() *google_protobuf2.StringValue {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *DescribeReposRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescribeReposRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type DescribeReposResponse struct {
	TotalCount uint32  `protobuf:"varint,1,opt,name=total_count,json=totalCount" json:"total_count,omitempty"`
	RepoSet    []*Repo `protobuf:"bytes,2,rep,name=repo_set,json=repoSet" json:"repo_set,omitempty"`
}

func (m *DescribeReposResponse) Reset()                    { *m = DescribeReposResponse{} }
func (m *DescribeReposResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeReposResponse) ProtoMessage()               {}
func (*DescribeReposResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *DescribeReposResponse) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *DescribeReposResponse) GetRepoSet() []*Repo {
	if m != nil {
		return m.RepoSet
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRepoRequest)(nil), "openpitrix.CreateRepoRequest")
	proto.RegisterType((*CreateRepoResponse)(nil), "openpitrix.CreateRepoResponse")
	proto.RegisterType((*ModifyRepoRequest)(nil), "openpitrix.ModifyRepoRequest")
	proto.RegisterType((*ModifyRepoResponse)(nil), "openpitrix.ModifyRepoResponse")
	proto.RegisterType((*DeleteRepoRequest)(nil), "openpitrix.DeleteRepoRequest")
	proto.RegisterType((*DeleteRepoResponse)(nil), "openpitrix.DeleteRepoResponse")
	proto.RegisterType((*RepoLabel)(nil), "openpitrix.RepoLabel")
	proto.RegisterType((*RepoSelector)(nil), "openpitrix.RepoSelector")
	proto.RegisterType((*Repo)(nil), "openpitrix.Repo")
	proto.RegisterType((*DescribeReposRequest)(nil), "openpitrix.DescribeReposRequest")
	proto.RegisterType((*DescribeReposResponse)(nil), "openpitrix.DescribeReposResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RepoManager service

type RepoManagerClient interface {
	CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error)
	DescribeRepos(ctx context.Context, in *DescribeReposRequest, opts ...grpc.CallOption) (*DescribeReposResponse, error)
	ModifyRepo(ctx context.Context, in *ModifyRepoRequest, opts ...grpc.CallOption) (*ModifyRepoResponse, error)
	DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*DeleteRepoResponse, error)
}

type repoManagerClient struct {
	cc *grpc.ClientConn
}

func NewRepoManagerClient(cc *grpc.ClientConn) RepoManagerClient {
	return &repoManagerClient{cc}
}

func (c *repoManagerClient) CreateRepo(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error) {
	out := new(CreateRepoResponse)
	err := grpc.Invoke(ctx, "/openpitrix.RepoManager/CreateRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoManagerClient) DescribeRepos(ctx context.Context, in *DescribeReposRequest, opts ...grpc.CallOption) (*DescribeReposResponse, error) {
	out := new(DescribeReposResponse)
	err := grpc.Invoke(ctx, "/openpitrix.RepoManager/DescribeRepos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoManagerClient) ModifyRepo(ctx context.Context, in *ModifyRepoRequest, opts ...grpc.CallOption) (*ModifyRepoResponse, error) {
	out := new(ModifyRepoResponse)
	err := grpc.Invoke(ctx, "/openpitrix.RepoManager/ModifyRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoManagerClient) DeleteRepo(ctx context.Context, in *DeleteRepoRequest, opts ...grpc.CallOption) (*DeleteRepoResponse, error) {
	out := new(DeleteRepoResponse)
	err := grpc.Invoke(ctx, "/openpitrix.RepoManager/DeleteRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RepoManager service

type RepoManagerServer interface {
	CreateRepo(context.Context, *CreateRepoRequest) (*CreateRepoResponse, error)
	DescribeRepos(context.Context, *DescribeReposRequest) (*DescribeReposResponse, error)
	ModifyRepo(context.Context, *ModifyRepoRequest) (*ModifyRepoResponse, error)
	DeleteRepo(context.Context, *DeleteRepoRequest) (*DeleteRepoResponse, error)
}

func RegisterRepoManagerServer(s *grpc.Server, srv RepoManagerServer) {
	s.RegisterService(&_RepoManager_serviceDesc, srv)
}

func _RepoManager_CreateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoManagerServer).CreateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoManager/CreateRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoManagerServer).CreateRepo(ctx, req.(*CreateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoManager_DescribeRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeReposRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoManagerServer).DescribeRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoManager/DescribeRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoManagerServer).DescribeRepos(ctx, req.(*DescribeReposRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoManager_ModifyRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoManagerServer).ModifyRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoManager/ModifyRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoManagerServer).ModifyRepo(ctx, req.(*ModifyRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RepoManager_DeleteRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoManagerServer).DeleteRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openpitrix.RepoManager/DeleteRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoManagerServer).DeleteRepo(ctx, req.(*DeleteRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepoManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openpitrix.RepoManager",
	HandlerType: (*RepoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepo",
			Handler:    _RepoManager_CreateRepo_Handler,
		},
		{
			MethodName: "DescribeRepos",
			Handler:    _RepoManager_DescribeRepos_Handler,
		},
		{
			MethodName: "ModifyRepo",
			Handler:    _RepoManager_ModifyRepo_Handler,
		},
		{
			MethodName: "DeleteRepo",
			Handler:    _RepoManager_DeleteRepo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repo.proto",
}

func init() { proto.RegisterFile("repo.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 954 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x4b, 0x73, 0x1b, 0x45,
	0x10, 0x46, 0x4f, 0x6b, 0x7b, 0x2d, 0x57, 0x3c, 0x95, 0xc0, 0x96, 0xca, 0x38, 0x42, 0x95, 0x83,
	0x09, 0x58, 0x0a, 0xe2, 0x51, 0x10, 0x9e, 0xc1, 0x29, 0xaa, 0x02, 0xe4, 0xb2, 0xa1, 0x38, 0x70,
	0x71, 0xad, 0xb4, 0x2d, 0x31, 0xc5, 0x6a, 0x67, 0x3d, 0x33, 0xb2, 0xd1, 0x95, 0x03, 0xc5, 0x39,
	0x5c, 0xf8, 0x2f, 0xfc, 0x00, 0xae, 0x54, 0xf1, 0x17, 0xb8, 0x73, 0xe3, 0x4c, 0x4d, 0xef, 0xae,
	0xb4, 0x2b, 0x25, 0x30, 0x52, 0xe5, 0xc0, 0x21, 0x27, 0x7b, 0xa6, 0xbf, 0x6f, 0xfa, 0xe1, 0x6f,
	0xbb, 0xdb, 0x00, 0x12, 0x13, 0xd1, 0x4f, 0xa4, 0xd0, 0x82, 0x81, 0x48, 0x30, 0x4e, 0xb8, 0x96,
	0xfc, 0xfb, 0xce, 0xf1, 0x54, 0x88, 0x69, 0x84, 0x03, 0xb2, 0x8c, 0xe6, 0x93, 0xc1, 0x95, 0x0c,
	0x92, 0x04, 0xa5, 0x4a, 0xb1, 0x9d, 0x9b, 0xeb, 0x76, 0xcd, 0x67, 0xa8, 0x74, 0x30, 0x4b, 0x32,
	0xc0, 0x51, 0x06, 0x08, 0x12, 0x3e, 0x08, 0xe2, 0x58, 0xe8, 0x40, 0x73, 0x11, 0xe7, 0xf4, 0xd7,
	0xe9, 0xc7, 0xf8, 0x74, 0x8a, 0xf1, 0xa9, 0xba, 0x0a, 0xa6, 0x53, 0x94, 0x03, 0x91, 0x10, 0x62,
	0x13, 0xdd, 0xfb, 0xa5, 0x0e, 0x87, 0x67, 0x12, 0x03, 0x8d, 0x3e, 0x26, 0xc2, 0xc7, 0x8b, 0x39,
	0x2a, 0xcd, 0x5e, 0x85, 0xca, 0xb9, 0x57, 0xe9, 0x56, 0x4e, 0xdc, 0xe1, 0x51, 0x3f, 0xf5, 0xd6,
	0xcf, 0xc3, 0xe9, 0x3f, 0xd2, 0x92, 0xc7, 0xd3, 0xaf, 0x83, 0x68, 0x8e, 0xfe, 0x0b, 0xec, 0x0e,
	0xd4, 0xe3, 0x60, 0x86, 0x5e, 0xd5, 0x02, 0x4d, 0x48, 0xf6, 0x11, 0xb8, 0x21, 0xaa, 0xb1, 0xe4,
	0x14, 0x94, 0x57, 0xb3, 0x20, 0x16, 0x09, 0xc6, 0xa3, 0x5e, 0x24, 0xe8, 0xd5, 0x6d, 0x3c, 0x1a,
	0x24, 0xeb, 0x43, 0x6d, 0x2e, 0x23, 0xaf, 0x61, 0x41, 0x30, 0x40, 0xf6, 0x01, 0xc0, 0x58, 0x62,
	0x88, 0xb1, 0xe6, 0x41, 0xe4, 0x35, 0x2d, 0x68, 0x05, 0xbc, 0x61, 0x5f, 0x72, 0xc5, 0x47, 0x3c,
	0xe2, 0x7a, 0xe1, 0xed, 0xd9, 0xb0, 0x57, 0x78, 0xd6, 0x81, 0x56, 0x22, 0xc5, 0x25, 0x0f, 0x51,
	0x7a, 0xad, 0x6e, 0xed, 0xc4, 0xf1, 0x97, 0x67, 0x36, 0x84, 0x46, 0x14, 0x8c, 0x30, 0xf2, 0x1c,
	0x8b, 0x47, 0x53, 0x28, 0x7b, 0x17, 0x5a, 0x0a, 0x23, 0x1c, 0x6b, 0x21, 0x3d, 0xb0, 0xa0, 0x2d,
	0xd1, 0xbd, 0xbb, 0xc0, 0x8a, 0xca, 0x50, 0x89, 0x88, 0x15, 0xb2, 0x5b, 0x50, 0x37, 0xba, 0xce,
	0xd4, 0x71, 0xad, 0xbf, 0x12, 0x76, 0x9f, 0x70, 0x64, 0xed, 0xfd, 0x54, 0x87, 0xc3, 0x87, 0x22,
	0xe4, 0x93, 0x45, 0x51, 0x56, 0x6f, 0xc3, 0x9e, 0xb1, 0x9e, 0xf3, 0xd0, 0x4a, 0x5c, 0x4d, 0x03,
	0x7e, 0x10, 0x3e, 0x97, 0xd8, 0x33, 0x96, 0xd8, 0x52, 0x46, 0xad, 0xdd, 0x64, 0xe4, 0x6c, 0x2b,
	0xa3, 0xa2, 0x12, 0xb6, 0x92, 0xd1, 0xe7, 0x70, 0x78, 0x1f, 0x23, 0x2c, 0x37, 0xa7, 0xdd, 0x54,
	0x64, 0xe2, 0x28, 0xbe, 0xb5, 0x55, 0x1c, 0xbf, 0x56, 0xc1, 0x31, 0xc7, 0x2f, 0xa9, 0x16, 0x9f,
	0x40, 0x9b, 0x02, 0xa0, 0xca, 0xd8, 0x86, 0xe1, 0xca, 0x9c, 0xff, 0x20, 0x2c, 0xa6, 0x50, 0xdd,
	0xe2, 0x43, 0x78, 0x0f, 0x9c, 0xd4, 0xe7, 0x77, 0xb8, 0xb0, 0x12, 0x75, 0x8b, 0xe0, 0x5f, 0xe0,
	0x82, 0x7d, 0x08, 0x6e, 0x4a, 0xbd, 0x34, 0x06, 0x2b, 0x61, 0x03, 0x11, 0xe8, 0x77, 0xf6, 0x3e,
	0xb8, 0x63, 0xea, 0x05, 0xe7, 0x66, 0x18, 0x65, 0x32, 0xef, 0x6c, 0xd0, 0xbf, 0xca, 0x27, 0x15,
	0xa9, 0x35, 0xd0, 0x68, 0x2e, 0x7a, 0xbf, 0x57, 0x61, 0xdf, 0x54, 0xef, 0x51, 0x26, 0x09, 0xf6,
	0x19, 0x5c, 0xa3, 0xf4, 0x73, 0x8d, 0xd8, 0xd6, 0xf0, 0x40, 0x16, 0x5e, 0xd9, 0xbd, 0x8c, 0x1f,
	0xc3, 0xfe, 0xd2, 0xb3, 0x6d, 0x25, 0xdd, 0x9c, 0x61, 0x8a, 0x79, 0x06, 0x07, 0xcb, 0x07, 0xec,
	0xeb, 0xd9, 0xce, 0x39, 0xcf, 0xa0, 0xa4, 0x7f, 0x37, 0xa0, 0x6e, 0x4a, 0xfa, 0xbc, 0xa5, 0xfe,
	0x6f, 0x5a, 0xaa, 0xb8, 0x8a, 0x69, 0x64, 0x5b, 0xb4, 0x54, 0x82, 0xb2, 0x23, 0x70, 0xf2, 0xc9,
	0xae, 0x3c, 0x87, 0x46, 0xfd, 0xea, 0x82, 0x9d, 0x42, 0x93, 0xbe, 0x3f, 0xe5, 0x41, 0xb7, 0x76,
	0xe2, 0x0e, 0x6f, 0xac, 0xb7, 0x26, 0xea, 0x25, 0x7e, 0x06, 0x62, 0xef, 0x80, 0x93, 0xcb, 0x4b,
	0x79, 0x2e, 0x31, 0xbc, 0x75, 0x46, 0xfe, 0xe5, 0xf8, 0x2b, 0x28, 0x7b, 0x0b, 0x9a, 0x4a, 0x07,
	0x7a, 0xae, 0xbc, 0x7d, 0x1b, 0xf9, 0xa4, 0xd8, 0x75, 0xed, 0xb6, 0xb7, 0xd1, 0xae, 0x21, 0xa7,
	0xcf, 0xa4, 0xe4, 0x83, 0xff, 0x26, 0xa7, 0x70, 0x12, 0xfe, 0x6f, 0x55, 0xb8, 0x7e, 0x9f, 0x64,
	0x35, 0xa2, 0x46, 0xae, 0xf2, 0xa9, 0xf0, 0x52, 0xf1, 0x43, 0x30, 0xb5, 0xcc, 0xa5, 0xce, 0x96,
	0x52, 0x37, 0xb7, 0xa9, 0x98, 0x59, 0x26, 0xc6, 0x5a, 0x7a, 0x47, 0x72, 0x3b, 0x2e, 0x09, 0xa0,
	0x4e, 0x96, 0xe2, 0x9f, 0xf8, 0xc5, 0x65, 0xa5, 0x1a, 0xe9, 0xfb, 0x59, 0x2d, 0x8a, 0x0b, 0x5b,
	0xf3, 0x69, 0x0b, 0xdb, 0xde, 0x6e, 0x93, 0xb6, 0xb5, 0xcd, 0xa4, 0x65, 0xd7, 0xa1, 0x11, 0xf1,
	0x19, 0xd7, 0x34, 0xa0, 0xdb, 0x7e, 0x7a, 0x30, 0x71, 0x8b, 0xc9, 0x44, 0xa1, 0xa6, 0xf5, 0xaf,
	0xed, 0x67, 0xa7, 0x1e, 0xc2, 0x8d, 0xb5, 0x42, 0x66, 0x23, 0xf1, 0x26, 0xb8, 0x5a, 0xe8, 0x20,
	0x3a, 0x1f, 0x8b, 0x79, 0xac, 0xa9, 0xad, 0xb4, 0x7d, 0xa0, 0xab, 0x33, 0x73, 0xc3, 0x5e, 0x83,
	0x56, 0xd6, 0xbe, 0x35, 0x55, 0xf5, 0x49, 0x73, 0x73, 0x2f, 0x6d, 0xd5, 0x7a, 0xf8, 0x57, 0x0d,
	0x5c, 0x73, 0xf3, 0x30, 0x88, 0x83, 0x29, 0x4a, 0x76, 0x01, 0xb0, 0xda, 0x2a, 0xd9, 0xcb, 0x45,
	0xe2, 0xc6, 0xff, 0x21, 0x9d, 0xe3, 0xa7, 0x99, 0xd3, 0x50, 0x7b, 0xb7, 0x1e, 0xdf, 0x6b, 0xb3,
	0x4c, 0x8a, 0x5d, 0xe3, 0xf1, 0x87, 0x3f, 0xfe, 0xfc, 0xb9, 0x7a, 0xd0, 0x73, 0x06, 0x97, 0x6f,
	0x0c, 0xcc, 0x59, 0xdd, 0xad, 0xdc, 0x66, 0x3f, 0x56, 0xa0, 0x5d, 0x4a, 0x95, 0x75, 0x8b, 0xef,
	0x3e, 0x49, 0x4e, 0x9d, 0x57, 0xfe, 0x05, 0x91, 0x39, 0xbf, 0xf3, 0xf8, 0xde, 0x11, 0xeb, 0x84,
	0x99, 0x8d, 0xdc, 0xab, 0xee, 0x15, 0xd7, 0xdf, 0x76, 0x27, 0x3c, 0xd2, 0x28, 0x29, 0x16, 0x97,
	0xad, 0x62, 0x31, 0xb9, 0xaf, 0x56, 0xa1, 0x72, 0xee, 0x1b, 0xcb, 0x72, 0x39, 0xf7, 0xcd, 0x0d,
	0x2a, 0xcb, 0x7d, 0x46, 0x86, 0x42, 0xee, 0xc3, 0x72, 0xee, 0x17, 0x00, 0xab, 0xad, 0xa7, 0xec,
	0x72, 0x63, 0xb3, 0x2a, 0xbb, 0xdc, 0x5c, 0x96, 0x32, 0x97, 0x21, 0x19, 0x0a, 0x2e, 0x6f, 0x97,
	0x5c, 0x7e, 0x5a, 0xff, 0xa6, 0x9a, 0x8c, 0x46, 0x4d, 0x12, 0xeb, 0x9b, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x5d, 0x23, 0xf7, 0x06, 0x0f, 0x00, 0x00,
}
