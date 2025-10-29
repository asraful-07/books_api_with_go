package wish

import (
	"books/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetWish(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json");

	idStr := r.PathValue("id");
	id, err := strconv.Atoi(idStr);
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body");
		return
	}

	wish, err := h.svc.Get(id);
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error");
		return
	}

	if wish == nil {
		utils.SendError(w, http.StatusNotFound, "Wish not found");
		return
	}

	utils.SendData(w, wish, http.StatusOK);
}