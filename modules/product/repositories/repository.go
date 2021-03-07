package repositories

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	uuid "github.com/satori/go.uuid"
)

type ProductRepository interface {
	FindAll() ([]*models.Product, error)
	FindByID(id string) (*models.Product, error)
	GetQuantityByID(id string) (*models.Product, error)
	UpdateQuantityByID(id uuid.UUID, quantity int64) (bool, error)
}

type CartRepository interface {
	Create(cart models.Cart) (*models.Cart, error)
	Update(*models.Cart) (bool, error)
	Delete(*models.Cart) (bool, error)
	FindByUserIDWithItems(userID uuid.UUID) (*models.Cart, error)
	IsExists(id uuid.UUID) (*bool, error)
}

type CartItemRepository interface {
	Create(cartItem models.CartItem) (*models.CartItem, error)
	FindByID(id uuid.UUID) (*models.CartItem, error)
	Update(cartItem *models.CartItem) (bool, error)
	Delete(cartItem *models.CartItem) (bool, error)
}

type OrderRepository interface {
	Create(order models.Order) (*models.Order, error)
	FindByIDWithItem(id string) (*models.Order, error)
	FindByUserIDWithItem(id string) ([]models.Order, error)
	IsExists(id string) (*bool, error)
	UpdateStatusToPaid(id string) (bool, error)
}
