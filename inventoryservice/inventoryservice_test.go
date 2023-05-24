package inventoryserviceapi

import (
	"assignment_crossnokaye/cod/inventoryservice/gen/inventoryservice"
	"context"
	"io"
	"log"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestService(t *testing.T) inventoryservice.Service {
	logger := log.New(io.Discard, "", 0)
	return NewInventoryservice(logger)
}

func TestAdd(t *testing.T) {

	inventoryService := newTestService(t)

	characterID := 1
	itemID := 100

	err := inventoryService.Add(context.Background(), &inventoryservice.AddPayload{
		CharacterID: characterID,
		ItemID:      &itemID,
	})
	assert.NoError(t, err, "Failed to add item to inventory")
	getPayload := &inventoryservice.GetPayload{
		ID: characterID,
	}
	response, err := inventoryService.Get(context.Background(), getPayload)
	assert.NoError(t, err, "Failed to get inventory")

	assert.NotNil(t, response)
	assert.Equal(t, characterID, *response.CharacterID)
	assert.Contains(t, response.Items, itemID)

	err = inventoryService.Delete(context.Background(), &inventoryservice.DeletePayload{
		CharacterID: characterID,
		ItemID:      itemID,
	})
	assert.NoError(t, err, "Failed to delete item from inventory")

	// Get the inventory again
	response, err = inventoryService.Get(context.Background(), getPayload)
	assert.NoError(t, err, "Failed to get inventory")

	// Ensure the item is no longer in the inventory
	assert.NotNil(t, response)
	assert.Equal(t, characterID, *response.CharacterID)
	assert.NotContains(t, response.Items, itemID)
}

func TestDelete(t *testing.T) {

	inventoryService := newTestService(t)

	characterID := 1
	itemID := 100

	err := inventoryService.Add(context.Background(), &inventoryservice.AddPayload{
		CharacterID: characterID,
		ItemID:      &itemID,
	})
	getPayload := &inventoryservice.GetPayload{
		ID: characterID,
	}
	response, err := inventoryService.Get(context.Background(), getPayload)
	assert.Equal(t, characterID, *response.CharacterID)

	err = inventoryService.Delete(context.Background(), &inventoryservice.DeletePayload{
		CharacterID: characterID,
		ItemID:      itemID,
	})
	assert.NoError(t, err, "Failed to delete item from inventory")
	response, err = inventoryService.Get(context.Background(), getPayload)

	assert.Equal(t, characterID, *response.CharacterID)
	assert.NotContains(t, response.Items, itemID)
	assert.Empty(t, response.Items)
}

func TestGet(t *testing.T) {

	inventoryService := newTestService(t)

	characterID := 1
	itemID1 := 100
	itemID2 := 101

	getPayload := &inventoryservice.GetPayload{
		ID: characterID,
	}

	response, _ := inventoryService.Get(context.Background(), getPayload)

	assert.Nil(t, response.Items)

	inventoryService.Add(context.Background(), &inventoryservice.AddPayload{
		CharacterID: characterID,
		ItemID:      &itemID1,
	})
	inventoryService.Add(context.Background(), &inventoryservice.AddPayload{
		CharacterID: characterID,
		ItemID:      &itemID2,
	})

	response, _ = inventoryService.Get(context.Background(), getPayload)
	assert.Equal(t, characterID, *response.CharacterID)
	expected := []int{itemID1, itemID2}
	sort.Ints(response.Items)
	assert.Equal(t, expected, response.Items)

}

func TestDeleteWithError(t *testing.T) {
	inventoryService := newTestService(t)

	err := inventoryService.Delete(context.Background(), &inventoryservice.DeletePayload{
		ItemID:      1,
		CharacterID: 1,
	})

	assert.Equal(t, err, inventoryservice.NotFound("character not found"))

	characterID := 1
	itemID1 := 100

	inventoryService.Add(context.Background(), &inventoryservice.AddPayload{
		CharacterID: characterID,
		ItemID:      &itemID1,
	})

	err = inventoryService.Delete(context.Background(), &inventoryservice.DeletePayload{
		ItemID:      1,
		CharacterID: 1,
	})

	assert.Equal(t, err, inventoryservice.NotFound("item not found"))
}
