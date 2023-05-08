package repository

import (
	"github.com/dedihartono801/go-clean-architecture/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SkuRepository interface {
	Create(sku *domain.Sku) error
	List() ([]domain.Sku, error)
	GetSkuById(tx *gorm.DB, id string) (domain.Sku, error)
	UpdateStockSku(tx *gorm.DB, sku *domain.Sku) error
}

type skuRepository struct {
	database *gorm.DB
}

func NewSkuRepository(database *gorm.DB) SkuRepository {
	return &skuRepository{database}
}

func (r *skuRepository) Create(sku *domain.Sku) error {
	return r.database.Table("sku").Create(sku).Error
}

func (r *skuRepository) List() ([]domain.Sku, error) {
	skus := []domain.Sku{}
	err := r.database.Table("sku").Find(&skus).Error
	return skus, err
}

func (r *skuRepository) GetSkuById(tx *gorm.DB, id string) (domain.Sku, error) {
	sku := domain.Sku{ID: id}
	err := tx.Table("sku").Clauses(clause.Locking{Strength: "UPDATE"}).First(&sku).Error
	return sku, err
}

func (r *skuRepository) UpdateStockSku(tx *gorm.DB, sku *domain.Sku) error {

	return tx.Table("sku").Save(sku).Error
}
