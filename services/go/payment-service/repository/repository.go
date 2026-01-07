package repository

import "monorepo-demonstration/services/go/payment-service/model"

// PaymentRepository describes persistence behaviour for payments.
type PaymentRepository interface {
	List() ([]model.Payment, error)
}

// InMemoryRepository is a trivial implementation for development.
type InMemoryRepository struct {
	payments []model.Payment
}

// NewInMemory seeds the repository with sample data.
func NewInMemory() *InMemoryRepository {
	return &InMemoryRepository{
		payments: []model.Payment{
			{ID: 201, Amount: 42.50, Currency: "USD", Status: "captured", OrderID: 101, Processed: true},
		},
	}
}

// List returns all known payments.
func (r *InMemoryRepository) List() ([]model.Payment, error) {
	return r.payments, nil
}
