package reviews

import (
	"books/utils"
	"net/http"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json");

	reviews, err := h.svc.List()

	if err != nil{
		utils.SendError(w, http.StatusInternalServerError, "Internal server error");
		return
	}

	utils.SendData(w, reviews, http.StatusOK);
}