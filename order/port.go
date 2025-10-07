package order

import (
	"books/domain"
	orderHandler "books/rest/handlers/order"
)

type Service interface {
	orderHandler.Service
}

type OrderRepo interface {
	Create(order domain.Order) (*domain.Order, error)
	Get(id int) (*domain.Order, error)
	List() ([]*domain.Order, error)
	Update(o domain.Order) (*domain.Order, error)
	Delete(id int) error
}