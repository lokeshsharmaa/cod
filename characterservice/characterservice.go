package characterserviceapi

import (
	characterservice "assignment_crossnokaye/cod/characterservice/gen/characterservice"
	"context"
	"fmt"
	"log"
)

type characterservicesrvc struct {
	logger            *log.Logger
	index             int
	idCharacterMapper map[int]*characterservice.Character
	uniqueNameMapper  map[string]bool
}

func NewCharacterservice(logger *log.Logger) characterservice.Service {
	return &characterservicesrvc{
		logger:            logger,
		idCharacterMapper: make(map[int]*characterservice.Character),
		uniqueNameMapper:  make(map[string]bool),
	}
}

func (s *characterservicesrvc) Create(ctx context.Context, p *characterservice.CreatePayload) (res int, err error) {
	if !s.isNameUnique(p.Name) {
		return 0, characterservice.UniqueConstraint(fmt.Sprintf("Name '%s' is not unique", p.Name))
	}

	s.index++
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

func (s *characterservicesrvc) Get(ctx context.Context, p *characterservice.GetPayload) (res *characterservice.Character, err error) {
	character, ok := s.idCharacterMapper[p.ID]
	if !ok {
		return nil, characterservice.NotFound(fmt.Sprintf("Character with ID '%d' not found", p.ID))
	}
	return character, nil
}

func (s *characterservicesrvc) Update(ctx context.Context, p *characterservice.UpdatePayload) (err error) {
	character, ok := s.idCharacterMapper[p.ID]
	if !ok {
		return characterservice.NotFound(fmt.Sprintf("Character with ID '%d' not found", p.ID))
	}

	if *character.Name != *p.Name {
		if !s.isNameUnique(*p.Name) {
			return characterservice.UniqueConstraint(fmt.Sprintf("Name '%s' is not unique", *p.Name))
		}
		delete(s.uniqueNameMapper, *character.Name)
	}
	character.Name = p.Name
	character.Description = p.Description
	character.Health = p.Health
	character.Experience = p.Experience

	return nil
}

func (s *characterservicesrvc) Delete(ctx context.Context, p *characterservice.DeletePayload) (err error) {
	character, ok := s.idCharacterMapper[p.ID]
	if !ok {
		return characterservice.NotFound(fmt.Sprintf("Character with ID '%d' not found", p.ID))
	}

	delete(s.idCharacterMapper, p.ID)
	delete(s.uniqueNameMapper, *character.Name)
	return nil
}

func (s *characterservicesrvc) isNameUnique(name string) bool {
	_, res := s.uniqueNameMapper[name]
	return !res
}
