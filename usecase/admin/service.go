package admin

import (
	"errors"
	"time"

	"github.com/dedihartono801/go-clean-architecture/domain"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/helper"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/middleware"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/repository"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
)

type Service interface {
	Find(id string) (*domain.Admin, error)
	Create(input *CreateDto) (*domain.Admin, int, error)
	Login(input *LoginDto) (*LoginResponse, int, error)
}

type service struct {
	repository repository.AdminRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewAdminService(
	repository repository.AdminRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

type LoginResponse struct {
	AdminId   string `json:"admin_id"`
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}

func (s *service) Find(id string) (*domain.Admin, error) {
	admin, err := s.repository.Find(id)
	return admin, err
}

func (s *service) Create(input *CreateDto) (*domain.Admin, int, error) {
	if input.Password != "" {
		input.Password = helper.EncryptPassword(input.Password)
	}
	admin := domain.Admin{
		ID:       s.identifier.NewUuid(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := s.validator.Validate(admin); err != nil {
		return nil, customstatus.ErrBadRequest.Code, errors.New(customstatus.ErrBadRequest.Message)
	}

	err := s.repository.Create(&admin)
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	return &admin, customstatus.StatusCreated.Code, nil
}

func (s *service) Login(input *LoginDto) (*LoginResponse, int, error) {

	admin, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return nil, customstatus.ErrEmailNotFound.Code, errors.New(customstatus.ErrEmailNotFound.Message)
	}

	if admin.Password != helper.EncryptPassword(input.Password) {
		return nil, customstatus.ErrPasswordWrong.Code, errors.New(customstatus.ErrPasswordWrong.Message)
	}
	expirationTime := time.Now().Add(time.Hour * time.Duration(24))
	token, err := middleware.GenerateToken(admin.ID, admin.Email)
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}
	responseParams := LoginResponse{
		AdminId:   admin.ID,
		Token:     token,
		ExpiredAt: expirationTime.Format(time.RFC3339),
	}
	return &responseParams, customstatus.StatusOk.Code, nil
}
