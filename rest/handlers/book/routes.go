package book

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /books", http.HandlerFunc(h.CreateBook))
	mux.Handle("GET /books", http.HandlerFunc(h.GetBooks))
	mux.Handle("GET /book/{id}", http.HandlerFunc(h.GetBookById))
	mux.Handle("GET /book/{id}", http.HandlerFunc(h.UpdateBook))
}