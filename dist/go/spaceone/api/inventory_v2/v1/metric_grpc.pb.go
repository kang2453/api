// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: spaceone/api/inventory_v2/v1/metric.proto

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
	Metric_Create_FullMethodName = "/spaceone.api.inventory_v2.v1.Metric/create"
	Metric_Update_FullMethodName = "/spaceone.api.inventory_v2.v1.Metric/update"
	Metric_Delete_FullMethodName = "/spaceone.api.inventory_v2.v1.Metric/delete"
	Metric_Run_FullMethodName    = "/spaceone.api.inventory_v2.v1.Metric/run"
	Metric_Test_FullMethodName   = "/spaceone.api.inventory_v2.v1.Metric/test"
	Metric_Get_FullMethodName    = "/spaceone.api.inventory_v2.v1.Metric/get"
	Metric_List_FullMethodName   = "/spaceone.api.inventory_v2.v1.Metric/list"
	Metric_Stat_FullMethodName   = "/spaceone.api.inventory_v2.v1.Metric/stat"
)

// MetricClient is the client API for Metric service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricClient interface {
	Create(ctx context.Context, in *CreateMetricRequest, opts ...grpc.CallOption) (*MetricInfo, error)
	Update(ctx context.Context, in *UpdateMetricRequest, opts ...grpc.CallOption) (*MetricInfo, error)
	Delete(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Run(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Test(ctx context.Context, in *MetricTestRequest, opts ...grpc.CallOption) (*_struct.Struct, error)
	Get(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*MetricInfo, error)
	List(ctx context.Context, in *MetricQuery, opts ...grpc.CallOption) (*MetricsInfo, error)
	Stat(ctx context.Context, in *MetricStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type metricClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricClient(cc grpc.ClientConnInterface) MetricClient {
	return &metricClient{cc}
}

func (c *metricClient) Create(ctx context.Context, in *CreateMetricRequest, opts ...grpc.CallOption) (*MetricInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MetricInfo)
	err := c.cc.Invoke(ctx, Metric_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) Update(ctx context.Context, in *UpdateMetricRequest, opts ...grpc.CallOption) (*MetricInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MetricInfo)
	err := c.cc.Invoke(ctx, Metric_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) Delete(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Metric_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) Run(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Metric_Run_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) Test(ctx context.Context, in *MetricTestRequest, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Metric_Test_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) Get(ctx context.Context, in *MetricRequest, opts ...grpc.CallOption) (*MetricInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MetricInfo)
	err := c.cc.Invoke(ctx, Metric_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) List(ctx context.Context, in *MetricQuery, opts ...grpc.CallOption) (*MetricsInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MetricsInfo)
	err := c.cc.Invoke(ctx, Metric_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) Stat(ctx context.Context, in *MetricStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, Metric_Stat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricServer is the server API for Metric service.
// All implementations must embed UnimplementedMetricServer
// for forward compatibility.
type MetricServer interface {
	Create(context.Context, *CreateMetricRequest) (*MetricInfo, error)
	Update(context.Context, *UpdateMetricRequest) (*MetricInfo, error)
	Delete(context.Context, *MetricRequest) (*empty.Empty, error)
	Run(context.Context, *MetricRequest) (*empty.Empty, error)
	Test(context.Context, *MetricTestRequest) (*_struct.Struct, error)
	Get(context.Context, *MetricRequest) (*MetricInfo, error)
	List(context.Context, *MetricQuery) (*MetricsInfo, error)
	Stat(context.Context, *MetricStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedMetricServer()
}

// UnimplementedMetricServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMetricServer struct{}

func (UnimplementedMetricServer) Create(context.Context, *CreateMetricRequest) (*MetricInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMetricServer) Update(context.Context, *UpdateMetricRequest) (*MetricInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedMetricServer) Delete(context.Context, *MetricRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMetricServer) Run(context.Context, *MetricRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedMetricServer) Test(context.Context, *MetricTestRequest) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedMetricServer) Get(context.Context, *MetricRequest) (*MetricInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMetricServer) List(context.Context, *MetricQuery) (*MetricsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedMetricServer) Stat(context.Context, *MetricStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedMetricServer) mustEmbedUnimplementedMetricServer() {}
func (UnimplementedMetricServer) testEmbeddedByValue()                {}

// UnsafeMetricServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricServer will
// result in compilation errors.
type UnsafeMetricServer interface {
	mustEmbedUnimplementedMetricServer()
}

func RegisterMetricServer(s grpc.ServiceRegistrar, srv MetricServer) {
	// If the following call pancis, it indicates UnimplementedMetricServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Metric_ServiceDesc, srv)
}

func _Metric_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).Create(ctx, req.(*CreateMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).Update(ctx, req.(*UpdateMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).Delete(ctx, req.(*MetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_Run_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).Run(ctx, req.(*MetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_Test_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).Test(ctx, req.(*MetricTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).Get(ctx, req.(*MetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).List(ctx, req.(*MetricQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).Stat(ctx, req.(*MetricStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Metric_ServiceDesc is the grpc.ServiceDesc for Metric service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metric_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.inventory_v2.v1.Metric",
	HandlerType: (*MetricServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _Metric_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _Metric_Update_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _Metric_Delete_Handler,
		},
		{
			MethodName: "run",
			Handler:    _Metric_Run_Handler,
		},
		{
			MethodName: "test",
			Handler:    _Metric_Test_Handler,
		},
		{
			MethodName: "get",
			Handler:    _Metric_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Metric_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _Metric_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/inventory_v2/v1/metric.proto",
}
