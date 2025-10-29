package wish

import (
	"books/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteWish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json");

	idStr := r.PathValue("id");
	id, err := strconv.Atoi(idStr);
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body");
		return
	}

	err = h.svc.Delete(id);

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error");
		return
	}

	utils.SendData(w, "Wish delete successfully done", http.StatusOK);
}