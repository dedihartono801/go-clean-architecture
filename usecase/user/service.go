package user

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture/domain"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/repository"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
)

type Service interface {
	List() ([]domain.User, error)
	Find(id string) (domain.User, error)
	Create(input *CreateDto) (*domain.User, int, error)
	Update(id string, input *UpdateDto) (*domain.User, int, error)
	Delete(id string) error
}

type service struct {
	repository repository.UserRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewUserService(
	repository repository.UserRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

func (s *service) List() ([]domain.User, error) {
	return s.repository.List()
}

func (s *service) Find(id string) (domain.User, error) {
	return s.repository.Find(id)
}

func (s *service) Create(input *CreateDto) (*domain.User, int, error) {
	user := domain.User{
		ID:    s.identifier.NewUuid(),
		Name:  input.Name,
		Email: input.Email,
	}

	if err := s.validator.Validate(user); err != nil {
		return &user, customstatus.ErrBadRequest.Code, err
	}

	err := s.repository.Create(&user)
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	return &user, customstatus.StatusCreated.Code, nil
}

func (s *service) Update(id string, input *UpdateDto) (*domain.User, int, error) {
	user, err := s.repository.Find(id)
	if err != nil {
		return nil, customstatus.ErrNotFound.Code, errors.New(customstatus.ErrNotFound.Message)
	}

	user.Name = input.Name
	user.Email = input.Email

	if err := s.repository.Update(&user); err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	return &user, customstatus.StatusOk.Code, nil
}

func (s *service) Delete(id string) error {
	user, err := s.repository.Find(id)
	if err != nil {
		return errors.New(customstatus.ErrNotFound.Message)
	}

	return s.repository.Delete(&user)
}
