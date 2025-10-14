package repo

import (
	"books/domain"
	"books/review"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ReviewRepo interface {
	review.ReviewRepo
}

type reviewRepo struct {
	db *sqlx.DB
}

func NewReviewRepo(db *sqlx.DB)  ReviewRepo {
	return &reviewRepo{db: db}
}

func (r *reviewRepo ) Create(review domain.Review) (*domain.Review, error) {
	query := `
	INSERT INTO reviews (email, description, review_date)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	err := r.db.QueryRow(
		query,
		review.Email,
		review.Description,
		review.ReviewDate,
	).Scan(review.ID)

	if err != nil {
		return  nil, err
	}
	return  &review, nil;
}


func (r *reviewRepo) Get(id int) (*domain.Review, error) {
	var review domain.Review
	query := `
	SELECT id, email, description, review_date
	FROM reviews
	WHERE id = $1
	`

	err := r.db.Get(&review, query, id)
	if err != nil {
		if err == sql.ErrNoRows{
			return  nil, nil
		}
		return  nil, err
	}
	return &review, nil
}

func (r *reviewRepo) List() ([]*domain.Review, error) {
	var review []*domain.Review
	query := `
	SELECT id, email, description, review_date
	FROM reviews
	ORDER BY id
	`

	err := r.db.Select(&review, query)
	if err != nil {
		return  nil, err
	}
	return  review, nil
}


func (r *reviewRepo) Update(review domain.Review) (*domain.Review, error) {
	query := `
	UPDATE reviews
	SET email = :email,
	    description = :description,
		review_date = :review_date,
    WHERE id = :id
	RETURNING id
	`

	rows, err := r.db.NamedQuery(query, review)
	if err != nil {
		return  nil, err
	}
	defer rows.Close()
	if rows.Next() {
	  if err:= rows.Scan(review.ID); err != nil {
		return  nil, err
	  }
	  return nil, err
	}
	return &review, nil
}

func (r *reviewRepo) Delete(reviewID int) error {
	query := `DELETE FORM reviews WHERE id = $1`

	_, err := r.db.Exec(query, reviewID)

	return err
}