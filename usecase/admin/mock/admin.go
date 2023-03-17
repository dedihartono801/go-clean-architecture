package mock

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture/domain"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	repoMock "github.com/dedihartono801/go-clean-architecture/infrastructure/repository/mock"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
	"github.com/dedihartono801/go-clean-architecture/usecase/admin"
)

type Service interface {
	Find(id string) (*domain.Admin, error)
	Create(input *admin.CreateDto) (*domain.Admin, int, error)
	Login(input *admin.LoginDto) (*admin.LoginResponse, int, error)
}

type mockService struct {
	repository repoMock.AdminRepository
	validator  validator.Validator
	identifier identifier.Identifier
}

func NewMockAdminService(
	repository repoMock.AdminRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &mockService{
		repository: repository,
		validator:  validator,
		identifier: identifier,
	}
}

func (s *mockService) Find(id string) (*domain.Admin, error) {
	admin := &domain.Admin{
		ID:       "4d35bf38-8c50-4c85-8072-fd9794803a167",
		Name:     "diding",
		Email:    "diding@gmail.com",
		Password: "56334b8232e95fb59b0fc93f2bc0d5c1fdbf5f120d91ac9f5d4c9db14544e007dd163cba5af3de3f027a6d47280f1407c19a5c1b8fc8ca10a4d7ef431341f135",
	}
	if admin.ID != id {
		return nil, errors.New("admin not found")
	}
	return admin, nil
}

func (s *mockService) Create(input *admin.CreateDto) (*domain.Admin, int, error) {
	admin := &domain.Admin{
		ID:       "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Name:     "diding",
		Email:    "diding@gmail.com",
		Password: "56334b8232e95fb59b0fc93f2bc0d5c1fdbf5f120d91ac9f5d4c9db14544e007dd163cba5af3de3f027a6d47280f1407c19a5c1b8fc8ca10a4d7ef431341f135",
	}
	return admin, customstatus.StatusCreated.Code, nil
}

func (s *mockService) Login(input *admin.LoginDto) (*admin.LoginResponse, int, error) {
	dt := &admin.LoginResponse{
		AdminId:   "4d35bf38-8c50-4c85-8072-fd9794803a16",
		Token:     "56334b8232e95fb59b0fc93f2bc0d5c1fdbf5f120d91ac9f5d4c9db14544e007dd163cba5af3de3f027a6d47280f1407c19a5c1b8fc8ca10a4d7ef431341f135",
		ExpiredAt: "2023-03-17T08:40:25Z",
	}
	return dt, customstatus.StatusOk.Code, nil
}
