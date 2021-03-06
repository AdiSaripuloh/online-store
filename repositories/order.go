package repositories

import (
	"github.com/AdiSaripuloh/online-store/database"
	"github.com/AdiSaripuloh/online-store/models"
	"github.com/jinzhu/gorm"
	"sync"
)

type IOrderRepository interface {
	Create(order models.Order) (*models.Order, error)
	FindByIDWithItem(id string) (*models.Order, error)
	FindByUserIDWithItem(id string) ([]models.Order, error)
	IsExists(id string) (*bool, error)
	Update(id string, order []models.OrderItem) (*models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

var (
	orderRepoLock sync.Once
	orderRepo     IOrderRepository
)

func NewOrderRepository(db *gorm.DB) IOrderRepository {
	orderRepoLock.Do(func() {
		orderRepo = &orderRepository{
			db: db,
		}
	})

	return orderRepo
}

func (u *orderRepository) Create(order models.Order) (*models.Order, error) {
	err := database.Mysql.Create(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (u *orderRepository) FindByIDWithItem(id string) (*models.Order, error) {
	var result models.Order
	err := database.Mysql.Select("id, userID, grandTotal, status").Preload("OrderItems").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (u *orderRepository) FindByUserIDWithItem(id string) ([]models.Order, error) {
	var result []models.Order
	err := database.Mysql.Select("id, userID, grandTotal, status").Preload("OrderItems").Where("userID = ?", id).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *orderRepository) IsExists(id string) (*bool, error) {
	var result bool
	row := database.Mysql.Raw("SELECT EXISTS(SELECT 1 FROM orders WHERE userID = ?)", id).Row()
	row.Scan(&result)
	return &result, nil
}

func (u *orderRepository) Update(id string, items []models.OrderItem) (*models.Order, error) {
	return nil, nil
}
