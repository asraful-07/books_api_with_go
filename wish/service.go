package wish

import "books/domain"

type service struct {
	wishRepo WishRepo
}

func NewService(wishRepo WishRepo) Service {
	return &service{
		wishRepo: wishRepo,
	}
}

func (svc *service) Create(wish domain.Wish) (*domain.Wish, error) {
	return svc.wishRepo.Create(wish)
}

func (svc *service) List() ([]*domain.Wish, error) {
	return svc.wishRepo.List()
}

func (svc *service) Get(id int) (*domain.Wish, error) {
	return svc.wishRepo.Get(id)
}

func (svc *service) Update(wish domain.Wish) (*domain.Wish, error) {
	return svc.wishRepo.Update(wish)
}

func (svc *service) Delete(id int) error {
	return svc.wishRepo.Delete(id)
}