package book

import (
	"books/domain"
	"books/utils"
	"encoding/json"
	"net/http"
)

type ReqCreateBook struct {
	Title       string  `json:"title"`
	Pot         string  `json:"pot"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req ReqCreateBook

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	book := domain.Book{
	Title:       req.Title,
	Pot:         req.Pot,
	Image:       req.Image,
	Quantity:    req.Quantity,
	Price:       req.Price,       
	Description: req.Description,  
	}

	createBook, err := h.svc.Create(book)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	} 
	utils.SendData(w, createBook, http.StatusCreated)
}