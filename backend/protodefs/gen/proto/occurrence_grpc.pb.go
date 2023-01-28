// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// OccurreceServiceClient is the client API for OccurreceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OccurreceServiceClient interface {
	Create(ctx context.Context, in *CreateOccurrence, opts ...grpc.CallOption) (*StatusResponse, error)
	GetById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*OccurrenceDetails, error)
	EditById(ctx context.Context, in *EditOccurrence, opts ...grpc.CallOption) (*OccurrenceDetails, error)
	SoftDeleteById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*StatusResponse, error)
	AcceptById(ctx context.Context, in *UpdateOccurrenceStatus, opts ...grpc.CallOption) (*StatusResponse, error)
	RefuseById(ctx context.Context, in *UpdateOccurrenceStatus, opts ...grpc.CallOption) (*StatusResponse, error)
	CancelById(ctx context.Context, in *UpdateOccurrenceStatus, opts ...grpc.CallOption) (*StatusResponse, error)
}

type occurreceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOccurreceServiceClient(cc grpc.ClientConnInterface) OccurreceServiceClient {
	return &occurreceServiceClient{cc}
}

func (c *occurreceServiceClient) Create(ctx context.Context, in *CreateOccurrence, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.OccurreceService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *occurreceServiceClient) GetById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*OccurrenceDetails, error) {
	out := new(OccurrenceDetails)
	err := c.cc.Invoke(ctx, "/proto.OccurreceService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *occurreceServiceClient) EditById(ctx context.Context, in *EditOccurrence, opts ...grpc.CallOption) (*OccurrenceDetails, error) {
	out := new(OccurrenceDetails)
	err := c.cc.Invoke(ctx, "/proto.OccurreceService/EditById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *occurreceServiceClient) SoftDeleteById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.OccurreceService/SoftDeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *occurreceServiceClient) AcceptById(ctx context.Context, in *UpdateOccurrenceStatus, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.OccurreceService/AcceptById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *occurreceServiceClient) RefuseById(ctx context.Context, in *UpdateOccurrenceStatus, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.OccurreceService/RefuseById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *occurreceServiceClient) CancelById(ctx context.Context, in *UpdateOccurrenceStatus, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.OccurreceService/CancelById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OccurreceServiceServer is the server API for OccurreceService service.
// All implementations must embed UnimplementedOccurreceServiceServer
// for forward compatibility
type OccurreceServiceServer interface {
	Create(context.Context, *CreateOccurrence) (*StatusResponse, error)
	GetById(context.Context, *Id) (*OccurrenceDetails, error)
	EditById(context.Context, *EditOccurrence) (*OccurrenceDetails, error)
	SoftDeleteById(context.Context, *Id) (*StatusResponse, error)
	AcceptById(context.Context, *UpdateOccurrenceStatus) (*StatusResponse, error)
	RefuseById(context.Context, *UpdateOccurrenceStatus) (*StatusResponse, error)
	CancelById(context.Context, *UpdateOccurrenceStatus) (*StatusResponse, error)
	mustEmbedUnimplementedOccurreceServiceServer()
}

// UnimplementedOccurreceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOccurreceServiceServer struct {
}

func (UnimplementedOccurreceServiceServer) Create(context.Context, *CreateOccurrence) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOccurreceServiceServer) GetById(context.Context, *Id) (*OccurrenceDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedOccurreceServiceServer) EditById(context.Context, *EditOccurrence) (*OccurrenceDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditById not implemented")
}
func (UnimplementedOccurreceServiceServer) SoftDeleteById(context.Context, *Id) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftDeleteById not implemented")
}
func (UnimplementedOccurreceServiceServer) AcceptById(context.Context, *UpdateOccurrenceStatus) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptById not implemented")
}
func (UnimplementedOccurreceServiceServer) RefuseById(context.Context, *UpdateOccurrenceStatus) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefuseById not implemented")
}
func (UnimplementedOccurreceServiceServer) CancelById(context.Context, *UpdateOccurrenceStatus) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelById not implemented")
}
func (UnimplementedOccurreceServiceServer) mustEmbedUnimplementedOccurreceServiceServer() {}

// UnsafeOccurreceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OccurreceServiceServer will
// result in compilation errors.
type UnsafeOccurreceServiceServer interface {
	mustEmbedUnimplementedOccurreceServiceServer()
}

func RegisterOccurreceServiceServer(s grpc.ServiceRegistrar, srv OccurreceServiceServer) {
	s.RegisterService(&OccurreceService_ServiceDesc, srv)
}

func _OccurreceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOccurrence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccurreceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OccurreceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccurreceServiceServer).Create(ctx, req.(*CreateOccurrence))
	}
	return interceptor(ctx, in, info, handler)
}

func _OccurreceService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccurreceServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OccurreceService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccurreceServiceServer).GetById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _OccurreceService_EditById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditOccurrence)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccurreceServiceServer).EditById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OccurreceService/EditById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccurreceServiceServer).EditById(ctx, req.(*EditOccurrence))
	}
	return interceptor(ctx, in, info, handler)
}

func _OccurreceService_SoftDeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccurreceServiceServer).SoftDeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OccurreceService/SoftDeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccurreceServiceServer).SoftDeleteById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _OccurreceService_AcceptById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOccurrenceStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccurreceServiceServer).AcceptById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OccurreceService/AcceptById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccurreceServiceServer).AcceptById(ctx, req.(*UpdateOccurrenceStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _OccurreceService_RefuseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOccurrenceStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccurreceServiceServer).RefuseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OccurreceService/RefuseById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccurreceServiceServer).RefuseById(ctx, req.(*UpdateOccurrenceStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _OccurreceService_CancelById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOccurrenceStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OccurreceServiceServer).CancelById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OccurreceService/CancelById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OccurreceServiceServer).CancelById(ctx, req.(*UpdateOccurrenceStatus))
	}
	return interceptor(ctx, in, info, handler)
}

// OccurreceService_ServiceDesc is the grpc.ServiceDesc for OccurreceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OccurreceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OccurreceService",
	HandlerType: (*OccurreceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OccurreceService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _OccurreceService_GetById_Handler,
		},
		{
			MethodName: "EditById",
			Handler:    _OccurreceService_EditById_Handler,
		},
		{
			MethodName: "SoftDeleteById",
			Handler:    _OccurreceService_SoftDeleteById_Handler,
		},
		{
			MethodName: "AcceptById",
			Handler:    _OccurreceService_AcceptById_Handler,
		},
		{
			MethodName: "RefuseById",
			Handler:    _OccurreceService_RefuseById_Handler,
		},
		{
			MethodName: "CancelById",
			Handler:    _OccurreceService_CancelById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "occurrence.proto",
}