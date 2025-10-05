package wish

import (
	"books/domain"
	wishHandler "books/rest/handlers/wish"
)

type Service interface {
	wishHandler.Service
}

type WishRepo interface {
	Create(wish domain.Wish) (*domain.Wish, error)
	List() ([]*domain.Wish, error)
	Get(id int) (*domain.Wish, error)
	Update(wish domain.Wish) (*domain.Wish, error)
	Delete(id int) error
}