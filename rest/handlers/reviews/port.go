package reviews

import "books/domain"

type Service interface {
	Create(b domain.Review) (*domain.Review, error)
	// Get(bookId int) (*domain.Review, error)
	// List() ([]*domain.Review, error)
	// Update(b domain.Review)  (*domain.Review, error)
	// Delete(bookId int) error
}