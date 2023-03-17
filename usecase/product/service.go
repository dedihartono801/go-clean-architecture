package product

import (
	"encoding/json"
	"errors"

	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/external"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
)

type Service interface {
	Product() (*ProductDto, int, error)
}
type service struct {
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewProductService(
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) Product() (*ProductDto, int, error) {
	prd, err := external.Product("https://run.mocky.io/v3/70437170-6f70-4c3b-8c27-f476ff697236")
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	detailPrd := &ProductDto{}
	err = json.Unmarshal(prd, &detailPrd)
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, err
	}

	return detailPrd, customstatus.StatusOk.Code, nil
}
