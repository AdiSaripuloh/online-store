package mysql

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type productRepository struct {
	db *gorm.DB
}

var (
	productRepoLock sync.Once
	productRepo     repositories.ProductRepository
)

func NewProductRepository(db *gorm.DB) repositories.ProductRepository {
	productRepoLock.Do(func() {
		productRepo = &productRepository{
			db: db,
		}
	})
	return productRepo
}

func (pr *productRepository) FindAll() ([]*models.Product, error) {
	var results []*models.Product
	err := pr.db.Select("id, name, price, quantity").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (pr *productRepository) UpdateQuantityByID(id uuid.UUID, quantity int64) (bool, error) {
	err := pr.db.Model(&models.Product{}).Where("id = ?", id).Update(&models.Product{
		Quantity: quantity,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (pr *productRepository) FindByID(id string) (*models.Product, error) {
	var result models.Product
	err := pr.db.Select("id, name, price, quantity").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (pr *productRepository) GetQuantityByID(id string) (*models.Product, error) {
	var result models.Product
	err := pr.db.Select("id, quantity").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
