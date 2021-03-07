package mysql

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type cartItemRepository struct {
	db *gorm.DB
}

var (
	cartItemRepoLock sync.Once
	cartItemRepo     repositories.CartItemRepository
)

func NewCartItemRepository(db *gorm.DB) repositories.CartItemRepository {
	cartItemRepoLock.Do(func() {
		cartItemRepo = &cartItemRepository{
			db: db,
		}
	})
	return cartItemRepo
}

func (cir *cartItemRepository) Create(cartItem models.CartItem) (*models.CartItem, error) {
	err := cir.db.Create(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}

func (cir *cartItemRepository) FindByID(id uuid.UUID) (*models.CartItem, error) {
	var result models.CartItem
	err := cir.db.Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cir *cartItemRepository) Update(cartItem *models.CartItem) (bool, error) {
	err := cir.db.Model(&models.CartItem{}).Update(cartItem).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cir *cartItemRepository) Delete(cartItem *models.CartItem) (bool, error) {
	err := cir.db.Delete(cartItem).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
