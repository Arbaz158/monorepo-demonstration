package service

import (
	"monorepo-demonstration/services/go/order-service/model"
	"monorepo-demonstration/services/go/order-service/repository"
)

// Service houses order domain logic.
type Service struct {
	repo repository.OrderRepository
}

// New constructs a Service.
func New(repo repository.OrderRepository) *Service {
	return &Service{repo: repo}
}

// ListOrders returns all orders; expand with richer logic later.
func (s *Service) ListOrders() ([]model.Order, error) {
	return s.repo.List()
}
