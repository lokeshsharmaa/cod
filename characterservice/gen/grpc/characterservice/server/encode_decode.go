// Code generated by goa v3.11.3, DO NOT EDIT.
//
// characterservice gRPC server encoders and decoders
//
// Command:
// $ goa gen assignment_crossnokaye/cod/characterservice/design

package server

import (
	characterservice "assignment_crossnokaye/cod/characterservice/gen/characterservice"
	characterservicepb "assignment_crossnokaye/cod/characterservice/gen/grpc/characterservice/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeCreateResponse encodes responses from the "characterservice" service
// "Create" endpoint.
func EncodeCreateResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("characterservice", "Create", "int", v)
	}
	resp := NewProtoCreateResponse(result)
	return resp, nil
}

// DecodeCreateRequest decodes requests sent to "characterservice" service
// "Create" endpoint.
func DecodeCreateRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *characterservicepb.CreateRequest
		ok      bool
	)
	{
		if message, ok = v.(*characterservicepb.CreateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("characterservice", "Create", "*characterservicepb.CreateRequest", v)
		}
	}
	var payload *characterservice.CreatePayload
	{
		payload = NewCreatePayload(message)
	}
	return payload, nil
}

// EncodeGetResponse encodes responses from the "characterservice" service
// "Get" endpoint.
func EncodeGetResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	result, ok := v.(*characterservice.Character)
	if !ok {
		return nil, goagrpc.ErrInvalidType("characterservice", "Get", "*characterservice.Character", v)
	}
	resp := NewProtoGetResponse(result)
	return resp, nil
}

// DecodeGetRequest decodes requests sent to "characterservice" service "Get"
// endpoint.
func DecodeGetRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *characterservicepb.GetRequest
		ok      bool
	)
	{
		if message, ok = v.(*characterservicepb.GetRequest); !ok {
			return nil, goagrpc.ErrInvalidType("characterservice", "Get", "*characterservicepb.GetRequest", v)
		}
	}
	var payload *characterservice.GetPayload
	{
		payload = NewGetPayload(message)
	}
	return payload, nil
}

// EncodeUpdateResponse encodes responses from the "characterservice" service
// "Update" endpoint.
func EncodeUpdateResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	resp := NewProtoUpdateResponse()
	return resp, nil
}

// DecodeUpdateRequest decodes requests sent to "characterservice" service
// "Update" endpoint.
func DecodeUpdateRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *characterservicepb.UpdateRequest
		ok      bool
	)
	{
		if message, ok = v.(*characterservicepb.UpdateRequest); !ok {
			return nil, goagrpc.ErrInvalidType("characterservice", "Update", "*characterservicepb.UpdateRequest", v)
		}
	}
	var payload *characterservice.UpdatePayload
	{
		payload = NewUpdatePayload(message)
	}
	return payload, nil
}

// EncodeDeleteResponse encodes responses from the "characterservice" service
// "Delete" endpoint.
func EncodeDeleteResponse(ctx context.Context, v any, hdr, trlr *metadata.MD) (any, error) {
	resp := NewProtoDeleteResponse()
	return resp, nil
}

// DecodeDeleteRequest decodes requests sent to "characterservice" service
// "Delete" endpoint.
func DecodeDeleteRequest(ctx context.Context, v any, md metadata.MD) (any, error) {
	var (
		message *characterservicepb.DeleteRequest
		ok      bool
	)
	{
		if message, ok = v.(*characterservicepb.DeleteRequest); !ok {
			return nil, goagrpc.ErrInvalidType("characterservice", "Delete", "*characterservicepb.DeleteRequest", v)
		}
	}
	var payload *characterservice.DeletePayload
	{
		payload = NewDeletePayload(message)
	}
	return payload, nil
}
