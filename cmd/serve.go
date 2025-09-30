package cmd

import (
	"books/book"
	"books/config"
	"books/infra/db"
	"books/repo"
	"books/rest"
	bookHandler "books/rest/handlers/book"
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

	// domains
	btkSvc := book.NewService(booksRepo)

	// handlers
	booksHandler := bookHandler.NewHandler(btkSvc)

	// start sever
	server := rest.NewServer(cfg, booksHandler)
	server.Start()
}