package wish

import (
	"books/domain"
	"books/utils"
	"encoding/json"
	"net/http"
)

type  ReqCreateWish struct {  
	Email       string  `json:"email"`  
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (h *Handler) CreateWith(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json");

	var req ReqCreateWish;
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body");
		return
	}

	wish := domain.Wish{
		Email: req.Email,
		Title: req.Title,
		Image: req.Image,
		Quantity: req.Quantity,
		Price: req.Price,
	}

	createWish, err := h.svc.Create(wish);

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error");
		return
	}

	utils.SendData(w, createWish, http.StatusCreated);
}