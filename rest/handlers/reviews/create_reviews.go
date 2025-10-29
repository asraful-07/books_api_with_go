package reviews

import (
	"books/domain"
	"books/utils"
	"encoding/json"
	"net/http"
	"time"
)

type ReqCreateReview struct {  
	Email       string    `json:"email"`  
	Description string    `json:"description"`
	ReviewDate  time.Time `json:"review_date"`
}

func (h *Handler) CreateReview(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req ReqCreateReview 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return;
	}

	review := domain.Review{
		Email: req.Email,
		Description: req.Description,
		ReviewDate: req.ReviewDate,
	}

	createReview, err := h.svc.Create(review)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "INternal server error");
	    return;
	}

	utils.SendData(w, createReview, http.StatusCreated);
}