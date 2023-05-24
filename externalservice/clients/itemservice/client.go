package itemservice

import (
	itemServiceClient "assignment_crossnokaye/cod/itemservice/gen/grpc/item_service/client"
	itemService "assignment_crossnokaye/cod/itemservice/gen/item_service"

	"context"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

type (
	Client interface {
		CreateItem(context.Context, *CreatePayload) (res *int, err error)
		GetItem(context.Context, *GetPayload) (*Item, error)
		UpdateItem(context.Context, *UpdatePayload) error
		DeleteItem(context.Context, *DeletePayload) error
	}
	GetPayload struct {
		ID int
	}
	CreatePayload struct {
		Name        string
		Description *string
		Damage      int
		Healing     *int
		Protection  *int
	}
	Item struct {
		ID          *int
		Name        *string
		Description *string
		Damage      *int
		Healing     *int
		Protection  *int
	}

	DeletePayload struct {
		ID int
	}

	UpdatePayload struct {
		ID          int
		Name        string
		Description *string
		Damage      int
		Healing     *int
		Protection  *int
	}
	client struct {
		Create goa.Endpoint
		Get    goa.Endpoint
		Update goa.Endpoint
		Delete goa.Endpoint
	}
)

func New(cc *grpc.ClientConn) Client {
	c := itemServiceClient.NewClient(cc, grpc.WaitForReady(true))
	return &client{Create: c.Create(), Get: c.Get(), Update: c.Update(), Delete: c.Delete()}
}

func (c *client) CreateItem(ctx context.Context, item *CreatePayload) (*int, error) {

	res, err := c.Create(ctx, &itemService.CreatePayload{
		Name:        item.Name,
		Description: item.Description,
		Damage:      item.Damage,
		Healing:     item.Healing,
		Protection:  item.Protection,
	})
	if err != nil {
		return nil, err
	}
	response := res.(int)
	return &response, nil
}
func (c *client) GetItem(ctx context.Context, payload *GetPayload) (*Item, error) {
	res, err := c.Get(ctx, &itemService.GetPayload{
		ID: payload.ID,
	})
	if err != nil {
		return nil, err
	}
	item := res.(*itemService.Item)

	response := &Item{
		ID:          item.ID,
		Name:        item.Name,
		Description: item.Description,
		Damage:      item.Damage,
		Healing:     item.Healing,
		Protection:  item.Protection,
	}
	return response, nil
}

func (c *client) UpdateItem(ctx context.Context, payload *UpdatePayload) error {
	_, err := c.Update(ctx, &itemService.UpdatePayload{
		ID:          payload.ID,
		Name:        &payload.Name,
		Description: payload.Description,
		Damage:      &payload.Damage,
		Healing:     payload.Healing,
		Protection:  payload.Protection,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *client) DeleteItem(ctx context.Context, payload *DeletePayload) error {
	_, err := c.Delete(ctx, &itemService.DeletePayload{
		ID: payload.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
