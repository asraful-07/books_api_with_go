package wish

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /wish", http.HandlerFunc(h.CreateWith))
	mux.Handle("GET /wishes", http.HandlerFunc(h.GetWishes))
	mux.Handle("GET /wish/{id}", http.HandlerFunc(h.GetWish))
	mux.Handle("PUT /wish/{id}", http.HandlerFunc(h.UpdateWith))
	mux.Handle("DELETE /wish/{id}", http.HandlerFunc(h.DeleteWish))
}
