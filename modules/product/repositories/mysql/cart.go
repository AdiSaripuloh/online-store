package mysql

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type cartRepository struct {
	db *gorm.DB
}

var (
	cartRepoLock sync.Once
	cartRepo     repositories.CartRepository
)

func NewCartRepository(db *gorm.DB) repositories.CartRepository {
	cartRepoLock.Do(func() {
		cartRepo = &cartRepository{
			db: db,
		}
	})
	return cartRepo
}

func (cr *cartRepository) Create(cart models.Cart) (*models.Cart, error) {
	err := cr.db.Create(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (cr *cartRepository) Update(cart *models.Cart) (bool, error) {
	err := cr.db.Model(&models.Cart{}).Update(cart).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cr *cartRepository) Delete(cart *models.Cart) (bool, error) {
	err := cr.db.Delete(cart).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cr *cartRepository) FindByUserIDWithItems(userID uuid.UUID) (*models.Cart, error) {
	var result models.Cart
	err := cr.db.Select("id, userID, grandTotal").Preload("CartItems").Where("userID = ?", userID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr *cartRepository) IsExists(id uuid.UUID) (*bool, error) {
	var result bool
	row := cr.db.Raw("SELECT EXISTS(SELECT 1 FROM carts WHERE userID = ?)", id).Row()
	row.Scan(&result)
	return &result, nil
}
