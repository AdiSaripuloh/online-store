package services

import (
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	uuid "github.com/satori/go.uuid"
)

type ProductService interface {
	All() ([]*dto.Product, *responses.HttpError)
	IsAvailable(id string, quantity int64) (bool, *responses.HttpError)
}

type CartService interface {
	GetByUserID(id string) (*dto.Cart, *responses.HttpError)
	Create(userID string, req dto.CreateCartRequest) (*dto.Cart, *responses.HttpError)
	Checkout(userID string, req dto.CheckoutCartRequest) (*dto.Order, *responses.HttpError)
}

type CartItemService interface {
	GetByID(id uuid.UUID) (*models.CartItem, *responses.HttpError)
	UpdateQuantityByID(id uuid.UUID, quantity int64) (bool, *responses.HttpError)
	DeleteByID(id uuid.UUID) (bool, *responses.HttpError)
}

type OrderService interface {
	GetOrderByID(id string) (*dto.Order, *responses.HttpError)
	GetOrderByUserID(id string) ([]*dto.Order, *responses.HttpError)
	Pay(id string, userID string, req dto.PayOrderRequest) (*dto.Order, *responses.HttpError)
}
