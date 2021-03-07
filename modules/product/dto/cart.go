package dto

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	uuid "github.com/satori/go.uuid"
)

type Cart struct {
	ID         uuid.UUID  `json:"id"`
	GrandTotal float64    `json:"price"`
	Items      []CartItem `json:"items"`
}

type CartItem struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"productID"`
	Quantity  int64     `json:"quantity"`
}

type CreateCartRequest struct {
	Items []ItemsRequest `form:"items" json:"items" binding:"required"`
}

type CheckoutCartRequest struct {
	Items []ItemsRequest `form:"items" json:"items" binding:"required"`
}

func CartResponse(cart *models.Cart) *Cart {
	var response Cart
	var items []CartItem
	for _, item := range cart.CartItems {
		items = append(items, CartItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	response = Cart{
		ID:         cart.ID,
		GrandTotal: cart.GrandTotal,
		Items:      items,
	}
	return &response
}
