// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: spaceone/api/identity/v2/trusted_account.proto

package v2

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TrustedAccount_Create_FullMethodName           = "/spaceone.api.identity.v2.TrustedAccount/create"
	TrustedAccount_Update_FullMethodName           = "/spaceone.api.identity.v2.TrustedAccount/update"
	TrustedAccount_UpdateSecretData_FullMethodName = "/spaceone.api.identity.v2.TrustedAccount/update_secret_data"
	TrustedAccount_Delete_FullMethodName           = "/spaceone.api.identity.v2.TrustedAccount/delete"
	TrustedAccount_Sync_FullMethodName             = "/spaceone.api.identity.v2.TrustedAccount/sync"
	TrustedAccount_Get_FullMethodName              = "/spaceone.api.identity.v2.TrustedAccount/get"
	TrustedAccount_List_FullMethodName             = "/spaceone.api.identity.v2.TrustedAccount/list"
	TrustedAccount_Stat_FullMethodName             = "/spaceone.api.identity.v2.TrustedAccount/stat"
)

// TrustedAccountClient is the client API for TrustedAccount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TrustedAccountClient interface {
	Create(ctx context.Context, in *CreateTrustedAccountRequest, opts ...grpc.CallOption) (*TrustedAccountInfo, error)
	Update(ctx context.Context, in *UpdateTrustedAccountRequest, opts ...grpc.CallOption) (*TrustedAccountInfo, error)
	UpdateSecretData(ctx context.Context, in *UpdateTrustedAccountSecretRequest, opts ...grpc.CallOption) (*TrustedAccountInfo, error)
	Delete(ctx context.Context, in *TrustedAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Sync(ctx context.Context, in *TrustedAccountRequest, opts ...grpc.CallOption) (*JobInfo, error)
	Get(ctx context.Context, in *TrustedAccountRequest, opts ...grpc.CallOption) (*TrustedAccountInfo, error)
	List(ctx context.Context, in *TrustedAccountSearchQuery, opts ...grpc.CallOption) (*TrustedAccountsInfo, error)
	Stat(ctx context.Context, in *TrustedAccountStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error)
}

type trustedAccountClient struct {
	cc grpc.ClientConnInterface
}

func NewTrustedAccountClient(cc grpc.ClientConnInterface) TrustedAccountClient {
	return &trustedAccountClient{cc}
}

func (c *trustedAccountClient) Create(ctx context.Context, in *CreateTrustedAccountRequest, opts ...grpc.CallOption) (*TrustedAccountInfo, error) {
	out := new(TrustedAccountInfo)
	err := c.cc.Invoke(ctx, TrustedAccount_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedAccountClient) Update(ctx context.Context, in *UpdateTrustedAccountRequest, opts ...grpc.CallOption) (*TrustedAccountInfo, error) {
	out := new(TrustedAccountInfo)
	err := c.cc.Invoke(ctx, TrustedAccount_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedAccountClient) UpdateSecretData(ctx context.Context, in *UpdateTrustedAccountSecretRequest, opts ...grpc.CallOption) (*TrustedAccountInfo, error) {
	out := new(TrustedAccountInfo)
	err := c.cc.Invoke(ctx, TrustedAccount_UpdateSecretData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedAccountClient) Delete(ctx context.Context, in *TrustedAccountRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, TrustedAccount_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedAccountClient) Sync(ctx context.Context, in *TrustedAccountRequest, opts ...grpc.CallOption) (*JobInfo, error) {
	out := new(JobInfo)
	err := c.cc.Invoke(ctx, TrustedAccount_Sync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedAccountClient) Get(ctx context.Context, in *TrustedAccountRequest, opts ...grpc.CallOption) (*TrustedAccountInfo, error) {
	out := new(TrustedAccountInfo)
	err := c.cc.Invoke(ctx, TrustedAccount_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedAccountClient) List(ctx context.Context, in *TrustedAccountSearchQuery, opts ...grpc.CallOption) (*TrustedAccountsInfo, error) {
	out := new(TrustedAccountsInfo)
	err := c.cc.Invoke(ctx, TrustedAccount_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trustedAccountClient) Stat(ctx context.Context, in *TrustedAccountStatQuery, opts ...grpc.CallOption) (*_struct.Struct, error) {
	out := new(_struct.Struct)
	err := c.cc.Invoke(ctx, TrustedAccount_Stat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrustedAccountServer is the server API for TrustedAccount service.
// All implementations must embed UnimplementedTrustedAccountServer
// for forward compatibility
type TrustedAccountServer interface {
	Create(context.Context, *CreateTrustedAccountRequest) (*TrustedAccountInfo, error)
	Update(context.Context, *UpdateTrustedAccountRequest) (*TrustedAccountInfo, error)
	UpdateSecretData(context.Context, *UpdateTrustedAccountSecretRequest) (*TrustedAccountInfo, error)
	Delete(context.Context, *TrustedAccountRequest) (*empty.Empty, error)
	Sync(context.Context, *TrustedAccountRequest) (*JobInfo, error)
	Get(context.Context, *TrustedAccountRequest) (*TrustedAccountInfo, error)
	List(context.Context, *TrustedAccountSearchQuery) (*TrustedAccountsInfo, error)
	Stat(context.Context, *TrustedAccountStatQuery) (*_struct.Struct, error)
	mustEmbedUnimplementedTrustedAccountServer()
}

// UnimplementedTrustedAccountServer must be embedded to have forward compatible implementations.
type UnimplementedTrustedAccountServer struct {
}

func (UnimplementedTrustedAccountServer) Create(context.Context, *CreateTrustedAccountRequest) (*TrustedAccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTrustedAccountServer) Update(context.Context, *UpdateTrustedAccountRequest) (*TrustedAccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTrustedAccountServer) UpdateSecretData(context.Context, *UpdateTrustedAccountSecretRequest) (*TrustedAccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSecretData not implemented")
}
func (UnimplementedTrustedAccountServer) Delete(context.Context, *TrustedAccountRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTrustedAccountServer) Sync(context.Context, *TrustedAccountRequest) (*JobInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedTrustedAccountServer) Get(context.Context, *TrustedAccountRequest) (*TrustedAccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTrustedAccountServer) List(context.Context, *TrustedAccountSearchQuery) (*TrustedAccountsInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTrustedAccountServer) Stat(context.Context, *TrustedAccountStatQuery) (*_struct.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (UnimplementedTrustedAccountServer) mustEmbedUnimplementedTrustedAccountServer() {}

// UnsafeTrustedAccountServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TrustedAccountServer will
// result in compilation errors.
type UnsafeTrustedAccountServer interface {
	mustEmbedUnimplementedTrustedAccountServer()
}

func RegisterTrustedAccountServer(s grpc.ServiceRegistrar, srv TrustedAccountServer) {
	s.RegisterService(&TrustedAccount_ServiceDesc, srv)
}

func _TrustedAccount_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTrustedAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedAccountServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedAccount_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedAccountServer).Create(ctx, req.(*CreateTrustedAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedAccount_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrustedAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedAccountServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedAccount_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedAccountServer).Update(ctx, req.(*UpdateTrustedAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedAccount_UpdateSecretData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTrustedAccountSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedAccountServer).UpdateSecretData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedAccount_UpdateSecretData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedAccountServer).UpdateSecretData(ctx, req.(*UpdateTrustedAccountSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedAccount_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustedAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedAccountServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedAccount_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedAccountServer).Delete(ctx, req.(*TrustedAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedAccount_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustedAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedAccountServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedAccount_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedAccountServer).Sync(ctx, req.(*TrustedAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedAccount_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustedAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedAccountServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedAccount_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedAccountServer).Get(ctx, req.(*TrustedAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedAccount_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustedAccountSearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedAccountServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedAccount_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedAccountServer).List(ctx, req.(*TrustedAccountSearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _TrustedAccount_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrustedAccountStatQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrustedAccountServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TrustedAccount_Stat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrustedAccountServer).Stat(ctx, req.(*TrustedAccountStatQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// TrustedAccount_ServiceDesc is the grpc.ServiceDesc for TrustedAccount service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TrustedAccount_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spaceone.api.identity.v2.TrustedAccount",
	HandlerType: (*TrustedAccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "create",
			Handler:    _TrustedAccount_Create_Handler,
		},
		{
			MethodName: "update",
			Handler:    _TrustedAccount_Update_Handler,
		},
		{
			MethodName: "update_secret_data",
			Handler:    _TrustedAccount_UpdateSecretData_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _TrustedAccount_Delete_Handler,
		},
		{
			MethodName: "sync",
			Handler:    _TrustedAccount_Sync_Handler,
		},
		{
			MethodName: "get",
			Handler:    _TrustedAccount_Get_Handler,
		},
		{
			MethodName: "list",
			Handler:    _TrustedAccount_List_Handler,
		},
		{
			MethodName: "stat",
			Handler:    _TrustedAccount_Stat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spaceone/api/identity/v2/trusted_account.proto",
}
