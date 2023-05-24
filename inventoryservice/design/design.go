package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("inventoryservice", func() {
	Title("Inventories Service")
	Description("Managing inventory for a character")
	Server("inventoryservice", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8082")
		})
	})
})

var _ = Service("inventoryservice", func() {
	Description("The inventoryservice service handles CRUD operations for character inventories.")

	Method("Get", func() {
		Description("Retrieve an inventory.")
		Payload(func() {
			Field(1, "id", Int, "Character ID")
			Required("id")
		})
		Result(Inventory, "inventory", "Retrieved inventory")
		Error("not_found", String, "Inventory not found")
		Error("internal_error", String, "Internal server error")
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("add", func() {
		Description("Update an inventory.")
		Payload(func() {
			Field(1, "character_id", Int, "Character ID")
			Field(2, "item_id", Int, "Item ID")
			Required("character_id", "character_id")
		})
		Error("not_found", String, "Inventory not found")
		Error("bad_request", String, "Bad request")
		Error("internal_error", String, "Internal server error")
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("Delete", func() {
		Description("Delete an inventory by ID.")
		Payload(func() {
			Field(1, "character_id", Int, "Character ID")
			Field(2, "item_id", Int, "Item ID")
			Required("character_id", "item_id")
		})
		Error("not_found", String, "Inventory not found")
		Error("internal_error", String, "Internal server error")
		GRPC(func() {
			Response(CodeOK)
		})
	})

})

var Inventory = Type("Inventory", func() {
	Description("Inventory holds the association of items with a character.")
	Field(1, "character_id", Int, "Character ID")
	Field(2, "items", ArrayOf(Int), "Items in the inventory")
})
