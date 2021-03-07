package mysql

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type orderRepository struct {
	db *gorm.DB
}

var (
	orderRepoLock sync.Once
	orderRepo     repositories.OrderRepository
)

func NewOrderRepository(db *gorm.DB) repositories.OrderRepository {
	orderRepoLock.Do(func() {
		orderRepo = &orderRepository{
			db: db,
		}
	})
	return orderRepo
}

func (or *orderRepository) Create(order models.Order) (*models.Order, error) {
	err := or.db.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (or *orderRepository) FindByIDWithItems(id uuid.UUID) (*models.Order, error) {
	var result models.Order
	err := or.db.Select("id, userID, grandTotal, status").Preload("OrderItems").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (or *orderRepository) FindByUserIDWithItems(id uuid.UUID) ([]models.Order, error) {
	var result []models.Order
	err := or.db.Select("id, userID, grandTotal, status").Preload("OrderItems").Where("userID = ?", id).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (or *orderRepository) Update(order *models.Order) (bool, error) {
	err := or.db.Model(&models.Order{}).Update(order).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (or *orderRepository) IsExists(id uuid.UUID) (*bool, error) {
	var result bool
	row := or.db.Raw("SELECT EXISTS(SELECT 1 FROM orders WHERE userID = ?)", id).Row()
	row.Scan(&result)
	return &result, nil
}
