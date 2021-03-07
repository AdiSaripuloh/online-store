package dto

import (
	uuid "github.com/satori/go.uuid"
)

type Product struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Quantity int64     `json:"quantity"`
}
