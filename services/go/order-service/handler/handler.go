package handler

import (
	"encoding/json"
	"net/http"

	"monorepo-demonstration/services/go/common/pkg/errors"
	"monorepo-demonstration/services/go/order-service/service"
)

// Handler wires HTTP endpoints to the service layer.
type Handler struct {
	svc *service.Service
}

// New constructs a Handler.
func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

// RegisterRoutes registers order routes onto the provided ServeMux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", h.health)
	mux.HandleFunc("/orders", h.listOrders)
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("order service ok"))
}

func (h *Handler) listOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.svc.ListOrders()
	if err != nil {
		appErr := errors.Wrap(http.StatusInternalServerError, "listing orders", err)
		http.Error(w, appErr.Error(), appErr.Code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(orders)
}
