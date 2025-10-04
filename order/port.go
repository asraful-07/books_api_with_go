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
	// Get(bookId int) (*domain.Review, error)
	// List() ([]*domain.Review, error)
	// Update(b domain.Review)  (*domain.Review, error)
	// Delete(bookId int) error
}