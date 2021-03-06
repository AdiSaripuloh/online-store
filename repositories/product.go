package repositories

import (
	"github.com/AdiSaripuloh/online-store/database"
	"github.com/AdiSaripuloh/online-store/models"
	"github.com/jinzhu/gorm"
	"sync"
)

type IProductRepository interface {
	GetAll() ([]models.Product, error)
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
