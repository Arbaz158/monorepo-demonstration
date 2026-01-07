package model

// Order represents a minimal order domain model.
type Order struct {
	ID     int     `json:"id"`
	Total  float64 `json:"total"`
	Status string  `json:"status"`
}
