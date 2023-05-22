// Code generated by goa v3.11.3, DO NOT EDIT.
//
// characterservice gRPC client
//
// Command:
// $ goa gen assignment_crossnokaye/cod/characterservice/design

package client

import (
	characterservicepb "assignment_crossnokaye/cod/characterservice/gen/grpc/characterservice/pb"
	"context"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli characterservicepb.CharacterserviceClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the characterservice service
// servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: characterservicepb.NewCharacterserviceClient(cc),
		opts:    opts,
	}
}

// Create calls the "Create" function in
// characterservicepb.CharacterserviceClient interface.
func (c *Client) Create() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildCreateFunc(c.grpccli, c.opts...),
			EncodeCreateRequest,
			DecodeCreateResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Get calls the "Get" function in characterservicepb.CharacterserviceClient
// interface.
func (c *Client) Get() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildGetFunc(c.grpccli, c.opts...),
			EncodeGetRequest,
			DecodeGetResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Update calls the "Update" function in
// characterservicepb.CharacterserviceClient interface.
func (c *Client) Update() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildUpdateFunc(c.grpccli, c.opts...),
			EncodeUpdateRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Delete calls the "Delete" function in
// characterservicepb.CharacterserviceClient interface.
func (c *Client) Delete() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildDeleteFunc(c.grpccli, c.opts...),
			EncodeDeleteRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
