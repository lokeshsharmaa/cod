// Code generated by goa v3.11.3, DO NOT EDIT.
//
// externalservice HTTP client types
//
// Command:
// $ goa gen assignment_crossnokaye/cod/externalservice/design

package client

import (
	externalservice "assignment_crossnokaye/cod/externalservice/gen/externalservice"
)

// CreateCharacterRequestBody is the type of the "externalservice" service
// "create_character" endpoint HTTP request body.
type CreateCharacterRequestBody struct {
	// Character name
	Name string `form:"name" json:"name" xml:"name"`
	// Character description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Character health
	Health int `form:"health" json:"health" xml:"health"`
	// Character experience
	Experience int `form:"experience" json:"experience" xml:"experience"`
}

// UpdateCharacterRequestBody is the type of the "externalservice" service
// "update_character" endpoint HTTP request body.
type UpdateCharacterRequestBody struct {
	// Character name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Character description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Character health
	Health *int `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Character experience
	Experience *int `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// AddItemToInventoryRequestBody is the type of the "externalservice" service
// "add_item_to_inventory" endpoint HTTP request body.
type AddItemToInventoryRequestBody struct {
	// Item ID
	ItemID int `form:"item_id" json:"item_id" xml:"item_id"`
}

// CreateItemRequestBody is the type of the "externalservice" service
// "create_item" endpoint HTTP request body.
type CreateItemRequestBody struct {
	// Item name
	Name string `form:"name" json:"name" xml:"name"`
	// Item description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Item damage
	Damage int `form:"damage" json:"damage" xml:"damage"`
	// Item healing
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Item protection
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// UpdateItemRequestBody is the type of the "externalservice" service
// "update_item" endpoint HTTP request body.
type UpdateItemRequestBody struct {
	// Item name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Item description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Item damage
	Damage *int `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Item healing
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Item protection
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// CreateCharacterResponseBody is the type of the "externalservice" service
// "create_character" endpoint HTTP response body.
type CreateCharacterResponseBody struct {
	// Character ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Character name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Character description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Character health
	Health *int `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Character experience
	Experience *int `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// GetCharacterResponseBody is the type of the "externalservice" service
// "get_character" endpoint HTTP response body.
type GetCharacterResponseBody struct {
	// Character ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Character name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Character description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Character health
	Health *int `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Character experience
	Experience *int `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// GetInventoryResponseBody is the type of the "externalservice" service
// "get_inventory" endpoint HTTP response body.
type GetInventoryResponseBody struct {
	// Character ID
	CharacterID *int `form:"character_id,omitempty" json:"character_id,omitempty" xml:"character_id,omitempty"`
	// Item IDs
	Items []int `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
}

// CreateItemResponseBody is the type of the "externalservice" service
// "create_item" endpoint HTTP response body.
type CreateItemResponseBody struct {
	// Item ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Item name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Item description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Item damage
	Damage *int `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Item healing
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Item protection
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// GetItemResponseBody is the type of the "externalservice" service "get_item"
// endpoint HTTP response body.
type GetItemResponseBody struct {
	// Item ID
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Item name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Item description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Item damage
	Damage *int `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Item healing
	Healing *int `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Item protection
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// NewCreateCharacterRequestBody builds the HTTP request body from the payload
// of the "create_character" endpoint of the "externalservice" service.
func NewCreateCharacterRequestBody(p *externalservice.CreateCharacterPayload) *CreateCharacterRequestBody {
	body := &CreateCharacterRequestBody{
		Name:        p.Name,
		Description: p.Description,
		Health:      p.Health,
		Experience:  p.Experience,
	}
	return body
}

// NewUpdateCharacterRequestBody builds the HTTP request body from the payload
// of the "update_character" endpoint of the "externalservice" service.
func NewUpdateCharacterRequestBody(p *externalservice.UpdateCharacterPayload) *UpdateCharacterRequestBody {
	body := &UpdateCharacterRequestBody{
		Name:        p.Name,
		Description: p.Description,
		Health:      p.Health,
		Experience:  p.Experience,
	}
	return body
}

// NewAddItemToInventoryRequestBody builds the HTTP request body from the
// payload of the "add_item_to_inventory" endpoint of the "externalservice"
// service.
func NewAddItemToInventoryRequestBody(p *externalservice.AddItemToInventoryPayload) *AddItemToInventoryRequestBody {
	body := &AddItemToInventoryRequestBody{
		ItemID: p.ItemID,
	}
	return body
}

// NewCreateItemRequestBody builds the HTTP request body from the payload of
// the "create_item" endpoint of the "externalservice" service.
func NewCreateItemRequestBody(p *externalservice.CreateItemPayload) *CreateItemRequestBody {
	body := &CreateItemRequestBody{
		Name:        p.Name,
		Description: p.Description,
		Damage:      p.Damage,
		Healing:     p.Healing,
		Protection:  p.Protection,
	}
	return body
}

// NewUpdateItemRequestBody builds the HTTP request body from the payload of
// the "update_item" endpoint of the "externalservice" service.
func NewUpdateItemRequestBody(p *externalservice.UpdateItemPayload) *UpdateItemRequestBody {
	body := &UpdateItemRequestBody{
		Name:        p.Name,
		Description: p.Description,
		Damage:      p.Damage,
		Healing:     p.Healing,
		Protection:  p.Protection,
	}
	return body
}

// NewCreateCharacterCharacterCreated builds a "externalservice" service
// "create_character" endpoint result from a HTTP "Created" response.
func NewCreateCharacterCharacterCreated(body *CreateCharacterResponseBody) *externalservice.Character {
	v := &externalservice.Character{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Health:      body.Health,
		Experience:  body.Experience,
	}

	return v
}

// NewGetCharacterCharacterOK builds a "externalservice" service
// "get_character" endpoint result from a HTTP "OK" response.
func NewGetCharacterCharacterOK(body *GetCharacterResponseBody) *externalservice.Character {
	v := &externalservice.Character{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Health:      body.Health,
		Experience:  body.Experience,
	}

	return v
}

// NewGetInventoryInventoryOK builds a "externalservice" service
// "get_inventory" endpoint result from a HTTP "OK" response.
func NewGetInventoryInventoryOK(body *GetInventoryResponseBody) *externalservice.Inventory {
	v := &externalservice.Inventory{
		CharacterID: body.CharacterID,
	}
	if body.Items != nil {
		v.Items = make([]int, len(body.Items))
		for i, val := range body.Items {
			v.Items[i] = val
		}
	}

	return v
}

// NewCreateItemItemCreated builds a "externalservice" service "create_item"
// endpoint result from a HTTP "Created" response.
func NewCreateItemItemCreated(body *CreateItemResponseBody) *externalservice.Item {
	v := &externalservice.Item{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Damage:      body.Damage,
		Healing:     body.Healing,
		Protection:  body.Protection,
	}

	return v
}

// NewGetItemItemOK builds a "externalservice" service "get_item" endpoint
// result from a HTTP "OK" response.
func NewGetItemItemOK(body *GetItemResponseBody) *externalservice.Item {
	v := &externalservice.Item{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Damage:      body.Damage,
		Healing:     body.Healing,
		Protection:  body.Protection,
	}

	return v
}
