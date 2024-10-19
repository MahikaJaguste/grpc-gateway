// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: runtime/internal/examplepb/non_standard_names.proto

package examplepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NonStandardService_Update_FullMethodName              = "/grpc.gateway.runtime.internal.examplepb.NonStandardService/Update"
	NonStandardService_UpdateWithJSONNames_FullMethodName = "/grpc.gateway.runtime.internal.examplepb.NonStandardService/UpdateWithJSONNames"
)

// NonStandardServiceClient is the client API for NonStandardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// NonStandardService responds to incoming messages, applies a field mask and returns the masked response.
type NonStandardServiceClient interface {
	// Apply field mask to empty NonStandardMessage and return result.
	Update(ctx context.Context, in *NonStandardUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessage, error)
	// Apply field mask to empty NonStandardMessageWithJSONNames and return result.
	UpdateWithJSONNames(ctx context.Context, in *NonStandardWithJSONNamesUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessageWithJSONNames, error)
}

type nonStandardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNonStandardServiceClient(cc grpc.ClientConnInterface) NonStandardServiceClient {
	return &nonStandardServiceClient{cc}
}

func (c *nonStandardServiceClient) Update(ctx context.Context, in *NonStandardUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NonStandardMessage)
	err := c.cc.Invoke(ctx, NonStandardService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nonStandardServiceClient) UpdateWithJSONNames(ctx context.Context, in *NonStandardWithJSONNamesUpdateRequest, opts ...grpc.CallOption) (*NonStandardMessageWithJSONNames, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NonStandardMessageWithJSONNames)
	err := c.cc.Invoke(ctx, NonStandardService_UpdateWithJSONNames_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NonStandardServiceServer is the server API for NonStandardService service.
// All implementations should embed UnimplementedNonStandardServiceServer
// for forward compatibility.
//
// NonStandardService responds to incoming messages, applies a field mask and returns the masked response.
type NonStandardServiceServer interface {
	// Apply field mask to empty NonStandardMessage and return result.
	Update(context.Context, *NonStandardUpdateRequest) (*NonStandardMessage, error)
	// Apply field mask to empty NonStandardMessageWithJSONNames and return result.
	UpdateWithJSONNames(context.Context, *NonStandardWithJSONNamesUpdateRequest) (*NonStandardMessageWithJSONNames, error)
}

// UnimplementedNonStandardServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNonStandardServiceServer struct{}

func (UnimplementedNonStandardServiceServer) Update(context.Context, *NonStandardUpdateRequest) (*NonStandardMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNonStandardServiceServer) UpdateWithJSONNames(context.Context, *NonStandardWithJSONNamesUpdateRequest) (*NonStandardMessageWithJSONNames, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWithJSONNames not implemented")
}
func (UnimplementedNonStandardServiceServer) testEmbeddedByValue() {}

// UnsafeNonStandardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NonStandardServiceServer will
// result in compilation errors.
type UnsafeNonStandardServiceServer interface {
	mustEmbedUnimplementedNonStandardServiceServer()
}

func RegisterNonStandardServiceServer(s grpc.ServiceRegistrar, srv NonStandardServiceServer) {
	// If the following call pancis, it indicates UnimplementedNonStandardServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NonStandardService_ServiceDesc, srv)
}

func _NonStandardService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonStandardUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonStandardServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NonStandardService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonStandardServiceServer).Update(ctx, req.(*NonStandardUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NonStandardService_UpdateWithJSONNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NonStandardWithJSONNamesUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NonStandardServiceServer).UpdateWithJSONNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NonStandardService_UpdateWithJSONNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NonStandardServiceServer).UpdateWithJSONNames(ctx, req.(*NonStandardWithJSONNamesUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NonStandardService_ServiceDesc is the grpc.ServiceDesc for NonStandardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NonStandardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.runtime.internal.examplepb.NonStandardService",
	HandlerType: (*NonStandardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _NonStandardService_Update_Handler,
		},
		{
			MethodName: "UpdateWithJSONNames",
			Handler:    _NonStandardService_UpdateWithJSONNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "runtime/internal/examplepb/non_standard_names.proto",
}
