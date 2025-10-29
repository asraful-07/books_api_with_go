package reviews

import (
	"books/domain"
	"books/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type ReqUpdateReview struct {  
	Email       string    `json:"email"`  
	Description string    `json:"description"`
	ReviewDate  time.Time `json:"review_date"`
}

func (h *Handler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id");
	id, err := strconv.Atoi(idStr);
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid id");
		return;
	}
   
	var req ReqUpdateReview
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body");
		return;
	}

	updateReview := domain.Review{
		ID: id,
		Email: req.Email,
		Description: req.Description,
		ReviewDate: req.ReviewDate,
	}

	review, err := h.svc.Update(updateReview);
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Interval server error");
		return;
	}

	if review == nil {
		utils.SendError(w, http.StatusNotFound, "Review not found");
		return;
	}

	utils.SendData(w, review, http.StatusOK);
}