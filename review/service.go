package review

import (
	"books/domain"
)

type service struct {
	reviewRepo ReviewRepo
}

func NewService(reviewRepo ReviewRepo) Service {
	return &service{
		reviewRepo: reviewRepo,
	}
}

func (svc *service) Create(review domain.Review) (*domain.Review, error) {
	return svc.reviewRepo.Create(review)
}

// func (svc *service) Get(id int) (*domain.Book, error) {
// 	return svc.reviewRepo.Get(id)
// }

// func (svc *service) List() ([]*domain.Book, error) {
// 	return svc.reviewRepo.List()
// }

// func (svc *service) Update(book domain.Book) (*domain.Book, error) {
// 	return svc.reviewRepo.Update(book)
// }

// func (svc *service) Delete(id int) error {
// 	return svc.reviewRepo.Delete(id)
// }
