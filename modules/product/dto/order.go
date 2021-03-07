package dto

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	uuid "github.com/satori/go.uuid"
)

type Order struct {
	ID         uuid.UUID     `json:"id"`
	GrandTotal float64       `json:"grandTotal"`
	Status     models.Status `json:"status"`
	Items      []OrderItem   `json:"items"`
}

type OrderItem struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"productID"`
	Quantity  int64     `json:"quantity"`
}

type CreateOrderRequest struct {
	UserID     uuid.UUID      `form:"userID"`
	GrandTotal float64        `form:"grantTotal"`
	Items      []ItemsRequest `form:"items"`
}

type PayOrderRequest struct {
	Amount float64 `form:"amount"`
}

func OrderResponse(order *models.Order) *Order {
	var response Order
	var items []OrderItem
	for _, item := range order.OrderItems {
		items = append(items, OrderItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	response = Order{
		ID:         order.ID,
		GrandTotal: order.GrandTotal,
		Status:     order.Status,
		Items:      items,
	}
	return &response
}

func OrdersResponse(orders []models.Order) []*Order {
	var response []*Order
	for _, order := range orders {
		response = append(response, OrderResponse(&order))
	}
	return response
}
