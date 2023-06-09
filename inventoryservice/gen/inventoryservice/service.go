// Code generated by goa v3.11.3, DO NOT EDIT.
//
// inventoryservice service
//
// Command:
// $ goa gen assignment_crossnokaye/cod/inventoryservice/design

package inventoryservice

import (
	"context"
)

// The inventoryservice service handles CRUD operations for character
// inventories.
type Service interface {
	// Retrieve an inventory.
	Get(context.Context, *GetPayload) (res *Inventory, err error)
	// Update an inventory.
	Add(context.Context, *AddPayload) (err error)
	// Delete an inventory by ID.
	Delete(context.Context, *DeletePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "inventoryservice"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"Get", "add", "Delete"}

// AddPayload is the payload type of the inventoryservice service add method.
type AddPayload struct {
	// Character ID
	CharacterID int
	// Item ID
	ItemID *int
}

// DeletePayload is the payload type of the inventoryservice service Delete
// method.
type DeletePayload struct {
	// Character ID
	CharacterID int
	// Item ID
	ItemID int
}

// GetPayload is the payload type of the inventoryservice service Get method.
type GetPayload struct {
	// Character ID
	ID int
}

// inventory
type Inventory struct {
	// Character ID
	CharacterID *int
	// Items in the inventory
	Items []int
}

// Bad request
type BadRequest string

// Internal server error
type InternalError string

// Inventory not found
type NotFound string

// Error returns an error description.
func (e BadRequest) Error() string {
	return "Bad request"
}

// ErrorName returns "bad_request".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e BadRequest) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "bad_request".
func (e BadRequest) GoaErrorName() string {
	return "bad_request"
}

// Error returns an error description.
func (e InternalError) Error() string {
	return "Internal server error"
}

// ErrorName returns "internal_error".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e InternalError) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "internal_error".
func (e InternalError) GoaErrorName() string {
	return "internal_error"
}

// Error returns an error description.
func (e NotFound) Error() string {
	return "Inventory not found"
}

// ErrorName returns "not_found".
//
// Deprecated: Use GoaErrorName - https://github.com/goadesign/goa/issues/3105
func (e NotFound) ErrorName() string {
	return e.GoaErrorName()
}

// GoaErrorName returns "not_found".
func (e NotFound) GoaErrorName() string {
	return "not_found"
}
