package mappers

import (
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/models"
)

func CartResponse(cart *models.Cart) *dto.Cart {
	var response dto.Cart
	var items []dto.CartItem
	for _, item := range cart.CartItems {
		items = append(items, dto.CartItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	response = dto.Cart{
		ID:         cart.ID,
		UserID:     cart.UserID,
		GrandTotal: cart.GrandTotal,
		Items:      items,
	}
	return &response
}
