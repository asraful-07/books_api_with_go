package book

import (
	"books/domain"
	"books/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateBook struct {
	Title       string  `json:"title"`
	Pot         string  `json:"pot"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}


func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
     
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid id")
		return
	}

	var req ReqUpdateBook
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedBook := domain.Book{
	ID:          id,
	Title:       req.Title,
	Pot:         req.Pot,
	Image:       req.Image,
	Quantity:    req.Quantity,
	Price:       req.Price,       
	Description: req.Description,  
	}


	book, err :=  h.svc.Update(updatedBook)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if book == nil {
		utils.SendError(w, http.StatusNotFound, "Not found")
	}

	utils.SendData(w, book, http.StatusCreated)
}
