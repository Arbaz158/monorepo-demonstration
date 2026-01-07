package repository

import "monorepo-demonstration/services/go/order-service/model"

// OrderRepository describes persistence behaviour for orders.
type OrderRepository interface {
	List() ([]model.Order, error)
}

// InMemoryRepository is a trivial implementation for development.
type InMemoryRepository struct {
	orders []model.Order
}

// NewInMemory seeds the repository with sample data.
func NewInMemory() *InMemoryRepository {
	return &InMemoryRepository{
		orders: []model.Order{
			{ID: 101, Total: 42.50, Status: "processing"},
		},
	}
}

// List returns all known orders.
func (r *InMemoryRepository) List() ([]model.Order, error) {
	return r.orders, nil
}
