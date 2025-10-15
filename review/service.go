package review

import "books/domain"

type service struct {
	reviewRepo ReviewRepo
}

func NewService(reviewRepo ReviewRepo) Service {
	return &service{
		reviewRepo: reviewRepo,
	}
}

func (svc *service) Create(r domain.Review) (*domain.Review, error) {
	return svc.reviewRepo.Create(r)
}

func (svc *service) Get(id int) (*domain.Review, error) {
	return svc.reviewRepo.Get(id)
}

func (svc *service) 	List() ([]*domain.Review, error) {
	return svc.reviewRepo.List()
}

func (svc *service) Update(r domain.Review) (*domain.Review, error) {
	return svc.reviewRepo.Update(r)
}

func (svc *service) Delete(reviewID int) error {
	return svc.reviewRepo.Delete(reviewID)
}