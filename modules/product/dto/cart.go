package dto

import (
	uuid "github.com/satori/go.uuid"
)

type Cart struct {
	ID         uuid.UUID  `json:"id"`
	UserID     uuid.UUID  `json:"userID"`
	GrandTotal float64    `json:"price"`
	Items      []CartItem `json:"items"`
}

type CartItem struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"productID"`
	Quantity  int64     `json:"quantity"`
}
