// Code generated with goa v3.11.3, DO NOT EDIT.
//
// inventoryservice protocol buffer definition
//
// Command:
// $ goa gen assignment_crossnokaye/cod/inventoryservice/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.1
// source: goadesign_goagen_inventoryservice.proto

package inventoryservicepb

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

const (
	Inventoryservice_Get_FullMethodName    = "/inventoryservice.Inventoryservice/Get"
	Inventoryservice_Add_FullMethodName    = "/inventoryservice.Inventoryservice/Add"
	Inventoryservice_Delete_FullMethodName = "/inventoryservice.Inventoryservice/Delete"
)

// InventoryserviceClient is the client API for Inventoryservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryserviceClient interface {
	// Retrieve an inventory.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Update an inventory.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Delete an inventory by ID.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type inventoryserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryserviceClient(cc grpc.ClientConnInterface) InventoryserviceClient {
	return &inventoryserviceClient{cc}
}

func (c *inventoryserviceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Inventoryservice_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryserviceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, Inventoryservice_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryserviceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, Inventoryservice_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryserviceServer is the server API for Inventoryservice service.
// All implementations must embed UnimplementedInventoryserviceServer
// for forward compatibility
type InventoryserviceServer interface {
	// Retrieve an inventory.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Update an inventory.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Delete an inventory by ID.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedInventoryserviceServer()
}

// UnimplementedInventoryserviceServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryserviceServer struct {
}

func (UnimplementedInventoryserviceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedInventoryserviceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedInventoryserviceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedInventoryserviceServer) mustEmbedUnimplementedInventoryserviceServer() {}

// UnsafeInventoryserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryserviceServer will
// result in compilation errors.
type UnsafeInventoryserviceServer interface {
	mustEmbedUnimplementedInventoryserviceServer()
}

func RegisterInventoryserviceServer(s grpc.ServiceRegistrar, srv InventoryserviceServer) {
	s.RegisterService(&Inventoryservice_ServiceDesc, srv)
}

func _Inventoryservice_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryserviceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventoryservice_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryserviceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventoryservice_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryserviceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventoryservice_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryserviceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inventoryservice_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryserviceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inventoryservice_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryserviceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inventoryservice_ServiceDesc is the grpc.ServiceDesc for Inventoryservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inventoryservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventoryservice.Inventoryservice",
	HandlerType: (*InventoryserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Inventoryservice_Get_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Inventoryservice_Add_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Inventoryservice_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goadesign_goagen_inventoryservice.proto",
}
