// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC server
//
// Command:
// $ goa gen assignment_crossnokaye/cod/itemservice/design

package server

import (
	item_servicepb "assignment_crossnokaye/cod/itemservice/gen/grpc/item_service/pb"
	itemservice "assignment_crossnokaye/cod/itemservice/gen/item_service"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the item_servicepb.ItemServiceServer interface.
type Server struct {
	CreateH goagrpc.UnaryHandler
	GetH    goagrpc.UnaryHandler
	UpdateH goagrpc.UnaryHandler
	DeleteH goagrpc.UnaryHandler
	item_servicepb.UnimplementedItemServiceServer
}

// New instantiates the server struct with the ItemService service endpoints.
func New(e *itemservice.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		CreateH: NewCreateHandler(e.Create, uh),
		GetH:    NewGetHandler(e.Get, uh),
		UpdateH: NewUpdateHandler(e.Update, uh),
		DeleteH: NewDeleteHandler(e.Delete, uh),
	}
}

// NewCreateHandler creates a gRPC handler which serves the "ItemService"
// service "Create" endpoint.
func NewCreateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateRequest, EncodeCreateResponse)
	}
	return h
}

// Create implements the "Create" method in item_servicepb.ItemServiceServer
// interface.
func (s *Server) Create(ctx context.Context, message *item_servicepb.CreateRequest) (*item_servicepb.CreateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "Create")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ItemService")
	resp, err := s.CreateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*item_servicepb.CreateResponse), nil
}

// NewGetHandler creates a gRPC handler which serves the "ItemService" service
// "Get" endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetRequest, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in item_servicepb.ItemServiceServer
// interface.
func (s *Server) Get(ctx context.Context, message *item_servicepb.GetRequest) (*item_servicepb.GetResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "Get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ItemService")
	resp, err := s.GetH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*item_servicepb.GetResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "ItemService"
// service "Update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in item_servicepb.ItemServiceServer
// interface.
func (s *Server) Update(ctx context.Context, message *item_servicepb.UpdateRequest) (*item_servicepb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "Update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ItemService")
	resp, err := s.UpdateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*item_servicepb.UpdateResponse), nil
}

// NewDeleteHandler creates a gRPC handler which serves the "ItemService"
// service "Delete" endpoint.
func NewDeleteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDeleteRequest, EncodeDeleteResponse)
	}
	return h
}

// Delete implements the "Delete" method in item_servicepb.ItemServiceServer
// interface.
func (s *Server) Delete(ctx context.Context, message *item_servicepb.DeleteRequest) (*item_servicepb.DeleteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "Delete")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ItemService")
	resp, err := s.DeleteH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*item_servicepb.DeleteResponse), nil
}
