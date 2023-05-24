package characterservice

import (
	"assignment_crossnokaye/cod/characterservice/gen/characterservice"
	characterserviceClient "assignment_crossnokaye/cod/characterservice/gen/grpc/characterservice/client"
	"context"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

type (
	Client interface {
		CreateCharacter(context.Context, *CreatePayload) (res *int, err error)
		GetCharacter(context.Context, *GetPayload) (*Character, error)
		UpdateCharacter(context.Context, *UpdatePayload) error
		DeleteCharacter(context.Context, *DeletePayload) error
	}
	GetPayload struct {
		ID int
	}
	CreatePayload struct {
		Name        string
		Description *string
		Health      int
		Experience  int
	}
	Character struct {
		ID          *int
		Name        *string
		Description *string
		Health      *int
		Experience  *int
	}

	DeletePayload struct {
		ID int
	}

	UpdatePayload struct {
		ID          int
		Name        *string
		Description *string
		Health      *int
		Experience  *int
	}
	client struct {
		Create goa.Endpoint
		Get    goa.Endpoint
		Update goa.Endpoint
		Delete goa.Endpoint
	}
)

func New(cc *grpc.ClientConn) Client {
	c := characterserviceClient.NewClient(cc, grpc.WaitForReady(true))
	return &client{Create: c.Create(), Get: c.Get(), Update: c.Update(), Delete: c.Delete()}
}

func (c *client) CreateCharacter(ctx context.Context, character *CreatePayload) (*int, error) {

	res, err := c.Create(ctx, &characterservice.CreatePayload{
		Name:        character.Name,
		Description: character.Description,
		Health:      character.Health,
		Experience:  character.Experience})
	if err != nil {
		return nil, err
	}
	response := res.(int)
	return &response, nil
}
func (c *client) GetCharacter(ctx context.Context, payload *GetPayload) (*Character, error) {
	res, err := c.Get(ctx, &characterservice.GetPayload{
		ID: payload.ID,
	})
	if err != nil {
		return nil, err
	}
	character := res.(*Character)
	return character, nil
}

func (c *client) UpdateCharacter(ctx context.Context, payload *UpdatePayload) error {
	_, err := c.Update(ctx, &characterservice.UpdatePayload{
		ID:          payload.ID,
		Name:        payload.Name,
		Description: payload.Description,
		Health:      payload.Health,
		Experience:  payload.Experience,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *client) DeleteCharacter(ctx context.Context, payload *DeletePayload) error {
	_, err := c.Delete(ctx, &characterservice.DeletePayload{
		ID: payload.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
