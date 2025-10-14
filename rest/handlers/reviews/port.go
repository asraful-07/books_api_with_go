package reviews

import "books/domain"

type Service interface {
	Create(r domain.Review) (*domain.Review, error)
	Get(id int) (*domain.Review, error)
	List() ([]*domain.Review, error)
	Update(r domain.Review) (*domain.Review, error)
	Delete(reviewID int) error
}