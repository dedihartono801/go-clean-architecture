package repository

import (
	"github.com/dedihartono801/go-clean-architecture/domain"
	"gorm.io/gorm"
)

type AdminRepository interface {
	Find(id string) (*domain.Admin, error)
	Create(user *domain.Admin) error
	FindByEmail(email string) (*domain.Admin, error)
}

type adminRepository struct {
	database *gorm.DB
}

func NewAdminRepository(database *gorm.DB) AdminRepository {
	return &adminRepository{database}
}

func (r *adminRepository) Find(id string) (*domain.Admin, error) {
	admin := &domain.Admin{ID: id}
	err := r.database.First(&admin).Error
	return admin, err
}

func (r *adminRepository) Create(admin *domain.Admin) error {
	return r.database.Create(admin).Error
}

func (r *adminRepository) FindByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	err := r.database.Where("email = ?", email).First(&admin).Error
	return &admin, err
}
