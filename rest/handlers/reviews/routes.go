package reviews

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /review", http.HandlerFunc(h.CreateReview))
	mux.Handle("GET /reviews", http.HandlerFunc(h.GetReviews))
	mux.Handle("GET /review/{id}", http.HandlerFunc(h.GetReview))
	mux.Handle("PUT /review/{id}", http.HandlerFunc(h.UpdateReview))
	mux.Handle("DELETE /review/{id}", http.HandlerFunc(h.DeleteReview))
}