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
	reviewHAndler "books/rest/handlers/reviews"
	"books/review"
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

	// repos
	booksRepo := repo.NewBooksRepo(dbCon)
	reviewRepo := repo.NewReviewRepo(dbCon)
	orderRepo := repo.NewOrderRepo(dbCon)

	// domains
	bokSvc := book.NewService(booksRepo)
	rvwSvc := review.NewService(reviewRepo)
	ordSvc := order.NewService(orderRepo)

	// handlers
	booksHandler := bookHandler.NewHandler(bokSvc)
	reviewsHandler := reviewHAndler.NewHandler(rvwSvc)
	ordersHandler := orderHandler.NewHandler(ordSvc)

	// start sever
	server := rest.NewServer(cfg, booksHandler, reviewsHandler, ordersHandler)
	server.Start()
}