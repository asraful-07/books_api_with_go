package wish

import "books/domain"

type Service interface {
	Create(w domain.Wish) (*domain.Wish, error)
	List() ([]*domain.Wish, error)
	Get(id int) (*domain.Wish, error)
	Update(domain.Wish) (*domain.Wish, error)
	Delete(id int) error
}
