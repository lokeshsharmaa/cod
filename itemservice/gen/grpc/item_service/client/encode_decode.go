// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC client encoders and decoders
//
// Command:
// $ goa gen assignment_crossnokaye/cod/itemservice/design

package client

import (
	item_servicepb "assignment_crossnokaye/cod/itemservice/gen/grpc/item_service/pb"
	itemservice "assignment_crossnokaye/cod/itemservice/gen/item_service"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildCreateFunc builds the remote method to invoke for "ItemService" service
// "Create" endpoint.
func BuildCreateFunc(grpccli item_servicepb.ItemServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Create(ctx, reqpb.(*item_servicepb.CreateRequest), opts...)
		}
		return grpccli.Create(ctx, &item_servicepb.CreateRequest{}, opts...)
	}
}

// EncodeCreateRequest encodes requests sent to ItemService Create endpoint.
func EncodeCreateRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*itemservice.CreatePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "Create", "*itemservice.CreatePayload", v)
	}
	return NewProtoCreateRequest(payload), nil
}

// DecodeCreateResponse decodes responses from the ItemService Create endpoint.
func DecodeCreateResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*item_servicepb.CreateResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "Create", "*item_servicepb.CreateResponse", v)
	}
	res := NewCreateResult(message)
	return res, nil
}

// BuildGetFunc builds the remote method to invoke for "ItemService" service
// "Get" endpoint.
func BuildGetFunc(grpccli item_servicepb.ItemServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Get(ctx, reqpb.(*item_servicepb.GetRequest), opts...)
		}
		return grpccli.Get(ctx, &item_servicepb.GetRequest{}, opts...)
	}
}

// EncodeGetRequest encodes requests sent to ItemService Get endpoint.
func EncodeGetRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*itemservice.GetPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "Get", "*itemservice.GetPayload", v)
	}
	return NewProtoGetRequest(payload), nil
}

// DecodeGetResponse decodes responses from the ItemService Get endpoint.
func DecodeGetResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*item_servicepb.GetResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "Get", "*item_servicepb.GetResponse", v)
	}
	res := NewGetResult(message)
	return res, nil
}

// BuildUpdateFunc builds the remote method to invoke for "ItemService" service
// "Update" endpoint.
func BuildUpdateFunc(grpccli item_servicepb.ItemServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Update(ctx, reqpb.(*item_servicepb.UpdateRequest), opts...)
		}
		return grpccli.Update(ctx, &item_servicepb.UpdateRequest{}, opts...)
	}
}

// EncodeUpdateRequest encodes requests sent to ItemService Update endpoint.
func EncodeUpdateRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*itemservice.UpdatePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "Update", "*itemservice.UpdatePayload", v)
	}
	return NewProtoUpdateRequest(payload), nil
}

// BuildDeleteFunc builds the remote method to invoke for "ItemService" service
// "Delete" endpoint.
func BuildDeleteFunc(grpccli item_servicepb.ItemServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Delete(ctx, reqpb.(*item_servicepb.DeleteRequest), opts...)
		}
		return grpccli.Delete(ctx, &item_servicepb.DeleteRequest{}, opts...)
	}
}

// EncodeDeleteRequest encodes requests sent to ItemService Delete endpoint.
func EncodeDeleteRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*itemservice.DeletePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ItemService", "Delete", "*itemservice.DeletePayload", v)
	}
	return NewProtoDeleteRequest(payload), nil
}