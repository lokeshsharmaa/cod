package externalserviceapi

import (
	"assignment_crossnokaye/cod/externalservice/clients/characterservice"
	characterServiceMock "assignment_crossnokaye/cod/externalservice/clients/characterservice/mocks"
	"assignment_crossnokaye/cod/externalservice/clients/inventoryservice"
	inventoryServiceMock "assignment_crossnokaye/cod/externalservice/clients/inventoryservice/mocks"
	"assignment_crossnokaye/cod/externalservice/gen/externalservice"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCreateCharacter(t *testing.T) {
	logger := log.New(nil, "", 0)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCharacterServiceClient := characterServiceMock.NewMockClient(ctrl)
	mockInventoryServiceClient := inventoryServiceMock.NewMockClient(ctrl)
	service := NewExternalservice(logger, mockCharacterServiceClient, mockInventoryServiceClient, nil)

	ctx := context.Background()
	name := "test"
	description := "Dummy"
	rets := 1

	payload := &externalservice.CreateCharacterPayload{
		Name:        name,
		Description: &description,
		Health:      10,
		Experience:  10,
	}

	expectedPayload := &characterservice.CreatePayload{
		Name:        name,
		Description: &description,
		Health:      10,
		Experience:  10,
	}

	mockCharacterServiceClient.EXPECT().CreateCharacter(ctx, expectedPayload).Return(&rets, nil)

	res, err := service.CreateCharacter(ctx, payload)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 1, *res.ID)
	assert.Equal(t, name, *res.Name)
	assert.Equal(t, description, *res.Description)
	assert.Equal(t, 10, *res.Health)
	assert.Equal(t, 10, *res.Experience)

}

func TestGetCharacter(t *testing.T) {
	logger := log.New(nil, "", 0)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCharacterServiceClient := characterServiceMock.NewMockClient(ctrl)
	mockInventoryServiceClient := inventoryServiceMock.NewMockClient(ctrl)
	service := NewExternalservice(logger, mockCharacterServiceClient, mockInventoryServiceClient, nil)

	ctx := context.Background()

	id := 1
	name := "test"
	description := "dummy"
	health := 100
	experience := 0

	payload := &externalservice.GetCharacterPayload{
		ID: id,
	}

	expectedPayload := &characterservice.GetPayload{
		ID: id,
	}

	character := &characterservice.Character{
		ID:          &id,
		Name:        &name,
		Description: &description,
		Health:      &health,
		Experience:  &experience,
	}

	mockCharacterServiceClient.EXPECT().GetCharacter(ctx, expectedPayload).Return(character, nil)

	res, err := service.GetCharacter(ctx, payload)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, id, *res.ID)
	assert.Equal(t, name, *res.Name)
	assert.Equal(t, description, *res.Description)
	assert.Equal(t, health, *res.Health)
	assert.Equal(t, experience, *res.Experience)

}

func TestGetInventory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCharacterServiceClient := characterServiceMock.NewMockClient(ctrl)
	mockInventoryServiceClient := inventoryServiceMock.NewMockClient(ctrl)

	logger := log.New(nil, "", 0)
	service := NewExternalservice(logger, mockCharacterServiceClient, mockInventoryServiceClient, nil)

	ctx := context.Background()
	characterID := 1
	payload := &externalservice.GetInventoryPayload{
		CharacterID: characterID,
	}

	expectedPayload := &inventoryservice.GetPayload{
		ID: characterID,
	}

	expectedResponse := &inventoryservice.Inventory{
		CharacterID: &characterID,
		Items:       []int{1, 2},
	}

	mockInventoryServiceClient.EXPECT().GetInventory(ctx, expectedPayload).Return(expectedResponse, nil)

	res, err := service.GetInventory(ctx, payload)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, characterID, *res.CharacterID)
	assert.Equal(t, []int{1, 2}, res.Items)
}

func TestAddItemToInventory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCharacterServiceClient := characterServiceMock.NewMockClient(ctrl)
	mockInventoryServiceClient := inventoryServiceMock.NewMockClient(ctrl)

	logger := log.New(nil, "", 0)
	service := NewExternalservice(logger, mockCharacterServiceClient, mockInventoryServiceClient, nil)

	ctx := context.Background()
	id := 1
	payload := &externalservice.AddItemToInventoryPayload{
		CharacterID: id,
		ItemID:      id,
	}

	expectedPayload := &inventoryservice.AddPayload{
		CharacterID: id,
		ItemID:      &id,
	}

	mockInventoryServiceClient.EXPECT().AddInventory(ctx, expectedPayload).Return(nil)

	err := service.AddItemToInventory(ctx, payload)

	assert.NoError(t, err)
}

func TestRemoveItemFromInventory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCharacterServiceClient := characterServiceMock.NewMockClient(ctrl)
	mockInventoryServiceClient := inventoryServiceMock.NewMockClient(ctrl)

	logger := log.New(nil, "", 0)
	service := NewExternalservice(logger, mockCharacterServiceClient, mockInventoryServiceClient, nil)

	ctx := context.Background()
	payload := &externalservice.RemoveItemFromInventoryPayload{
		CharacterID: 1,
		ItemID:      1,
	}
	expectedPayload := &inventoryservice.DeletePayload{
		CharacterID: 1,
		ItemID:      1,
	}

	mockInventoryServiceClient.EXPECT().DeleteInventory(ctx, expectedPayload).Return(nil)

	err := service.RemoveItemFromInventory(ctx, payload)

	assert.NoError(t, err)
}
