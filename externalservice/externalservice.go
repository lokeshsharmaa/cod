package externalserviceapi

import (
	"assignment_crossnokaye/cod/externalservice/clients/characterservice"
	externalservice "assignment_crossnokaye/cod/externalservice/gen/externalservice"
	"context"
	"log"
)

// externalservice service example implementation.
// The example methods log the requests and return zero values.
type externalservicesrvc struct {
	logger                 *log.Logger
	characterServiceClient characterservice.Client
}

func NewExternalservice(logger *log.Logger, client characterservice.Client) externalservice.Service {
	return &externalservicesrvc{logger: logger, characterServiceClient: client}
}

func (s *externalservicesrvc) CreateCharacter(ctx context.Context, p *externalservice.CreateCharacterPayload) (res *externalservice.Character, err error) {
	payload := &characterservice.CreatePayload{
		Name:        p.Name,
		Description: p.Description,
		Health:      p.Health,
		Experience:  p.Experience,
	}

	responseCode, err := s.characterServiceClient.CreateCharacter(ctx, payload)

	if err != nil {
		return nil, err
	}
	response := &externalservice.Character{
		ID:          responseCode,
		Name:        &p.Name,
		Description: p.Description,
		Health:      &p.Health,
		Experience:  &p.Experience,
	}
	return response, nil
}

func (s *externalservicesrvc) GetCharacter(ctx context.Context, p *externalservice.GetCharacterPayload) (res *externalservice.Character, err error) {
	payload := &characterservice.GetPayload{
		ID: p.ID,
	}

	response, err := s.characterServiceClient.GetCharacter(ctx, payload)

	if err != nil {
		return nil, err
	}
	res = &externalservice.Character{
		ID:          response.ID,
		Name:        response.Name,
		Description: response.Description,
		Health:      response.Health,
		Experience:  response.Experience,
	}

	return res, nil

}

func (s *externalservicesrvc) UpdateCharacter(ctx context.Context, p *externalservice.UpdateCharacterPayload) (err error) {
	payload := &characterservice.UpdatePayload{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Health:      p.Health,
		Experience:  p.Experience,
	}

	err = s.characterServiceClient.UpdateCharacter(ctx, payload)

	if err != nil {
		return err
	}

	return nil
}

func (s *externalservicesrvc) DeleteCharacter(ctx context.Context, p *externalservice.DeleteCharacterPayload) (err error) {
	payload := &characterservice.DeletePayload{
		ID: p.ID,
	}

	err = s.characterServiceClient.DeleteCharacter(ctx, payload)

	if err != nil {
		return err
	}

	return nil
}
