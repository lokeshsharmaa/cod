package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("characterservice", func() {
	Title("Character Service")
	Description("CRUD operations on characters")
	Server("characterservice", func() {
		Host("localhost", func() {
			URI("grpc://localhost:8081")
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

var _ = Service("characterservice", func() {
	Description("Handles CRUD operations for characters.")
	Method("Create", func() {
		Description("Create a new character.")
		Payload(func() {
			Field(1, "name", String, "Character name")
			Field(2, "description", String, "Character description")
			Field(3, "health", Int, "Character health")
			Field(4, "experience", Int, "Character experience")
			Required("name", "health", "experience")
		})
		Result(Int, "id", "ID of character")
		Error("unique_constraint", String, "Name not unique")
		GRPC(func() {
			Response(CodeOK)
		})
	})

	// Define the "Get" method for retrieving a character by ID
	Method("Get", func() {
		Description("Retrieve a character by ID.")
		Payload(func() {
			Field(1, "id", Int, "Character ID")
			Required("id")
		})
		Result(Character, "character", "Retrieved character")
		Error("not_found", String, "Character not found")
		Error("unique_constraint", String, "Name not unique")
		GRPC(func() {
			Response(CodeOK)
		})
	})

	// Define the "Update" method for updating a character
	Method("Update", func() {
		Description("Update a character.")
		Payload(func() {
			Field(1, "id", Int, "Character ID")
			Field(2, "name", String, "Character name")
			Field(3, "description", String, "Character description")
			Field(4, "health", Int, "Character health")
			Field(5, "experience", Int, "Character experience")
			Required("id")
		})
		Error("not_found", String, "Character not found")
		Error("unique_constraint", String, "Name not unique")
		GRPC(func() {
			Response(CodeOK)
		})
	})

	// Define the "Delete" method for deleting a character
	Method("Delete", func() {
		Description("Delete a character by ID.")
		Payload(func() {
			Field(1, "id", Int, "Character ID")
			Required("id")
		})
		Error("not_found", String, "Character not found")

		GRPC(func() {
			Response(CodeOK)
		})
	})

})
