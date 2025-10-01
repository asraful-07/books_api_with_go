package cmd

import (
	"books/book"
	"books/config"
	"books/infra/db"
	"books/repo"
	"books/rest"
	bookHandler "books/rest/handlers/book"
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

	// domains
	bokSvc := book.NewService(booksRepo)
	rvwSvc := review.NewService(reviewRepo)

	// handlers
	booksHandler := bookHandler.NewHandler(bokSvc)
	reviewsHandler := reviewHAndler.NewHandler(rvwSvc)

	// start sever
	server := rest.NewServer(cfg, booksHandler, reviewsHandler)
	server.Start()
}