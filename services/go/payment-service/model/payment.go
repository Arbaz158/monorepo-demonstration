package model

// Payment represents a minimal payment record.
type Payment struct {
	ID        int     `json:"id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Status    string  `json:"status"`
	OrderID   int     `json:"orderId"`
	Processed bool    `json:"processed"`
}
