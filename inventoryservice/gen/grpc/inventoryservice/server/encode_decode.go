// Code generated by goa v3.11.3, DO NOT EDIT.
//
// inventoryservice gRPC server encoders and decoders
//
// Command:
// $ goa gen assignment_crossnokaye/cod/inventoryservice/design

package server

import (
	inventoryservicepb "assignment_crossnokaye/cod/inventoryservice/gen/grpc/inventoryservice/pb"
	inventoryservice "assignment_crossnokaye/cod/inventoryservice/gen/inventoryservice"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeGetResponse encodes responses from the "inventoryservice" service
// "Get" endpoint.
func EncodeGetResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(*inventoryservice.Inventory)
	if !ok {
		return nil, goagrpc.ErrInvalidType("inventoryservice", "Get", "*inventoryservice.Inventory", v)
	}
	resp := NewProtoGetResponse(result)
	return resp, nil
}

// DecodeGetRequest decodes requests sent to "inventoryservice" service "Get"
// endpoint.
func DecodeGetRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *inventoryservicepb.GetRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventoryservicepb.GetRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventoryservice", "Get", "*inventoryservicepb.GetRequest", v)
		}
	}
	var payload *inventoryservice.GetPayload
	{
		payload = NewGetPayload(message)
	}
	return payload, nil
}

// EncodeAddResponse encodes responses from the "inventoryservice" service
// "add" endpoint.
func EncodeAddResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	resp := NewProtoAddResponse()
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "inventoryservice" service "add"
// endpoint.
func DecodeAddRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *inventoryservicepb.AddRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventoryservicepb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventoryservice", "add", "*inventoryservicepb.AddRequest", v)
		}
	}
	var payload *inventoryservice.AddPayload
	{
		payload = NewAddPayload(message)
	}
	return payload, nil
}

// EncodeDeleteResponse encodes responses from the "inventoryservice" service
// "Delete" endpoint.
func EncodeDeleteResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	resp := NewProtoDeleteResponse()
	return resp, nil
}

// DecodeDeleteRequest decodes requests sent to "inventoryservice" service
// "Delete" endpoint.
func DecodeDeleteRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *inventoryservicepb.DeleteRequest
		ok      bool
	)
	{
		if message, ok = v.(*inventoryservicepb.DeleteRequest); !ok {
			return nil, goagrpc.ErrInvalidType("inventoryservice", "Delete", "*inventoryservicepb.DeleteRequest", v)
		}
	}
	var payload *inventoryservice.DeletePayload
	{
		payload = NewDeletePayload(message)
	}
	return payload, nil
}
