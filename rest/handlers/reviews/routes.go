package reviews

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /review", http.HandlerFunc(h.CreateReview))
}