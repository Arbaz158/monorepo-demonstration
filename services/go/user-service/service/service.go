package service

import (
	"monorepo-demonstration/services/go/user-service/model"
	"monorepo-demonstration/services/go/user-service/repository"
)

// Service houses user domain logic.
type Service struct {
	repo repository.UserRepository
}

// New constructs a Service.
func New(repo repository.UserRepository) *Service {
	return &Service{repo: repo}
}

// ListUsers returns all users; expand with richer logic later.
func (s *Service) ListUsers() ([]model.User, error) {
	return s.repo.List()
}
