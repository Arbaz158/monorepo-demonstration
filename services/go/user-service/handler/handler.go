package handler

import (
	"encoding/json"
	"net/http"

	"monorepo-demonstration/services/go/common/pkg/errors"
	"monorepo-demonstration/services/go/user-service/service"
)

// Handler wires HTTP endpoints to the service layer.
type Handler struct {
	svc *service.Service
}

// New constructs a Handler.
func New(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

// RegisterRoutes registers user routes onto the provided ServeMux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", h.health)
	mux.HandleFunc("/users", h.listUsers)
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("user service ok"))
}

func (h *Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.svc.ListUsers()
	if err != nil {
		appErr := errors.Wrap(http.StatusInternalServerError, "listing users", err)
		http.Error(w, appErr.Error(), appErr.Code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}
