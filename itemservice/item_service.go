package itemserviceapi

import (
	itemservice "assignment_crossnokaye/cod/itemservice/gen/item_service"
	"context"
	"log"
)

// ItemService service example implementation.
// The example methods log the requests and return zero values.
type itemServicesrvc struct {
	logger           *log.Logger
	index            int
	idItemMapper     map[int]*itemservice.Item
	uniqueNameMapper map[string]bool
}

// NewItemService returns the ItemService service implementation.
func NewItemService(logger *log.Logger) itemservice.Service {
	return &itemServicesrvc{logger: logger,
		idItemMapper:     make(map[int]*itemservice.Item),
		uniqueNameMapper: make(map[string]bool)}
}

// Create a new item
func (s *itemServicesrvc) Create(ctx context.Context, p *itemservice.CreatePayload) (res int, err error) {
	s.logger.Print("itemService.Create")
	if _, ok := s.uniqueNameMapper[p.Name]; ok {
		return 0, itemservice.UniqueConstraint("Name not unique")
	}
	// TODO: To store INDEX in DB
	s.index += 1
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

// Retrieve an item by ID
func (s *itemServicesrvc) Get(ctx context.Context, p *itemservice.GetPayload) (res *itemservice.Item, err error) {

	s.logger.Print("itemService.Get")
	res, ok := s.idItemMapper[p.ID]
	if !ok {
		return nil, itemservice.NotFound("Not Found")
	}
	return res, nil
}

// Update an item
func (s *itemServicesrvc) Update(ctx context.Context, p *itemservice.UpdatePayload) (err error) {
	s.logger.Print("itemService.Update")
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
	character := &itemservice.Item{
		ID:          res.ID,
		Name:        p.Name,
		Description: p.Description,
		Damage:      p.Damage,
		Healing:     p.Healing,
		Protection:  p.Protection,
	}
	s.idItemMapper[p.ID] = character
	s.uniqueNameMapper[*p.Name] = true
	return nil
}

// Delete an item by ID
func (s *itemServicesrvc) Delete(ctx context.Context, p *itemservice.DeletePayload) (err error) {
	s.logger.Print("itemService.Delete")
	res, ok := s.idItemMapper[p.ID]
	if !ok {
		return itemservice.NotFound("Not Found")
	}
	delete(s.idItemMapper, p.ID)
	delete(s.uniqueNameMapper, *res.Name)
	return nil
}
