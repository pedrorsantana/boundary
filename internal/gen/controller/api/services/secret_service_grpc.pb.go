// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretServiceClient interface {
	GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error)
	ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (*ListSecretsResponse, error)
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error)
	RequestSecretAccess(ctx context.Context, in *RequestSecretAccessRequest, opts ...grpc.CallOption) (*RequestSecretAccessResponse, error)
	ListSecretSessionEvents(ctx context.Context, in *ListSecretSessionEventsRequest, opts ...grpc.CallOption) (*ListSecretSessionEventsResponse, error)
	// Authorize Secret Execution
	AuthorizeSecretSession(ctx context.Context, in *AuthorizeSecretSessionRequest, opts ...grpc.CallOption) (*AuthorizeSecretSessionResponse, error)
	GetSecretSession(ctx context.Context, in *GetSecretSessionRequest, opts ...grpc.CallOption) (*GetSecretSessionResponse, error)
}

type secretServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) GetSecret(ctx context.Context, in *GetSecretRequest, opts ...grpc.CallOption) (*GetSecretResponse, error) {
	out := new(GetSecretResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.SecretService/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (*ListSecretsResponse, error) {
	out := new(ListSecretsResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.SecretService/ListSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error) {
	out := new(CreateSecretResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.SecretService/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error) {
	out := new(DeleteSecretResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.SecretService/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) RequestSecretAccess(ctx context.Context, in *RequestSecretAccessRequest, opts ...grpc.CallOption) (*RequestSecretAccessResponse, error) {
	out := new(RequestSecretAccessResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.SecretService/RequestSecretAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) ListSecretSessionEvents(ctx context.Context, in *ListSecretSessionEventsRequest, opts ...grpc.CallOption) (*ListSecretSessionEventsResponse, error) {
	out := new(ListSecretSessionEventsResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.SecretService/ListSecretSessionEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) AuthorizeSecretSession(ctx context.Context, in *AuthorizeSecretSessionRequest, opts ...grpc.CallOption) (*AuthorizeSecretSessionResponse, error) {
	out := new(AuthorizeSecretSessionResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.SecretService/AuthorizeSecretSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) GetSecretSession(ctx context.Context, in *GetSecretSessionRequest, opts ...grpc.CallOption) (*GetSecretSessionResponse, error) {
	out := new(GetSecretSessionResponse)
	err := c.cc.Invoke(ctx, "/controller.api.services.v1.SecretService/GetSecretSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServiceServer is the server API for SecretService service.
// All implementations must embed UnimplementedSecretServiceServer
// for forward compatibility
type SecretServiceServer interface {
	GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error)
	ListSecrets(context.Context, *ListSecretsRequest) (*ListSecretsResponse, error)
	CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error)
	RequestSecretAccess(context.Context, *RequestSecretAccessRequest) (*RequestSecretAccessResponse, error)
	ListSecretSessionEvents(context.Context, *ListSecretSessionEventsRequest) (*ListSecretSessionEventsResponse, error)
	// Authorize Secret Execution
	AuthorizeSecretSession(context.Context, *AuthorizeSecretSessionRequest) (*AuthorizeSecretSessionResponse, error)
	GetSecretSession(context.Context, *GetSecretSessionRequest) (*GetSecretSessionResponse, error)
	mustEmbedUnimplementedSecretServiceServer()
}

// UnimplementedSecretServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (UnimplementedSecretServiceServer) GetSecret(context.Context, *GetSecretRequest) (*GetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedSecretServiceServer) ListSecrets(context.Context, *ListSecretsRequest) (*ListSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecrets not implemented")
}
func (UnimplementedSecretServiceServer) CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedSecretServiceServer) DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedSecretServiceServer) RequestSecretAccess(context.Context, *RequestSecretAccessRequest) (*RequestSecretAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestSecretAccess not implemented")
}
func (UnimplementedSecretServiceServer) ListSecretSessionEvents(context.Context, *ListSecretSessionEventsRequest) (*ListSecretSessionEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecretSessionEvents not implemented")
}
func (UnimplementedSecretServiceServer) AuthorizeSecretSession(context.Context, *AuthorizeSecretSessionRequest) (*AuthorizeSecretSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeSecretSession not implemented")
}
func (UnimplementedSecretServiceServer) GetSecretSession(context.Context, *GetSecretSessionRequest) (*GetSecretSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecretSession not implemented")
}
func (UnimplementedSecretServiceServer) mustEmbedUnimplementedSecretServiceServer() {}

// UnsafeSecretServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServiceServer will
// result in compilation errors.
type UnsafeSecretServiceServer interface {
	mustEmbedUnimplementedSecretServiceServer()
}

func RegisterSecretServiceServer(s grpc.ServiceRegistrar, srv SecretServiceServer) {
	s.RegisterService(&SecretService_ServiceDesc, srv)
}

func _SecretService_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.SecretService/GetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).GetSecret(ctx, req.(*GetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_ListSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).ListSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.SecretService/ListSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).ListSecrets(ctx, req.(*ListSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.SecretService/CreateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).CreateSecret(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.SecretService/DeleteSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).DeleteSecret(ctx, req.(*DeleteSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_RequestSecretAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSecretAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).RequestSecretAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.SecretService/RequestSecretAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).RequestSecretAccess(ctx, req.(*RequestSecretAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_ListSecretSessionEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecretSessionEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).ListSecretSessionEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.SecretService/ListSecretSessionEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).ListSecretSessionEvents(ctx, req.(*ListSecretSessionEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_AuthorizeSecretSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeSecretSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).AuthorizeSecretSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.SecretService/AuthorizeSecretSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).AuthorizeSecretSession(ctx, req.(*AuthorizeSecretSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_GetSecretSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).GetSecretSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/controller.api.services.v1.SecretService/GetSecretSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).GetSecretSession(ctx, req.(*GetSecretSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretService_ServiceDesc is the grpc.ServiceDesc for SecretService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "controller.api.services.v1.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecret",
			Handler:    _SecretService_GetSecret_Handler,
		},
		{
			MethodName: "ListSecrets",
			Handler:    _SecretService_ListSecrets_Handler,
		},
		{
			MethodName: "CreateSecret",
			Handler:    _SecretService_CreateSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _SecretService_DeleteSecret_Handler,
		},
		{
			MethodName: "RequestSecretAccess",
			Handler:    _SecretService_RequestSecretAccess_Handler,
		},
		{
			MethodName: "ListSecretSessionEvents",
			Handler:    _SecretService_ListSecretSessionEvents_Handler,
		},
		{
			MethodName: "AuthorizeSecretSession",
			Handler:    _SecretService_AuthorizeSecretSession_Handler,
		},
		{
			MethodName: "GetSecretSession",
			Handler:    _SecretService_GetSecretSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller/api/services/v1/secret_service.proto",
}
