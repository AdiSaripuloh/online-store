package services

import (
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/requests"
)

type ProductService interface {
	GetAll() ([]dto.Product, error)
	IsAvailable(id string, quantity int64) (bool, error)
}

type CartService interface {
	GetCartByUserID(id string) (*dto.Cart, error)
	Create(userID string, req requests.CreateCart) (*dto.Cart, error)
	Checkout(userID string, req requests.Checkout) (*dto.Order, error)
}

type OrderService interface {
	GetOrderByID(id string) (*dto.Order, error)
	GetOrderByUserID(id string) ([]*dto.Order, error)
	Pay(id string, userID string, req requests.PayOrder) (*dto.Order, error)
}
