package characterserviceapi

import (
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"testing"

	"assignment_crossnokaye/cod/characterservice/gen/characterservice"
	"context"
)

func newTestService(t *testing.T) characterservice.Service {
	t.Helper()
	logger := log.New(io.Discard, "", 0)
	return NewCharacterservice(logger)
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
	payload := &characterservice.CreatePayload{
		Name:        "Test",
		Description: StringPtr("Dummy"),
		Health:      100,
		Experience:  0,
	}

	res, err := service.Create(ctx, payload)

	if err != nil {
		t.Errorf("Create failed: %v", err)
	}
	if res <= 0 {
		t.Errorf("Invalid character ID: %d", res)
	}
}

func TestGet(t *testing.T) {
	service := newTestService(t)

	ctx := context.Background()
	createPayload := &characterservice.CreatePayload{
		Name:        "Test",
		Description: StringPtr("Dummy"),
		Health:      100,
		Experience:  0,
	}
	characterID, err := service.Create(ctx, createPayload)
	if err != nil {
		t.Fatalf("Failed to create character: %v", err)
	}

	getPayload := &characterservice.GetPayload{
		ID: characterID,
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
	createPayload := &characterservice.CreatePayload{
		Name:        "Test",
		Description: StringPtr("Dummy"),
		Health:      100,
		Experience:  0,
	}
	characterID, err := service.Create(ctx, createPayload)
	if err != nil {
		t.Fatalf("Failed to create character: %v", err)
	}

	updatePayload := &characterservice.UpdatePayload{
		ID:          characterID,
		Name:        StringPtr("Dummy"),
		Description: StringPtr("test"),
		Health:      IntPtr(150),
		Experience:  IntPtr(50),
	}

	err = service.Update(ctx, updatePayload)

	if err != nil {
		t.Errorf("Update failed: %v", err)
	}

	getPayload := &characterservice.GetPayload{
		ID: characterID,
	}

	res, _ := service.Get(ctx, getPayload)

	assert.Equal(t, *res.Name, "Dummy")
	assert.Equal(t, *res.Description, "test")
	assert.Equal(t, *res.Health, 150)
	assert.Equal(t, *res.Experience, 50)
}

func TestDelete(t *testing.T) {
	service := newTestService(t)

	ctx := context.Background()
	createPayload := &characterservice.CreatePayload{
		Name:        "Test",
		Description: StringPtr("Dummy"),
		Health:      100,
		Experience:  0,
	}
	characterID, err := service.Create(ctx, createPayload)
	if err != nil {
		t.Fatalf("Failed to create character: %v", err)
	}

	deletePayload := &characterservice.DeletePayload{
		ID: characterID,
	}

	err = service.Delete(ctx, deletePayload)

	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}
	getPayload := &characterservice.GetPayload{
		ID: characterID,
	}
	res, err := service.Get(ctx, getPayload)
	assert.Empty(t, res)
}

func TestUniqueConstraints(t *testing.T) {
	service := newTestService(t)

	ctx := context.Background()
	createPayload := &characterservice.CreatePayload{
		Name:        "Test",
		Description: StringPtr("Dummy"),
		Health:      100,
		Experience:  0,
	}
	characterID, _ := service.Create(ctx, createPayload)
	_, err := service.Create(ctx, createPayload)

	assert.Equal(t, err.Error(), "Name not unique")

	deletePayload := &characterservice.DeletePayload{
		ID: characterID,
	}
	service.Delete(ctx, deletePayload)
	characterID, _ = service.Create(ctx, createPayload)
	assert.Equal(t, characterID, 2)

}
