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
	FindByUserIDWithItems(id string) (*models.Cart, error)
	IsExists(id string) (*bool, error)
	UpdateGrandTotalByID(id uuid.UUID, grandTotal float64) (bool, error)
	DeleteByID(id uuid.UUID) (bool, error)
	UpdateQtyCartItemByID(id uuid.UUID, quantity int64) (bool, error)
	DeleteCartItemByID(id uuid.UUID) (bool, error)
}

type OrderRepository interface {
	Create(order models.Order) (*models.Order, error)
	FindByIDWithItem(id string) (*models.Order, error)
	FindByUserIDWithItem(id string) ([]models.Order, error)
	IsExists(id string) (*bool, error)
	UpdateStatusToPaid(id string) (bool, error)
}
