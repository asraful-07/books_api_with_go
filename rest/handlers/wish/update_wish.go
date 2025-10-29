package wish

import (
	"books/domain"
	"books/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type  ReqUpdateWish struct {  
	Email       string  `json:"email"`  
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (h *Handler) UpdateWith(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json");

	idStr := r.PathValue("id");
	id, err := strconv.Atoi(idStr);
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body");
		return
	}

	var req ReqUpdateWish;
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body");
		return
	}

	wish := domain.Wish{
		ID: id,
		Email: req.Email,
		Title: req.Title,
		Image: req.Image,
		Quantity: req.Quantity,
		Price: req.Price,
	}

	updateWish, err := h.svc.Update(wish);

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error");
		return
	}
	
	if updateWish == nil {
		utils.SendError(w, http.StatusNotFound, "Wish not found")
		return
	}

	utils.SendData(w, updateWish, http.StatusCreated);
}