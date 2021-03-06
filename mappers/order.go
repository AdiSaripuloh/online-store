package mappers

import (
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/models"
)

func OrderResponse(order *models.Order) *dto.Order {
	var response dto.Order
	var items []dto.OrderItem
	for _, item := range order.OrderItems {
		items = append(items, dto.OrderItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	response = dto.Order{
		ID:         order.ID,
		UserID:     order.UserID,
		GrandTotal: order.GrandTotal,
		Status:     order.Status,
		Items:      items,
	}
	return &response
}
