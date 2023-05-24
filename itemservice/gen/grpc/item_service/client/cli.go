// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ItemService gRPC client CLI support package
//
// Command:
// $ goa gen assignment_crossnokaye/cod/itemservice/design

package client

import (
	item_servicepb "assignment_crossnokaye/cod/itemservice/gen/grpc/item_service/pb"
	itemservice "assignment_crossnokaye/cod/itemservice/gen/item_service"
	"encoding/json"
	"fmt"
)

// BuildCreatePayload builds the payload for the ItemService Create endpoint
// from CLI flags.
func BuildCreatePayload(itemServiceCreateMessage string) (*itemservice.CreatePayload, error) {
	var err error
	var message item_servicepb.CreateRequest
	{
		if itemServiceCreateMessage != "" {
			err = json.Unmarshal([]byte(itemServiceCreateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"damage\": 3205420170586316006,\n      \"description\": \"Et in alias est expedita quos eligendi.\",\n      \"healing\": 1329053061876795014,\n      \"name\": \"Est velit culpa aspernatur dolorem.\",\n      \"protection\": 120370757675553720\n   }'")
			}
		}
	}
	v := &itemservice.CreatePayload{
		Name:        message.Name,
		Description: message.Description,
		Damage:      int(message.Damage),
	}
	if message.Healing != nil {
		healing := int(*message.Healing)
		v.Healing = &healing
	}
	if message.Protection != nil {
		protection := int(*message.Protection)
		v.Protection = &protection
	}

	return v, nil
}

// BuildGetPayload builds the payload for the ItemService Get endpoint from CLI
// flags.
func BuildGetPayload(itemServiceGetMessage string) (*itemservice.GetPayload, error) {
	var err error
	var message item_servicepb.GetRequest
	{
		if itemServiceGetMessage != "" {
			err = json.Unmarshal([]byte(itemServiceGetMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 2011463225270065069\n   }'")
			}
		}
	}
	v := &itemservice.GetPayload{
		ID: int(message.Id),
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the ItemService Update endpoint
// from CLI flags.
func BuildUpdatePayload(itemServiceUpdateMessage string) (*itemservice.UpdatePayload, error) {
	var err error
	var message item_servicepb.UpdateRequest
	{
		if itemServiceUpdateMessage != "" {
			err = json.Unmarshal([]byte(itemServiceUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"damage\": 7849875426662081395,\n      \"description\": \"Aut facilis alias eos.\",\n      \"healing\": 1175548812336538195,\n      \"id\": 1771631765212308878,\n      \"name\": \"Commodi facilis maiores.\",\n      \"protection\": 1137017544043901924\n   }'")
			}
		}
	}
	v := &itemservice.UpdatePayload{
		ID:          int(message.Id),
		Name:        message.Name,
		Description: message.Description,
	}
	if message.Damage != nil {
		damage := int(*message.Damage)
		v.Damage = &damage
	}
	if message.Healing != nil {
		healing := int(*message.Healing)
		v.Healing = &healing
	}
	if message.Protection != nil {
		protection := int(*message.Protection)
		v.Protection = &protection
	}

	return v, nil
}

// BuildDeletePayload builds the payload for the ItemService Delete endpoint
// from CLI flags.
func BuildDeletePayload(itemServiceDeleteMessage string) (*itemservice.DeletePayload, error) {
	var err error
	var message item_servicepb.DeleteRequest
	{
		if itemServiceDeleteMessage != "" {
			err = json.Unmarshal([]byte(itemServiceDeleteMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 6263443980830511219\n   }'")
			}
		}
	}
	v := &itemservice.DeletePayload{
		ID: int(message.Id),
	}

	return v, nil
}
