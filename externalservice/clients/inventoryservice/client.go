package inventoryservice

import (
	inventoryServiceClient "assignment_crossnokaye/cod/inventoryservice/gen/grpc/inventoryservice/client"
	"assignment_crossnokaye/cod/inventoryservice/gen/inventoryservice"
	goa "goa.design/goa/v3/pkg"

	"context"
	"google.golang.org/grpc"
)

type (
	Client interface {
		GetInventory(context.Context, *GetPayload) (res *Inventory, err error)
		AddInventory(context.Context, *AddPayload) (err error)
		DeleteInventory(context.Context, *DeletePayload) (err error)
	}
	AddPayload struct {
		CharacterID int
		ItemID      *int
	}
	DeletePayload struct {
		CharacterID int
		ItemID      int
	}

	GetPayload struct {
		ID int
	}

	Inventory struct {
		CharacterID *int
		Items       []int
	}
	client struct {
		Get    goa.Endpoint
		Add    goa.Endpoint
		Delete goa.Endpoint
	}
)

func New(cc *grpc.ClientConn) Client {
	c := inventoryServiceClient.NewClient(cc, grpc.WaitForReady(true))
	return &client{Get: c.Get(), Add: c.Add(), Delete: c.Delete()}
}

func (c *client) GetInventory(ctx context.Context, g *GetPayload) (res *Inventory, err error) {

	response, err := c.Get(ctx, inventoryservice.GetPayload{
		ID: g.ID,
	})

	if err != nil {
		return nil, err
	}
	res = response.(*Inventory)
	return res, nil
}

func (c *client) AddInventory(ctx context.Context, p *AddPayload) (err error) {
	_, err = c.Add(ctx, &inventoryservice.AddPayload{
		CharacterID: p.CharacterID,
		ItemID:      p.ItemID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *client) DeleteInventory(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.Delete(ctx, &inventoryservice.DeletePayload{
		CharacterID: p.CharacterID,
		ItemID:      p.ItemID,
	})
	if err != nil {
		return err
	}
	return nil
}
