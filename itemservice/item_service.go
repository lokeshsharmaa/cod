package itemserviceapi

import (
	itemservice "assignment_crossnokaye/cod/itemservice/gen/item_service"
	"context"
	"log"
)

type itemServicesrvc struct {
	logger           *log.Logger
	index            int
	idItemMapper     map[int]*itemservice.Item
	uniqueNameMapper map[string]bool
}

func NewItemService(logger *log.Logger) itemservice.Service {
	return &itemServicesrvc{
		logger:           logger,
		index:            0,
		idItemMapper:     make(map[int]*itemservice.Item),
		uniqueNameMapper: make(map[string]bool),
	}
}

func (s *itemServicesrvc) Create(ctx context.Context, p *itemservice.CreatePayload) (res int, err error) {
	if _, ok := s.uniqueNameMapper[p.Name]; ok {
		return 0, itemservice.UniqueConstraint("Name not unique")
	}
	s.index++
	currentIndex := s.index
	item := &itemservice.Item{
		ID:          &currentIndex,
		Name:        &p.Name,
		Description: p.Description,
		Damage:      &p.Damage,
		Healing:     p.Healing,
		Protection:  p.Protection,
	}
	s.uniqueNameMapper[p.Name] = true
	s.idItemMapper[currentIndex] = item
	return s.index, nil
}

func (s *itemServicesrvc) Get(ctx context.Context, p *itemservice.GetPayload) (res *itemservice.Item, err error) {
	res, ok := s.idItemMapper[p.ID]
	if !ok {
		return nil, itemservice.NotFound("Not Found")
	}
	return res, nil
}

func (s *itemServicesrvc) Update(ctx context.Context, p *itemservice.UpdatePayload) (err error) {
	res, ok := s.idItemMapper[p.ID]
	if !ok {
		return itemservice.NotFound("Not Found")
	}
	if res.Name != p.Name {
		if _, ok := s.uniqueNameMapper[*p.Name]; ok {
			return itemservice.UniqueConstraint("Name not Unique")
		}
		delete(s.uniqueNameMapper, *res.Name)
	}
	item := &itemservice.Item{
		ID:          res.ID,
		Name:        p.Name,
		Description: p.Description,
		Damage:      p.Damage,
		Healing:     p.Healing,
		Protection:  p.Protection,
	}
	s.idItemMapper[p.ID] = item
	s.uniqueNameMapper[*p.Name] = true
	return nil
}

func (s *itemServicesrvc) Delete(ctx context.Context, p *itemservice.DeletePayload) (err error) {
	res, ok := s.idItemMapper[p.ID]
	if !ok {
		return itemservice.NotFound("Not Found")
	}
	delete(s.idItemMapper, p.ID)
	delete(s.uniqueNameMapper, *res.Name)
	return nil
}
