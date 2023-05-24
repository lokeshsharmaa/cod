package inventoryserviceapi

import (
	"assignment_crossnokaye/cod/inventoryservice/gen/inventoryservice"
	"context"
	"log"
)

type inventoryservicesrvc struct {
	logger               *log.Logger
	characterItemsMapper map[int]map[int]bool
}

func NewInventoryservice(logger *log.Logger) inventoryservice.Service {
	return &inventoryservicesrvc{logger, make(map[int]map[int]bool)}
}

func (s *inventoryservicesrvc) Get(ctx context.Context, p *inventoryservice.GetPayload) (res *inventoryservice.Inventory, err error) {
	items, found := s.characterItemsMapper[p.ID]
	keys := make([]int, 0, len(items))

	if !found {
		res = &inventoryservice.Inventory{CharacterID: &p.ID, Items: keys}
		return res, nil
	}

	for k := range items {
		keys = append(keys, k)
	}
	res = &inventoryservice.Inventory{CharacterID: &p.ID, Items: keys}
	return res, nil
}

func (s *inventoryservicesrvc) Add(ctx context.Context, p *inventoryservice.AddPayload) (err error) {
	_, found := s.characterItemsMapper[p.CharacterID]
	if !found {
		s.characterItemsMapper[p.CharacterID] = make(map[int]bool)
	}
	s.characterItemsMapper[p.CharacterID][*p.ItemID] = true
	return nil
}

func (s *inventoryservicesrvc) Delete(ctx context.Context, p *inventoryservice.DeletePayload) (err error) {
	items, found := s.characterItemsMapper[p.CharacterID]
	if !found {
		return inventoryservice.NotFound("character not found")
	}
	if _, exists := items[p.ItemID]; !exists {
		return inventoryservice.NotFound("item not found")
	}
	delete(s.characterItemsMapper[p.CharacterID], p.ItemID)
	return nil
}
