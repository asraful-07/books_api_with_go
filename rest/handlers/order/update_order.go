package order

import (
	"books/domain"
	"books/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateOrder struct {
	Email       string  `json:"email"`
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	var req ReqUpdateOrder
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w,http.StatusBadRequest, "Invalid request body")
		return
	}

	order := domain.Order{
	ID:          id,
	Title:       req.Title,
	Image:       req.Image,
	Quantity:    req.Quantity,
	Price:       req.Price,       
	}

	orderBook, err := h.svc.Update(order)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if orderBook == nil {
		utils.SendError(w, http.StatusNotFound, "Order not found")
		return
	}

	utils.SendData(w, orderBook, http.StatusOK)
}