package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("externalservice", func() {
	Title("External Service to interact with Character on Duty game")
	Description(" HTTP/JSON front service to manipulate characters, inventories, and items")
	Server("externalservice", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var Character = Type("Character", func() {
	Description("Character in the Game")
	Field(1, "id", Int, "Character ID")
	Field(2, "name", String, "Character name")
	Field(3, "description", String, "Character description")
	Field(4, "health", Int, "Character health")
	Field(5, "experience", Int, "Character experience")
})

var Inventory = Type("Inventory", func() {
	Description("Character's Inventory")
	Field(1, "character_id", Int, "Character ID")
	Field(2, "items", ArrayOf(Int), "Item IDs")
})

var _ = Service("externalservice", func() {
	Description("Contains API CRUD operations on characters, inventories, and items.")

	// Creating character with unique name and required payload
	Method("create_character", func() {
		Description("Creating a new Character")
		Payload(func() {
			Field(1, "name", String, "Character name")
			Field(2, "description", String, "Character description")
			Field(3, "health", Int, "Character health")
			Field(4, "experience", Int, "Character experience")
			Required("name", "health", "experience")
		})
		Result(Character)
		Error("bad_request", String, "Bad request")
		Error("unique_constraint", String, "Name not unique")
		HTTP(func() {
			POST("/characters")
			Response(StatusCreated)
		})
	})

	// Get a character by ID
	Method("get_character", func() {
		Description("Fetching a Character")
		Payload(func() {
			Field(1, "id", Int, "Character ID")
			Required("id")
		})
		Result(Character)
		Error("not_found", String, "Character not found")
		HTTP(func() {
			GET("/characters/{id}")
			Response(StatusOK)
		})
	})

	// Updating a Character by ID
	Method("update_character", func() {
		Description("Updating a Character")
		Payload(func() {
			Field(1, "id", Int, "Character ID")
			Field(2, "name", String, "Character name")
			Field(3, "description", String, "Character description")
			Field(4, "health", Int, "Character health")
			Field(5, "experience", Int, "Character experience")
			Required("id")
		})
		Error("not_found", String, "Character not found")
		HTTP(func() {
			PUT("/characters/{id}")
			Response(StatusOK)
		})
	})

	// Deleting a Character by ID
	Method("delete_character", func() {
		Description("Deleting a Character")
		Payload(func() {
			Field(1, "id", Int, "Character ID")
			Required("id")
		})
		Error("not_found", String, "Character not found")
		HTTP(func() {
			DELETE("/characters/{id}")
			Response(StatusOK)
		})
	})

	// Get character's inventory by ID
	Method("get_inventory", func() {
		Description("Fetching a Character's Inventory")
		Payload(func() {
			Field(1, "character_id", Int, "Character ID")
			Required("character_id")
		})
		Result(Inventory)
		Error("not_found", String, "Character not found")
		HTTP(func() {
			GET("/characters/{character_id}/inventory")
			Response(StatusOK)
		})
	})

	// Adding an item to character's inventory
	Method("add_item_to_inventory", func() {
		Description("Adding an Item to Character's Inventory")
		Payload(func() {
			Field(1, "character_id", Int, "Character ID")
			Field(2, "item_id", Int, "Item ID")
			Required("character_id", "item_id")
		})
		Error("not_found", String, "Character or Item not found")
		Error("not_valid_time", String, "Item not valid")
		HTTP(func() {
			POST("/characters/{character_id}/inventory/items")
			Response(StatusNoContent)
		})
	})

	// Removing an item from character's inventory
	Method("remove_item_from_inventory", func() {
		Description("Removing an Item from Character's Inventory")
		Payload(func() {
			Field(1, "character_id", Int, "Character ID")
			Field(2, "item_id", Int, "Item ID")
			Required("character_id", "item_id")
		})
		Error("not_found", String, "Character or Item not found")
		Error("not_valid_time", String, "Item not valid")
		HTTP(func() {
			DELETE("/characters/{character_id}/inventory/items/{item_id}")
			Response(StatusNoContent)
		})
	})
})
