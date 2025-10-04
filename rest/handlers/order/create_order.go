package order

import "net/http"

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
}