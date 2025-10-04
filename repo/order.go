package repo

import (
	"books/domain"
	"books/order"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type OrderRepo interface {
	order.OrderRepo
}

type orderRepo struct {
	db  *sqlx.DB
}

func NewOrderRepo(db  *sqlx.DB) OrderRepo {
	return &orderRepo{db: db}
}

func (r *orderRepo) Create(b domain.Order) (*domain.Order, error) {
	query := `
	INSERT INTO orders (title, image, quantity, price)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	err := r.db.QueryRow(
		query,
		b.Title,
		b.Image,
		b.Quantity,
		b.Price,
	).Scan(&b.ID)

	if err != nil {
		return nil, err
	}

	return &b, nil
}
