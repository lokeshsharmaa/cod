package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("ItemService", func() {
	Title("Items Service")
	Description("CRUD operations on items")
	Server("ItemService", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8083")
		})
	})
})

var _ = Service("ItemService", func() {
	Description("The ItemService service handles CRUD operations for items")

	Method("Create", func() {
		Description("Create a new item")
		Payload(func() {
			Field(1, "name", String, "Item name")
			Field(2, "description", String, "Item description")
			Field(3, "damage", Int, "Item damage")
			Field(4, "healing", Int, "Item healing")
			Field(5, "protection", Int, "Item protection")
			Required("name", "damage")
		})
		Result(Int, "id", "ID of the created item")
		Error("unique_constraint", String, "Name not unique")
		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("Get", func() {
		Description("Retrieve an item by ID")
		Payload(func() {
			Field(1, "id", Int, "Item ID")
			Required("id")
		})
		Result(Item, "item", "Retrieved item")
		Error("not_found", String, "Item not found")

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("Update", func() {
		Description("Update an item")
		Payload(func() {
			Field(1, "id", Int, "Item ID")
			Field(2, "name", String, "Item name")
			Field(3, "description", String, "Item description")
			Field(4, "damage", Int, "Item damage")
			Field(5, "healing", Int, "Item healing")
			Field(6, "protection", Int, "Item protection")
			Required("id")
		})
		Error("not_found", String, "Item not found")
		Error("unique_constraint", String, "Name not unique")

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("Delete", func() {
		Description("Delete an item by ID")
		Payload(func() {
			Field(1, "id", Int, "Item ID")
			Required("id")
		})
		Error("not_found", String, "Item not found")
		Error("unique_constraint", String, "Name not unique")

		GRPC(func() {
			Response(CodeOK)
		})
	})

})

var Item = Type("Item", func() {
	Description("Attributes of an item")
	Field(1, "id", Int, "Item ID")
	Field(2, "name", String, "Item name")
	Field(3, "description", String, "Item description")
	Field(4, "damage", Int, "Item damage")
	Field(5, "healing", Int, "Item healing")
	Field(6, "protection", Int, "Item protection")
})
