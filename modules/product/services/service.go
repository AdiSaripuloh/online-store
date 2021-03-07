package services

import (
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	uuid "github.com/satori/go.uuid"
)

type ProductService interface {
	All() ([]*dto.Product, error)
	IsAvailable(id string, quantity int64) (bool, error)
}

type CartService interface {
	GetByUserID(id string) (*dto.Cart, error)
	Create(userID string, req dto.CreateCartRequest) (*dto.Cart, error)
	Checkout(userID string, req dto.CheckoutCartRequest) (*dto.Order, error)
}

type CartItemService interface {
	GetByID(id uuid.UUID) (*models.CartItem, error)
	UpdateQuantityByID(id uuid.UUID, quantity int64) (bool, error)
	DeleteByID(id uuid.UUID) (bool, error)
}

type OrderService interface {
	GetOrderByID(id string) (*dto.Order, error)
	GetOrderByUserID(id string) ([]*dto.Order, error)
	Pay(id string, userID string, req dto.PayOrderRequest) (*dto.Order, error)
}
