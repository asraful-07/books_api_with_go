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

func (svc *service) Get(id int) (*domain.Order, error) {
	return svc.orderRepo.Get(id)
}

func (svc *service) List() ([]*domain.Order, error) {
	return  svc.orderRepo.List()
}

func (svc *service) Update(o domain.Order) (*domain.Order, error) {
	return svc.orderRepo.Update(o)
}

func (svc *service) Delete(id int) error {
	return svc.orderRepo.Delete(id)
}