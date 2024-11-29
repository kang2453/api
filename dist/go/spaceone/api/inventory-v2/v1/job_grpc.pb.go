// A Job is an act of collecting external cloud resources through plugins.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/inventory_v2/v1/job.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Job_Delete_FullMethodName  = "/spaceone.api.inventory_v2.v1.Job/delete"
	Job_Get_FullMethodName     = "/spaceone.api.inventory_v2.v1.Job/get"
	Job_List_FullMethodName    = "/spaceone.api.inventory_v2.v1.Job/list"
	Job_Analyze_FullMethodName = "/spaceone.api.inventory_v2.v1.Job/analyze"
	Job_Stat_FullMethodName    = "/spaceone.api.inventory_v2.v1.Job/stat"
)

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobClient interface {
	// Deletes a specific Job. You must specify the `job_id` of the Job to delete.
	Delete(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Gets a specific Job. Prints detailed information about the Job, including its state, total tasks, and collector info.
	Get(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobInfo, error)
	// Gets a list of all Jobs. You can use a query to get a filtered list of Jobs.
	List(ctx context.Context, in *JobsQuery, opts ...grpc.CallOption) (*JobsInfo, error)
	Analyze(ctx context.Context, in *JobAnalyzeQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
	Stat(ctx context.Context, in *JobStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) Delete(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Job_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Get(ctx context.Context, in *JobRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobInfo)
	err := c.cc.Invoke(ctx, Job_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) List(ctx context.Context, in *JobsQuery, opts ...grpc.CallOption) (*JobsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JobsInfo)
	err := c.cc.Invoke(ctx, Job_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Analyze(ctx context.Context, in *JobAnalyzeQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Job_Analyze_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) Stat(ctx context.Context, in *JobStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Job_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
// All implementations must embed UnimplementedJobServer
// for forward compatibility.
type JobServer interface {
	// Deletes a specific Job. You must specify the `job_id` of the Job to delete.
	Delete(context.Context, *JobRequest) (*empty.Empty, error)
	// Gets a specific Job. Prints detailed information about the Job, including its state, total tasks, and collector info.
	Get(context.Context, *JobRequest) (*JobInfo, error)
	// Gets a list of all Jobs. You can use a query to get a filtered list of Jobs.
	List(context.Context, *JobsQuery) (*JobsInfo, error)
	Analyze(context.Context, *JobAnalyzeQuery) (*_struct.Struct, error)
	Stat(context.Context, *JobStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedJobServer()
}

// UnimplementedJobServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJobServer struct{}

func (UnimplementedJobServer) Delete(context.Context, *JobRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedJobServer) Get(context.Context, *JobRequest) (*JobInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedJobServer) List(context.Context, *JobsQuery) (*JobsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedJobServer) Analyze(context.Context, *JobAnalyzeQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Analyze not implemented")
}
func (UnimplementedJobServer) Stat(context.Context, *JobStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedJobServer) mustEmbedUnimplementedJobServer() {}
func (UnimplementedJobServer) testEmbeddedByValue()             {}

// UnsafeJobServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServer will
// result in compilation errors.
type UnsafeJobServer interface {
	mustEmbedUnimplementedJobServer()
}

func RegisterJobServer(s grpc.ServiceRegistrar, srv JobServer) {
	// If the following call pancis, it indicates UnimplementedJobServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Job_ServiceDesc, srv)
}

func _Job_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Delete(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Get(ctx, req.(*JobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobsQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).List(ctx, req.(*JobsQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_Analyze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobAnalyzeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Analyze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_Analyze_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Analyze(ctx, req.(*JobAnalyzeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Job_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).Stat(ctx, req.(*JobStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Job_ServiceDesc is the grpc.ServiceDesc for Job service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Job_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.inventory_v2.v1.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "delete",
			Handler:    _Job_Delete_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Job_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Job_List_Handler,
		},
		{
			MethodName: "analyze",
			Handler:    _Job_Analyze_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Job_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/inventory_v2/v1/job.proto",
}
