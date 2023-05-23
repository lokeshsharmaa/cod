package inventoryserviceapi

import (
	"assignment_crossnokaye/cod/inventoryservice/gen/inventoryservice"
	"context"
	"log"
)

type inventoryservicesrvc struct {
	logger         *log.Logger
	characterItems map[int]map[int]bool
}

func NewInventoryservice(logger *log.Logger) inventoryservice.Service {
	return &inventoryservicesrvc{logger, make(map[int]map[int]bool)}
}

func (s *inventoryservicesrvc) Get(ctx context.Context, p *inventoryservice.GetPayload) (res *inventoryservice.Inventory, err error) {
	keys := make([]int, 0, len(s.characterItems[p.ID]))
	for k := range s.characterItems[p.ID] {
		keys = append(keys, k)
	}
	res = &inventoryservice.Inventory{CharacterID: &p.ID, Items: keys}
	s.logger.Print("inventoryservice.Get")
	return res, nil
}

func (s *inventoryservicesrvc) Add(ctx context.Context, p *inventoryservice.AddPayload) (err error) {
	s.logger.Print("inventoryservice.add")
	_, ok := s.characterItems[p.CharacterID]
	if !ok {
		s.characterItems[p.CharacterID] = make(map[int]bool)
	}
	s.characterItems[p.CharacterID][*p.ItemID] = true
	return nil
}

func (s *inventoryservicesrvc) Delete(ctx context.Context, p *inventoryservice.DeletePayload) (err error) {
	delete(s.characterItems[p.CharacterID], p.ItemID)
	return nil
}
