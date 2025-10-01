package repo

import (
	"books/domain"
	"books/review"

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

func (r *reviewRepo ) Create(b domain.Review) (*domain.Review, error) {
	query := `
		INSERT INTO reviews (email, description, review_date)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(
		query,
		b.Email,
		b.Description,
		b.ReviewDate,
	).Scan(&b.ID)

	if err != nil {
		return nil, err
	}

	return &b, nil
}
