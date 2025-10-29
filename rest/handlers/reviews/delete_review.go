package reviews

import (
	"books/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id");
	id, err := strconv.Atoi(idStr);
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid id");
		return;
	}

	err = h.svc.Delete(id);
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error");
		return;
	}

	utils.SendData(w, "Review delete successfully done", http.StatusOK)
}