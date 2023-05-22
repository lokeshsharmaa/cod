// Code generated by goa v3.11.3, DO NOT EDIT.
//
// characterservice client
//
// Command:
// $ goa gen assignment_crossnokaye/cod/characterservice/design

package characterservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "characterservice" service client.
type Client struct {
	CreateEndpoint goa.Endpoint
	GetEndpoint    goa.Endpoint
	UpdateEndpoint goa.Endpoint
	DeleteEndpoint goa.Endpoint
}

// NewClient initializes a "characterservice" service client given the
// endpoints.
func NewClient(create, get, update, delete_ goa.Endpoint) *Client {
	return &Client{
		CreateEndpoint: create,
		GetEndpoint:    get,
		UpdateEndpoint: update,
		DeleteEndpoint: delete_,
	}
}

// Create calls the "Create" endpoint of the "characterservice" service.
// Create may return the following errors:
//   - "unique_constraint" (type UniqueConstraint)
//   - error: internal error
func (c *Client) Create(ctx context.Context, p *CreatePayload) (res int, err error) {
	var ires any
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(int), nil
}

// Get calls the "Get" endpoint of the "characterservice" service.
// Get may return the following errors:
//   - "not_found" (type NotFound)
//   - "unique_constraint" (type UniqueConstraint)
//   - error: internal error
func (c *Client) Get(ctx context.Context, p *GetPayload) (res *Character, err error) {
	var ires any
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// Update calls the "Update" endpoint of the "characterservice" service.
// Update may return the following errors:
//   - "not_found" (type NotFound)
//   - "unique_constraint" (type UniqueConstraint)
//   - error: internal error
func (c *Client) Update(ctx context.Context, p *UpdatePayload) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}

// Delete calls the "Delete" endpoint of the "characterservice" service.
// Delete may return the following errors:
//   - "not_found" (type NotFound)
//   - error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}
