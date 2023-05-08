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
	BeginTransaction() (*gorm.DB, error)
	CommitTransaction(tx *gorm.DB) error
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

// BeginTransaction starts a new transaction and returns the begin operation
func (r *skuRepository) BeginTransaction() (*gorm.DB, error) {
	tx := r.database.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

// CommitTransaction commits a transaction and returns the commit operation
func (r *skuRepository) CommitTransaction(tx *gorm.DB) error {
	err := tx.Commit().Error
	if err != nil {
		return err
	}
	return nil
}
