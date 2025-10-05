package order

import "books/domain"

type service struct {
	orderRepo OrderRepo
}

func NewService(orderRepo OrderRepo) Service {
	return &service{
		orderRepo: orderRepo,
	}
}

func (svc *service) Create(order domain.Order) (*domain.Order, error) {
	return svc.orderRepo.Create(order)
}