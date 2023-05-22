package characterserviceapi

import (
	characterservice "assignment_crossnokaye/cod/characterservice/gen/characterservice"
	"context"
	"log"
)

// characterservice service example implementation.
// The example methods log the requests and return zero values.
type characterservicesrvc struct {
	logger            *log.Logger
	index             int
	idCharacterMapper map[int]*characterservice.Character
	uniqueNameMapper  map[string]bool
}

// NewCharacterservice returns the characterservice service implementation.
func NewCharacterservice(logger *log.Logger) characterservice.Service {
	return &characterservicesrvc{logger: logger,
		idCharacterMapper: make(map[int]*characterservice.Character),
		uniqueNameMapper:  make(map[string]bool)}
}

// Create a new character.
func (s *characterservicesrvc) Create(ctx context.Context, p *characterservice.CreatePayload) (res int, err error) {
	s.logger.Print("characterservice.Create")
	if _, ok := s.uniqueNameMapper[p.Name]; ok {
		return 0, characterservice.UniqueConstraint("Name not unique")
	}
	// TODO: To store INDEX in DB
	s.index += 1
	currentIndex := s.index
	character := &characterservice.Character{
		ID:          &currentIndex,
		Name:        &p.Name,
		Description: p.Description,
		Health:      &p.Health,
		Experience:  &p.Experience,
	}
	s.uniqueNameMapper[p.Name] = true
	s.idCharacterMapper[currentIndex] = character
	return s.index, nil
}

// Retrieve a character by ID.
func (s *characterservicesrvc) Get(ctx context.Context, p *characterservice.GetPayload) (res *characterservice.Character, err error) {
	s.logger.Print("characterservice.Get")
	res, ok := s.idCharacterMapper[p.ID]
	if !ok {
		return nil, characterservice.NotFound("Not Found")
	}
	return res, nil
}

// Update a character.
func (s *characterservicesrvc) Update(ctx context.Context, p *characterservice.UpdatePayload) (err error) {
	s.logger.Print("characterservice.Update")
	res, ok := s.idCharacterMapper[p.ID]
	if !ok {
		return characterservice.NotFound("Not Found")
	}
	if res.Name != p.Name {
		if _, ok := s.uniqueNameMapper[*p.Name]; ok {
			return characterservice.UniqueConstraint("Name not Unique")
		}
	}
	character := &characterservice.Character{
		ID:          res.ID,
		Name:        p.Name,
		Description: p.Description,
		Health:      p.Health,
		Experience:  p.Experience,
	}
	s.idCharacterMapper[p.ID] = character
	return nil
}

// Delete a character by ID.
func (s *characterservicesrvc) Delete(ctx context.Context, p *characterservice.DeletePayload) (err error) {
	s.logger.Print("characterservice.Delete")
	res, ok := s.idCharacterMapper[p.ID]
	if !ok {
		return characterservice.NotFound("Not Found")
	}
	delete(s.idCharacterMapper, p.ID)
	delete(s.uniqueNameMapper, *res.Name)
	return nil
}
