package order

import (
	"books/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	order, err := h.svc.Get(id)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if order != nil {
		utils.SendError(w, http.StatusNotFound, "Order not found")
		return
	}

	utils.SendData(w, order, http.StatusOK)
}