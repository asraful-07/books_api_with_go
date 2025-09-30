package book

import (
	"books/domain"
	bookHandler "books/rest/handlers/book"
)

type Service interface {
	bookHandler.Service
}

type BookRepo interface {
	Create(b domain.Book) (*domain.Book, error)
	Get(bookId int) (*domain.Book, error)
	List() ([]*domain.Book, error)
	Update(b domain.Book)  (*domain.Book, error)
	Delete(bookId int) error
}
