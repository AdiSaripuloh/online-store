package repositories

import (
	"github.com/AdiSaripuloh/online-store/database"
	"github.com/AdiSaripuloh/online-store/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type IProductRepository interface {
	GetAll() ([]models.Product, error)
	FindByID(id string) (*models.Product, error)
	GetQuantityByID(id string) (*models.Product, error)
	UpdateQuantityByID(id uuid.UUID, quantity int64) (bool, error)
}

type productRepository struct {
	db *gorm.DB
}

var (
	productRepoLock sync.Once
	productRepo     IProductRepository
)

func NewProductRepository(db *gorm.DB) IProductRepository {
	productRepoLock.Do(func() {
		productRepo = &productRepository{
			db: db,
		}
	})

	return productRepo
}

func (u *productRepository) GetAll() ([]models.Product, error) {
	var results []models.Product
	err := database.Mysql.Select("id, name, price, quantity").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (u *productRepository) UpdateQuantityByID(id uuid.UUID, quantity int64) (bool, error) {
	err := database.Mysql.Model(&models.Product{}).Where("id = ?", id).Update(&models.Product{
		Quantity: quantity,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *productRepository) FindByID(id string) (*models.Product, error) {
	var result models.Product
	err := database.Mysql.Select("id, name, price, quantity").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (u *productRepository) GetQuantityByID(id string) (*models.Product, error) {
	var result models.Product
	err := database.Mysql.Select("id, quantity").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
