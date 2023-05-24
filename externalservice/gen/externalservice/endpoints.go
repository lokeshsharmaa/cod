// Code generated by goa v3.11.3, DO NOT EDIT.
//
// externalservice endpoints
//
// Command:
// $ goa gen assignment_crossnokaye/cod/externalservice/design

package externalservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "externalservice" service endpoints.
type Endpoints struct {
	CreateCharacter         goa.Endpoint
	GetCharacter            goa.Endpoint
	UpdateCharacter         goa.Endpoint
	DeleteCharacter         goa.Endpoint
	GetInventory            goa.Endpoint
	AddItemToInventory      goa.Endpoint
	RemoveItemFromInventory goa.Endpoint
}

// NewEndpoints wraps the methods of the "externalservice" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		CreateCharacter:         NewCreateCharacterEndpoint(s),
		GetCharacter:            NewGetCharacterEndpoint(s),
		UpdateCharacter:         NewUpdateCharacterEndpoint(s),
		DeleteCharacter:         NewDeleteCharacterEndpoint(s),
		GetInventory:            NewGetInventoryEndpoint(s),
		AddItemToInventory:      NewAddItemToInventoryEndpoint(s),
		RemoveItemFromInventory: NewRemoveItemFromInventoryEndpoint(s),
	}
}

// Use applies the given middleware to all the "externalservice" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.CreateCharacter = m(e.CreateCharacter)
	e.GetCharacter = m(e.GetCharacter)
	e.UpdateCharacter = m(e.UpdateCharacter)
	e.DeleteCharacter = m(e.DeleteCharacter)
	e.GetInventory = m(e.GetInventory)
	e.AddItemToInventory = m(e.AddItemToInventory)
	e.RemoveItemFromInventory = m(e.RemoveItemFromInventory)
}

// NewCreateCharacterEndpoint returns an endpoint function that calls the
// method "create_character" of service "externalservice".
func NewCreateCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateCharacterPayload)
		return s.CreateCharacter(ctx, p)
	}
}

// NewGetCharacterEndpoint returns an endpoint function that calls the method
// "get_character" of service "externalservice".
func NewGetCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetCharacterPayload)
		return s.GetCharacter(ctx, p)
	}
}

// NewUpdateCharacterEndpoint returns an endpoint function that calls the
// method "update_character" of service "externalservice".
func NewUpdateCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdateCharacterPayload)
		return nil, s.UpdateCharacter(ctx, p)
	}
}

// NewDeleteCharacterEndpoint returns an endpoint function that calls the
// method "delete_character" of service "externalservice".
func NewDeleteCharacterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteCharacterPayload)
		return nil, s.DeleteCharacter(ctx, p)
	}
}

// NewGetInventoryEndpoint returns an endpoint function that calls the method
// "get_inventory" of service "externalservice".
func NewGetInventoryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetInventoryPayload)
		return s.GetInventory(ctx, p)
	}
}

// NewAddItemToInventoryEndpoint returns an endpoint function that calls the
// method "add_item_to_inventory" of service "externalservice".
func NewAddItemToInventoryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*AddItemToInventoryPayload)
		return nil, s.AddItemToInventory(ctx, p)
	}
}

// NewRemoveItemFromInventoryEndpoint returns an endpoint function that calls
// the method "remove_item_from_inventory" of service "externalservice".
func NewRemoveItemFromInventoryEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*RemoveItemFromInventoryPayload)
		return nil, s.RemoveItemFromInventory(ctx, p)
	}
}
