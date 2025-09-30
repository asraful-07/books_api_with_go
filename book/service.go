package book

import (
	"books/domain"
)

type service struct {
	bookRepo BookRepo
}

func NewService(bookRepo BookRepo) Service {
	return &service{
		bookRepo: bookRepo,
	}
}

func (svc *service) Create(book domain.Book) (*domain.Book, error) {
	return svc.bookRepo.Create(book)
}

func (svc *service) Get(id int) (*domain.Book, error) {
	return svc.bookRepo.Get(id)
}

func (svc *service) List() ([]*domain.Book, error) {
	return svc.bookRepo.List()
}

func (svc *service) Update(book domain.Book) (*domain.Book, error) {
	return svc.bookRepo.Update(book)
}

func (svc *service) Delete(id int) error {
	return svc.bookRepo.Delete(id)
}
