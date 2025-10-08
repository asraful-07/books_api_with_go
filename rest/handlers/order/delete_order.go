package order

import (
	"books/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")

	order, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	err = h.svc.Delete(order)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.SendData(w, "Successfully deleted done", http.StatusOK)
}