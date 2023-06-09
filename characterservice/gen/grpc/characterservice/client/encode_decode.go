// Code generated by goa v3.11.3, DO NOT EDIT.
//
// characterservice gRPC client encoders and decoders
//
// Command:
// $ goa gen assignment_crossnokaye/cod/characterservice/design

package client

import (
	characterservice "assignment_crossnokaye/cod/characterservice/gen/characterservice"
	characterservicepb "assignment_crossnokaye/cod/characterservice/gen/grpc/characterservice/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildCreateFunc builds the remote method to invoke for "characterservice"
// service "Create" endpoint.
func BuildCreateFunc(grpccli characterservicepb.CharacterserviceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Create(ctx, reqpb.(*characterservicepb.CreateRequest), opts...)
		}
		return grpccli.Create(ctx, &characterservicepb.CreateRequest{}, opts...)
	}
}

// EncodeCreateRequest encodes requests sent to characterservice Create
// endpoint.
func EncodeCreateRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*characterservice.CreatePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("characterservice", "Create", "*characterservice.CreatePayload", v)
	}
	return NewProtoCreateRequest(payload), nil
}

// DecodeCreateResponse decodes responses from the characterservice Create
// endpoint.
func DecodeCreateResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*characterservicepb.CreateResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("characterservice", "Create", "*characterservicepb.CreateResponse", v)
	}
	res := NewCreateResult(message)
	return res, nil
}

// BuildGetFunc builds the remote method to invoke for "characterservice"
// service "Get" endpoint.
func BuildGetFunc(grpccli characterservicepb.CharacterserviceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Get(ctx, reqpb.(*characterservicepb.GetRequest), opts...)
		}
		return grpccli.Get(ctx, &characterservicepb.GetRequest{}, opts...)
	}
}

// EncodeGetRequest encodes requests sent to characterservice Get endpoint.
func EncodeGetRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*characterservice.GetPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("characterservice", "Get", "*characterservice.GetPayload", v)
	}
	return NewProtoGetRequest(payload), nil
}

// DecodeGetResponse decodes responses from the characterservice Get endpoint.
func DecodeGetResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*characterservicepb.GetResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("characterservice", "Get", "*characterservicepb.GetResponse", v)
	}
	res := NewGetResult(message)
	return res, nil
}

// BuildUpdateFunc builds the remote method to invoke for "characterservice"
// service "Update" endpoint.
func BuildUpdateFunc(grpccli characterservicepb.CharacterserviceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Update(ctx, reqpb.(*characterservicepb.UpdateRequest), opts...)
		}
		return grpccli.Update(ctx, &characterservicepb.UpdateRequest{}, opts...)
	}
}

// EncodeUpdateRequest encodes requests sent to characterservice Update
// endpoint.
func EncodeUpdateRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*characterservice.UpdatePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("characterservice", "Update", "*characterservice.UpdatePayload", v)
	}
	return NewProtoUpdateRequest(payload), nil
}

// BuildDeleteFunc builds the remote method to invoke for "characterservice"
// service "Delete" endpoint.
func BuildDeleteFunc(grpccli characterservicepb.CharacterserviceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Delete(ctx, reqpb.(*characterservicepb.DeleteRequest), opts...)
		}
		return grpccli.Delete(ctx, &characterservicepb.DeleteRequest{}, opts...)
	}
}

// EncodeDeleteRequest encodes requests sent to characterservice Delete
// endpoint.
func EncodeDeleteRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*characterservice.DeletePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("characterservice", "Delete", "*characterservice.DeletePayload", v)
	}
	return NewProtoDeleteRequest(payload), nil
}
