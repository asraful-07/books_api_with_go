package repo

import (
	"books/domain"
	"books/wish"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type WishRepo interface {
	wish.WishRepo
}

type wishRepo struct {
	db *sqlx.DB
}

func NewWishRepo(db *sqlx.DB) WishRepo {
	return &wishRepo{db: db}
}

func (r *wishRepo) Create(w domain.Wish) (*domain.Wish, error) {
	query := `
	INSERT INTO wishes (email, title, image, quantity, price)
	VALUES (1$, $2, $3, $4, $5)
	RETURNING id
	`
	err := r.db.QueryRow(
		query,
		w.Title,
        w.Email,
		w.Image,
		w.Quantity,
		w.Price,
	).Scan(&w.ID)
	
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (r *wishRepo) List() ([]*domain.Wish, error) {
	var wish []*domain.Wish
	query := `
	SELECT id, email, title, image, quantity, price
	FROM wishes
	ORDER BY id 
	`

	err := r.db.Select(&wish, query)
	if err != nil {
		return nil, err
	}
	return wish, nil
}

func (r *wishRepo) Get(id int) (*domain.Wish, error) {
	var wish domain.Wish
	query := `
	SELECT id, email, title, image, quantity, price
	FROM wishes
	WHERE id = $1
	`
	err := r.db.Get( &wish, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil , nil
		}
		return nil, err
	}
	return &wish, nil 
}

func (r *wishRepo) Update(wish domain.Wish) (*domain.Wish, error) {
	query := `
	UPDATE wishes 
	SET email = :email,
	    title = : title,
		image = :image,
		quantity = :quantity
        price = :price 
	WHERE id = :id	
	RETURNING id
	`
	row, err := r.db.NamedQuery(query, wish) 
	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
     	if err := row.Scan(&wish.ID); err != nil {
			return  nil, err
		}
	}
	return &wish, nil
}

func (r *wishRepo) Delete(id int) error {
	query := `DELETE FROM wishes WHERE id = $1`

	_, err := r.db.Exec(query, id)
	return err
}