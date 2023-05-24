// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC client types
//
// Command:
// $ goa gen assignment_crossnokaye/cod/itemservice/design

package client

import (
	item_servicepb "assignment_crossnokaye/cod/itemservice/gen/grpc/item_service/pb"
	itemservice "assignment_crossnokaye/cod/itemservice/gen/item_service"
)

// NewProtoCreateRequest builds the gRPC request type from the payload of the
// "Create" endpoint of the "ItemService" service.
func NewProtoCreateRequest(payload *itemservice.CreatePayload) *item_servicepb.CreateRequest {
	message := &item_servicepb.CreateRequest{
		Name:        payload.Name,
		Description: payload.Description,
		Damage:      int32(payload.Damage),
	}
	if payload.Healing != nil {
		healing := int32(*payload.Healing)
		message.Healing = &healing
	}
	if payload.Protection != nil {
		protection := int32(*payload.Protection)
		message.Protection = &protection
	}
	return message
}

// NewCreateResult builds the result type of the "Create" endpoint of the
// "ItemService" service from the gRPC response type.
func NewCreateResult(message *item_servicepb.CreateResponse) int {
	result := int(message.Field)
	return result
}

// NewProtoGetRequest builds the gRPC request type from the payload of the
// "Get" endpoint of the "ItemService" service.
func NewProtoGetRequest(payload *itemservice.GetPayload) *item_servicepb.GetRequest {
	message := &item_servicepb.GetRequest{
		Id: int32(payload.ID),
	}
	return message
}

// NewGetResult builds the result type of the "Get" endpoint of the
// "ItemService" service from the gRPC response type.
func NewGetResult(message *item_servicepb.GetResponse) *itemservice.Item {
	result := &itemservice.Item{
		Name:        message.Name,
		Description: message.Description,
	}
	if message.Id != nil {
		id := int(*message.Id)
		result.ID = &id
	}
	if message.Damage != nil {
		damage := int(*message.Damage)
		result.Damage = &damage
	}
	if message.Healing != nil {
		healing := int(*message.Healing)
		result.Healing = &healing
	}
	if message.Protection != nil {
		protection := int(*message.Protection)
		result.Protection = &protection
	}
	return result
}

// NewProtoUpdateRequest builds the gRPC request type from the payload of the
// "Update" endpoint of the "ItemService" service.
func NewProtoUpdateRequest(payload *itemservice.UpdatePayload) *item_servicepb.UpdateRequest {
	message := &item_servicepb.UpdateRequest{
		Id:          int32(payload.ID),
		Name:        payload.Name,
		Description: payload.Description,
	}
	if payload.Damage != nil {
		damage := int32(*payload.Damage)
		message.Damage = &damage
	}
	if payload.Healing != nil {
		healing := int32(*payload.Healing)
		message.Healing = &healing
	}
	if payload.Protection != nil {
		protection := int32(*payload.Protection)
		message.Protection = &protection
	}
	return message
}

// NewProtoDeleteRequest builds the gRPC request type from the payload of the
// "Delete" endpoint of the "ItemService" service.
func NewProtoDeleteRequest(payload *itemservice.DeletePayload) *item_servicepb.DeleteRequest {
	message := &item_servicepb.DeleteRequest{
		Id: int32(payload.ID),
	}
	return message
}