package order

import (
	"books/utils"
	"net/http"
)

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	order, err := h.svc.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.SendData(w, order, http.StatusOK)
}