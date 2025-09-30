package rest

import (
	"books/config"
	"books/rest/handlers/book"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	cfg            *config.Config
	bookHandler    *book.Handler
}

func NewServer(cfg *config.Config, bookHandler *book.Handler) *Server {
	return &Server{
		cfg:         cfg,
        bookHandler: bookHandler,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()

	// Route registration
	server.bookHandler.RegisterRoutes(mux)

	//type casting (inter to string)
	addr := ":" + strconv.Itoa(int(server.cfg.HttpPort))

	fmt.Println("Server running on http://localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}