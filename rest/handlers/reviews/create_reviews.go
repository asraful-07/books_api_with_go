package reviews

import "net/http"

func (h *Handler) CreateReview(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
}