package repo

import (
	"books/book"
	"books/domain"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type BookRepo interface {
	book.BookRepo
}

type bookRepo struct {
	db *sqlx.DB
}

// constructor function
func NewBooksRepo(db *sqlx.DB) BookRepo {
	return &bookRepo{db: db}
}

func (r *bookRepo) Create(b domain.Book) (*domain.Book, error) {
	query := `
	INSERT INTO books (title, pot, image, quantity, price, description)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
	err := r.db.QueryRow(
		query,
		b.Title,
		b.Pot,
		b.Image,
		b.Quantity,
		b.Price,
		b.Description,
	).Scan(&b.ID)

	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (r *bookRepo) List() ([]*domain.Book, error) {
	var books []*domain.Book
	query := `
	 SELECT id, title, pot, image, quantity, price, description 
	 FROM books
	 ORDER BY id 
	`

	err := r.db.Select(&books, query)
	if err != nil {
		return  nil, err
	}
	return books, nil
}

func (r *bookRepo) Get(bookId int) (*domain.Book, error) {
	var book domain.Book
	query := `
	SELECT id, title, pot, image, quantity, price, description 
	FROM books
	WHERE id = $1
	`

	err := r.db.Get(&book, query, bookId)
	if err != nil {
		if err == sql.ErrNoRows {
			return  nil, nil
		}
		return  nil, err
	}
	return &book, err
}

func (r *bookRepo) Update(b domain.Book) (*domain.Book, error) {
	query := `
	UPDATE books
	SET title = :title,
	    pot = :pot,
		image = :image,
		quantity = :quantity,
		price = :price,
        description = :description
	WHERE id = :id
	RETURNING id
	`
	row, err := r.db.NamedQuery(query, b)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	if row.Next() {
	 if err := row.Scan(&b.ID); err != nil {
		return  nil, err
	 }
	}
	return &b, nil
}

func (r *bookRepo) Delete(id int) error {
	query := `
	DELETE FROM books WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	return err
}


