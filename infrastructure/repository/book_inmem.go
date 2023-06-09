package repository

import (
	"errors"

	"github.com/dedihartono801/go-clean-architecture/domain"
	"github.com/dedihartono801/go-clean-architecture/infrastructure/customstatus"
	"github.com/google/uuid"
)

type BookRepository interface {
	List() ([]*domain.Book, error)
	Find(id string) (*domain.Book, error)
	Create(book *domain.Book) error
	Update(book *domain.Book) error
	Delete(book *domain.Book) error
}

type bookRepository struct {
	data map[string]*domain.Book
}

func NewBookRepository() BookRepository {
	var mp = map[string]*domain.Book{}
	return &bookRepository{data: mp}
}

func (r *bookRepository) List() ([]*domain.Book, error) {
	var books []*domain.Book
	for _, book := range r.data {
		books = append(books, book)
	}
	return books, nil
}

func (r *bookRepository) Find(id string) (*domain.Book, error) {
	if r.data[id] == nil {
		return nil, errors.New(customstatus.ErrNotFound.Message)
	}
	return r.data[id], nil
}

func (r *bookRepository) Create(book *domain.Book) error {
	book.ID = uuid.New().String()
	r.data[book.ID] = book
	return nil
}

func (r *bookRepository) Update(book *domain.Book) error {
	_, err := r.Find(book.ID)
	if err != nil {
		return errors.New(customstatus.ErrNotFound.Message)
	}
	r.data[book.ID] = book
	return nil
}

func (r *bookRepository) Delete(book *domain.Book) error {
	r.data[book.ID] = nil
	return nil
}
