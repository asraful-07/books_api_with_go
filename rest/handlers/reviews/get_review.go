package reviews

import (
	"books/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id");
	id, err := strconv.Atoi(idStr);
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid id");
		return;
	}

	review, err := h.svc.Get(id);
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error");
		return;
	}

	if review == nil {
		utils.SendError(w, http.StatusNotFound, "Review Not found");
		return;
	}

	utils.SendData(w, review, http.StatusOK);
}