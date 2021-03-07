package dto

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	ID         uuid.UUID     `json:"id"`
	UserID     uuid.UUID     `json:"userID"`
	GrandTotal float64       `json:"grandTotal"`
	Status     models.Status `json:"status"`
	Items      []OrderItem   `json:"items"`
}

type OrderItem struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"productID"`
	Quantity  int64     `json:"quantity"`
}
