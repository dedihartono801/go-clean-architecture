package transaction

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dedihartono801/go-clean-architecture/domain"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/identifier"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/repository"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/validator"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	Checkout(input *CheckoutDto, adminID string) (*domain.Transaction, int, error)
}

type service struct {
	dbTransactionRepository repository.DbTransactionRepository
	transactionRepository   repository.TransactionRepository
	skuRepository           repository.SkuRepository
	validator               validator.Validator
	identifier              identifier.Identifier
	skuMutex                sync.Mutex
}

func NewTransactionService(
	dbTransactionRepository repository.DbTransactionRepository,
	transactionRepository repository.TransactionRepository,
	skuRepository repository.SkuRepository,
	validator validator.Validator,
	identifier identifier.Identifier,
) Service {
	return &service{
		dbTransactionRepository: dbTransactionRepository,
		transactionRepository:   transactionRepository,
		skuRepository:           skuRepository,
		validator:               validator,
		identifier:              identifier,
		skuMutex:                sync.Mutex{},
	}
}

func (s *service) Checkout(input *CheckoutDto, adminID string) (*domain.Transaction, int, error) {

	checkout := CheckoutDto{
		Items: input.Items,
	}

	if err := s.validator.Validate(checkout); err != nil {
		return nil, customstatus.ErrBadRequest.Code, err
	}

	// create an errgroup.Group instance
	var g errgroup.Group

	tx, err := s.dbTransactionRepository.BeginTransaction()
	if err != nil {
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}

	totalPrice := 0
	totalQuantity := 0
	for _, items := range input.Items {
		item := items
		g.Go(func() error {
			s.skuMutex.Lock()
			defer s.skuMutex.Unlock()

			// find the selected SKU
			sku, err := s.skuRepository.GetSkuById(tx, item.ID)
			if err != nil {
				tx.Rollback()
				return errors.New(customstatus.ErrNotFound.Message)
			}
			if sku.Stock < item.Quantity {
				tx.Rollback()
				return fmt.Errorf("stock item %s tidak mencukupi hanya ada %d", sku.Name, sku.Stock)
			}

			sku.Stock -= item.Quantity

			err = s.skuRepository.UpdateStockSku(tx, &sku)
			if err != nil {
				tx.Rollback()
				return errors.New(customstatus.ErrInternalServerError.Message)
			}

			totalPrice += sku.Price * item.Quantity

			return nil
		})
		totalQuantity += item.Quantity
	}

	// wait for all goroutines to finish
	if err := g.Wait(); err != nil {
		return nil, customstatus.ErrBadRequest.Code, errors.New(err.Error())
	}

	trx := &domain.Transaction{
		ID:               s.identifier.NewUuid(),
		AdminID:          adminID,
		TotalQuantity:    totalQuantity,
		TotalTransaction: totalPrice,
	}

	err = s.transactionRepository.Create(tx, trx)
	if err != nil {
		tx.Rollback()
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}

	err = s.dbTransactionRepository.CommitTransaction(tx)
	if err != nil {
		tx.Rollback()
		return nil, customstatus.ErrInternalServerError.Code, errors.New(customstatus.ErrInternalServerError.Message)
	}

	return nil, customstatus.StatusCreated.Code, nil
}
