package book

import "books/domain"

type Service interface {
	Create(b domain.Book) (*domain.Book, error)
	Get(bookId int)  (*domain.Book, error)
	List() ([]*domain.Book, error)
	Update(domain.Book)  (*domain.Book, error)
	Delete(bookId int) error
}