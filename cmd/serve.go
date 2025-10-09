package cmd

import (
	"books/book"
	"books/config"
	"books/infra/db"
	"books/order"
	"books/repo"
	"books/rest"
	bookHandler "books/rest/handlers/book"
	orderHandler "books/rest/handlers/order"
	reviewHandler "books/rest/handlers/reviews"
	wishHandler "books/rest/handlers/wish"
	"books/review"
	"books/wish"
	"fmt"
	"os"
)

func Serve() {
	cfg := config.GetConfig()

	dbCon, err := db.NewConnection(cfg.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos file
	// repos
	booksRepo := repo.NewBooksRepo(dbCon)
	reviewRepo := repo.NewReviewRepo(dbCon)
	orderRepo := repo.NewOrderRepo(dbCon)
	wishRepo := repo.NewWishRepo(dbCon)

	// domains
	bokSvc := book.NewService(booksRepo)
	rvwSvc := review.NewService(reviewRepo)
	ordSvc := order.NewService(orderRepo)
	wihSvc := wish.NewService(wishRepo) 

	// handlers
	booksHandler := bookHandler.NewHandler(bokSvc)
	reviewsHandler := reviewHandler.NewHandler(rvwSvc)
	ordersHandler := orderHandler.NewHandler(ordSvc)
	wishesHandler := wishHandler.NewHandler(wihSvc)

	// start sever
	server := rest.NewServer(cfg, booksHandler, reviewsHandler, ordersHandler, wishesHandler)
	server.Start()
}