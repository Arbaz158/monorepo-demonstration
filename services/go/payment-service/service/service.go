package service

import (
	"monorepo-demonstration/services/go/payment-service/model"
	"monorepo-demonstration/services/go/payment-service/repository"
)

// Service houses payment domain logic.
type Service struct {
	repo repository.PaymentRepository
}

// New constructs a Service.
func New(repo repository.PaymentRepository) *Service {
	return &Service{repo: repo}
}

// ListPayments returns all payments; expand with richer logic later.
func (s *Service) ListPayments() ([]model.Payment, error) {
	return s.repo.List()
}
