package itemserviceapi

import (
	itemservice "assignment_crossnokaye/cod/itemservice/gen/item_service"
	"context"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"testing"
)

func newTestService(t *testing.T) itemservice.Service {
	logger := log.New(io.Discard, "", 0)
	return NewItemService(logger)
}

func StringPtr(s string) *string {
	return &s
}

func IntPtr(s int) *int {
	return &s
}

func TestCreate(t *testing.T) {
	service := newTestService(t)

	ctx := context.Background()
	payload := createPayload()

	res, err := service.Create(ctx, payload)

	if err != nil {
		t.Errorf("Create failed: %v", err)
	}
	if res <= 0 {
		t.Errorf("Invalid character ID: %d", res)
	}
}

func createPayload() *itemservice.CreatePayload {
	return &itemservice.CreatePayload{
		Name:        "test",
		Description: StringPtr("dummy"),
		Damage:      12,
		Healing:     IntPtr(10),
		Protection:  IntPtr(10),
	}
}

func TestGet(t *testing.T) {
	service := newTestService(t)

	ctx := context.Background()
	createPayload := createPayload()
	itemID, err := service.Create(ctx, createPayload)
	if err != nil {
		t.Fatalf("Failed to create character: %v", err)
	}

	getPayload := &itemservice.GetPayload{
		ID: itemID,
	}

	res, err := service.Get(ctx, getPayload)

	if err != nil {
		t.Errorf("Get failed: %v", err)
	}
	if res == nil {
		t.Errorf("Character not found")
	}
}

func TestUpdate(t *testing.T) {
	service := newTestService(t)

	ctx := context.Background()
	createPayload := createPayload()
	itemID, err := service.Create(ctx, createPayload)
	if err != nil {
		t.Fatalf("Failed to create character: %v", err)
	}

	updatePayload := &itemservice.UpdatePayload{
		ID:          itemID,
		Name:        StringPtr("test"),
		Description: StringPtr("dummy"),
		Damage:      IntPtr(13),
		Healing:     IntPtr(16),
		Protection:  IntPtr(16),
	}

	err = service.Update(ctx, updatePayload)

	if err != nil {
		t.Errorf("Update failed: %v", err)
	}

	getPayload := &itemservice.GetPayload{
		ID: itemID,
	}

	res, _ := service.Get(ctx, getPayload)

	assert.Equal(t, *res.Name, "test")
	assert.Equal(t, *res.Description, "dummy")
	assert.Equal(t, *res.Damage, *updatePayload.Damage)
	assert.Equal(t, *res.Healing, *updatePayload.Healing)
	assert.Equal(t, *res.Protection, *updatePayload.Protection)
}

func TestDelete(t *testing.T) {
	service := newTestService(t)

	ctx := context.Background()
	createPayload := createPayload()
	itemID, err := service.Create(ctx, createPayload)
	if err != nil {
		t.Fatalf("Failed to create character: %v", err)
	}

	deletePayload := &itemservice.DeletePayload{
		ID: itemID,
	}

	err = service.Delete(ctx, deletePayload)

	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	getPayload := &itemservice.GetPayload{
		ID: itemID,
	}
	res, err := service.Get(ctx, getPayload)
	assert.Empty(t, res)
}

func TestUniqueConstraints(t *testing.T) {
	service := newTestService(t)

	ctx := context.Background()
	createPayload := createPayload()
	itemID, _ := service.Create(ctx, createPayload)
	_, err := service.Create(ctx, createPayload)

	assert.Equal(t, err, itemservice.UniqueConstraint("Name not unique"))

	deletePayload := &itemservice.DeletePayload{
		ID: itemID,
	}
	service.Delete(ctx, deletePayload)
	itemID, _ = service.Create(ctx, createPayload)
	assert.Equal(t, itemID, 2)

}
