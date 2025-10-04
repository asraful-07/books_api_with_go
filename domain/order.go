package domain

// model or entity -> existence
type  Order struct {  
	ID          int     `db:"id" json:"id"`
	Email       string  `db:"email" json:"email"`  
	Title       string  `db:"title" json:"title"`
	Image       string  `db:"image" json:"image"`
	Quantity    int     `db:"quantity" json:"quantity"`
	Price       float64 `db:"price" json:"price"`
}