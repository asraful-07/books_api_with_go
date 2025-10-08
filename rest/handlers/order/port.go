package order

import "books/domain"

type Service interface {
	Create(order domain.Order) (*domain.Order, error)
    Get(id int) (*domain.Order, error)
	List() ([]*domain.Order, error)
	Update(o domain.Order) (*domain.Order, error)
	Delete(id int) error
}