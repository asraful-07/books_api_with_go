package domain

// model or entity -> existence
type  Book struct {  
	ID          int     `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Pot         string  `db:"pot" json:"pot"`
	Image       string  `db:"image" json:"image"`
	Quantity    int     `db:"quantity" json:"quantity"`
	Price       float64 `db:"price" json:"price"`
	Description string  `db:"description" json:"description"`
}