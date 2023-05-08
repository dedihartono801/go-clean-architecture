package sku

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture/domain"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/repository"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
)

type Service interface {
	Create(input *CreateDto) (*domain.Sku, int, error)
	List() ([]domain.Sku, error)
}

type service struct {
	repository repository.SkuRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewSkuService(
	repository repository.SkuRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) Create(input *CreateDto) (*domain.Sku, int, error) {
	sku := domain.Sku{
		ID:    s.identifier.NewUuid(),
		Name:  input.Name,
		Stock: input.Stock,
		Price: input.Price,
	}

	if err := s.validator.Validate(sku); err != nil {
		return &sku, customstatus.ErrBadRequest.Code, err
	}

	err := s.repository.Create(&sku)
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	return &sku, customstatus.StatusCreated.Code, nil
}

func (s *service) List() ([]domain.Sku, error) {
	return s.repository.List()
}
