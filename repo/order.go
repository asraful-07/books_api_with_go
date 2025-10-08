package repo

import (
	"books/domain"
	"books/order"
	"database/sql"

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

func (r *orderRepo) Create(o domain.Order) (*domain.Order, error) {
	query := `
	INSERT INTO orders (email, title, image, quantity, price)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`
	err := r.db.QueryRow(
		query,
		o.Email,
		o.Title,
		o.Image,
		o.Quantity,
		o.Price,
	).Scan(&o.ID)

	if err != nil {
		return  nil, err;
	}
	return &o, nil;
}

func (r *orderRepo) Get(id int) (*domain.Order, error) {
	var order domain.Order
	query := `
	SELECT id, email, title, image, quantity, price
	FROM orders
    WHERE id = $1
	`

	err := r.db.Get(&order, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return  nil, err
	}
	return  &order, nil;
}

func (r *orderRepo) List() ([]*domain.Order, error) {
	var order []*domain.Order

	query := `
	SELECT id, email, title, image, quantity, price
	FROM orders
    ORDER BY id
	`

	err := r.db.Select(&order, query)

	if err != nil {
		return  nil, err
	}
	return  order, nil
}

func (r *orderRepo) Update(o domain.Order) (*domain.Order, error) {
	query := `
	UPDATE orders
	SET email= :email,
	    title= : title,
		image= :image,
		quantity= :quantity,
		price= :price,
	WHERE id = :id
	RETURNING id	
	`

	row, err := r.db.NamedQuery(query, o)

	if err != nil {
		return  nil, err
	}
	defer row.Close()

	if row.Next() {
	if err := row.Scan(&o.ID); err != nil {
		return nil, err
	}
	}
	return  &o, nil;
}

func (r *orderRepo) Delete(id int) error {
	query := `DELETE FROM orders WHERE id = $1`

	_, err := r.db.Exec(query, id)
	return err
}