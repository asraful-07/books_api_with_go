package domain

import "time"

// model or entity -> existence
type Review struct {  
	ID          int       `db:"id" json:"id"`
	Email       string    `db:"email" json:"email"`  
	Description string    `db:"description" json:"description"`
	ReviewDate  time.Time `db:"review_date" json:"review_date"`
}
