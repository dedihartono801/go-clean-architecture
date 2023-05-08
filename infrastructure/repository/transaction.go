package repository

import (
	"github.com/dedihartono801/go-clean-architecture/domain"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *domain.Transaction) error
}

type transactionRepository struct {
	database *gorm.DB
}

func NewTransactionRepository(database *gorm.DB) TransactionRepository {
	return &transactionRepository{database}
}

func (r *transactionRepository) Create(transaction *domain.Transaction) error {
	return r.database.Table("transaction").Create(transaction).Error
}
