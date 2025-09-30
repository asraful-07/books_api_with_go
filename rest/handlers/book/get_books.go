package book

import (
	"books/utils"
	"net/http"
)

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.svc.List()

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	utils.SendData(w, books, http.StatusOK)
}