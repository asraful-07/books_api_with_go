package review

import (
	"books/domain"
	reviewHandler "books/rest/handlers/reviews"
)

type Service interface {
	reviewHandler.Service
}

type ReviewRepo interface {
	Create(r domain.Review) (*domain.Review, error)
	Get(id int) (*domain.Review, error)
	List() ([]*domain.Review, error)
	Update(r domain.Review) (*domain.Review, error)
	Delete(reviewID int) error
}
