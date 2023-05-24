// Code generated by goa v3.11.3, DO NOT EDIT.
//
// inventoryservice gRPC server
//
// Command:
// $ goa gen assignment_crossnokaye/cod/inventoryservice/design

package server

import (
	inventoryservicepb "assignment_crossnokaye/cod/inventoryservice/gen/grpc/inventoryservice/pb"
	inventoryservice "assignment_crossnokaye/cod/inventoryservice/gen/inventoryservice"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the inventoryservicepb.InventoryserviceServer interface.
type Server struct {
	GetH    goagrpc.UnaryHandler
	AddH    goagrpc.UnaryHandler
	DeleteH goagrpc.UnaryHandler
	inventoryservicepb.UnimplementedInventoryserviceServer
}

// New instantiates the server struct with the inventoryservice service
// endpoints.
func New(e *inventoryservice.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		GetH:    NewGetHandler(e.Get, uh),
		AddH:    NewAddHandler(e.Add, uh),
		DeleteH: NewDeleteHandler(e.Delete, uh),
	}
}

// NewGetHandler creates a gRPC handler which serves the "inventoryservice"
// service "Get" endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetRequest, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in inventoryservicepb.InventoryserviceServer
// interface.
func (s *Server) Get(ctx context.Context, message *inventoryservicepb.GetRequest) (*inventoryservicepb.GetResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "Get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventoryservice")
	resp, err := s.GetH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventoryservicepb.GetResponse), nil
}

// NewAddHandler creates a gRPC handler which serves the "inventoryservice"
// service "add" endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in inventoryservicepb.InventoryserviceServer
// interface.
func (s *Server) Add(ctx context.Context, message *inventoryservicepb.AddRequest) (*inventoryservicepb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventoryservice")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventoryservicepb.AddResponse), nil
}

// NewDeleteHandler creates a gRPC handler which serves the "inventoryservice"
// service "Delete" endpoint.
func NewDeleteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDeleteRequest, EncodeDeleteResponse)
	}
	return h
}

// Delete implements the "Delete" method in
// inventoryservicepb.InventoryserviceServer interface.
func (s *Server) Delete(ctx context.Context, message *inventoryservicepb.DeleteRequest) (*inventoryservicepb.DeleteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "Delete")
	ctx = context.WithValue(ctx, goa.ServiceKey, "inventoryservice")
	resp, err := s.DeleteH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*inventoryservicepb.DeleteResponse), nil
}