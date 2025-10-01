package book

import (
	"books/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	book, err := h.svc.Get(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if book == nil {
		utils.SendError(w, http.StatusNotFound, "Book not found")
		return
	}
	
	utils.SendData(w, book, http.StatusOK)
}