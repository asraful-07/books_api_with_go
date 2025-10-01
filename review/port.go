package review

import (
	"books/domain"
	reviewHandler "books/rest/handlers/reviews"
)

type Service interface {
	reviewHandler.Service
}

type ReviewRepo interface {
	Create(b domain.Review) (*domain.Review, error)
	// Get(bookId int) (*domain.Review, error)
	// List() ([]*domain.Review, error)
	// Update(b domain.Review)  (*domain.Review, error)
	// Delete(bookId int) error
}
