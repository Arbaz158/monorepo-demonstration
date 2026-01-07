package repository

import "monorepo-demonstration/services/go/user-service/model"

// UserRepository describes persistence behaviour for users.
type UserRepository interface {
	List() ([]model.User, error)
}

// InMemoryRepository is a trivial implementation for development.
type InMemoryRepository struct {
	users []model.User
}

// NewInMemory initializes repository with seed data.
func NewInMemory() *InMemoryRepository {
	return &InMemoryRepository{users: []model.User{{ID: 1, Name: "Ada Lovelace", Email: "ada@example.com"}}}
}

// List returns all known users.
func (r *InMemoryRepository) List() ([]model.User, error) {
	return r.users, nil
}
