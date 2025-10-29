package wish

import (
	"books/utils"
	"net/http"
)

func (h *Handler) GetWishes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json");

	wishes, err := h.svc.List()
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error");
		return
	}

	utils.SendData(w, wishes, http.StatusOK);
}