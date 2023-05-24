// Code generated by goa v3.11.3, DO NOT EDIT.
//
// HTTP request path constructors for the externalservice service.
//
// Command:
// $ goa gen assignment_crossnokaye/cod/externalservice/design

package server

import (
	"fmt"
)

// CreateCharacterExternalservicePath returns the URL path to the externalservice service create_character HTTP endpoint.
func CreateCharacterExternalservicePath() string {
	return "/characters"
}

// GetCharacterExternalservicePath returns the URL path to the externalservice service get_character HTTP endpoint.
func GetCharacterExternalservicePath(id int) string {
	return fmt.Sprintf("/characters/%v", id)
}

// UpdateCharacterExternalservicePath returns the URL path to the externalservice service update_character HTTP endpoint.
func UpdateCharacterExternalservicePath(id int) string {
	return fmt.Sprintf("/characters/%v", id)
}

// DeleteCharacterExternalservicePath returns the URL path to the externalservice service delete_character HTTP endpoint.
func DeleteCharacterExternalservicePath(id int) string {
	return fmt.Sprintf("/characters/%v", id)
}

// GetInventoryExternalservicePath returns the URL path to the externalservice service get_inventory HTTP endpoint.
func GetInventoryExternalservicePath(characterID int) string {
	return fmt.Sprintf("/characters/%v/inventory", characterID)
}

// AddItemToInventoryExternalservicePath returns the URL path to the externalservice service add_item_to_inventory HTTP endpoint.
func AddItemToInventoryExternalservicePath(characterID int) string {
	return fmt.Sprintf("/characters/%v/inventory/items", characterID)
}

// RemoveItemFromInventoryExternalservicePath returns the URL path to the externalservice service remove_item_from_inventory HTTP endpoint.
func RemoveItemFromInventoryExternalservicePath(characterID int, itemID int) string {
	return fmt.Sprintf("/characters/%v/inventory/items/%v", characterID, itemID)
}

// CreateItemExternalservicePath returns the URL path to the externalservice service create_item HTTP endpoint.
func CreateItemExternalservicePath() string {
	return "/items"
}

// GetItemExternalservicePath returns the URL path to the externalservice service get_item HTTP endpoint.
func GetItemExternalservicePath(id int) string {
	return fmt.Sprintf("/items/%v", id)
}

// UpdateItemExternalservicePath returns the URL path to the externalservice service update_item HTTP endpoint.
func UpdateItemExternalservicePath(id int) string {
	return fmt.Sprintf("/items/%v", id)
}

// DeleteItemExternalservicePath returns the URL path to the externalservice service delete_item HTTP endpoint.
func DeleteItemExternalservicePath(id int) string {
	return fmt.Sprintf("/items/%v", id)
}
