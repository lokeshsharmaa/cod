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
