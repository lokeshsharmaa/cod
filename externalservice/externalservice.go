package externalserviceapi

import (
	externalservice "assignment_crossnokaye/cod/externalservice/gen/externalservice"
	"context"
	"log"
)

// externalservice service example implementation.
// The example methods log the requests and return zero values.
type externalservicesrvc struct {
	logger *log.Logger
}

// NewExternalservice returns the externalservice service implementation.
func NewExternalservice(logger *log.Logger) externalservice.Service {
	return &externalservicesrvc{logger}
}

// Creating a new Character
func (s *externalservicesrvc) CreateCharacter(ctx context.Context, p *externalservice.CreateCharacterPayload) (res *externalservice.Character, err error) {
	res = &externalservice.Character{}
	s.logger.Print("externalservice.create_character")
	return
}

// Fetching a Character
func (s *externalservicesrvc) GetCharacter(ctx context.Context, p *externalservice.GetCharacterPayload) (res *externalservice.Character, err error) {
	res = &externalservice.Character{}
	s.logger.Print("externalservice.get_character")
	return
}

// Updating a Character
func (s *externalservicesrvc) UpdateCharacter(ctx context.Context, p *externalservice.UpdateCharacterPayload) (err error) {
	s.logger.Print("externalservice.update_character")
	return
}

// Deleting a Character
func (s *externalservicesrvc) DeleteCharacter(ctx context.Context, p *externalservice.DeleteCharacterPayload) (err error) {
	s.logger.Print("externalservice.delete_character")
	return
}
