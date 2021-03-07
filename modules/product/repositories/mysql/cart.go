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

func (cr *cartRepository) FindByUserIDWithItems(id string) (*models.Cart, error) {
	var result models.Cart
	err := cr.db.Select("id, userID, grandTotal").Preload("CartItems").Where("userID = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (cr *cartRepository) IsExists(id string) (*bool, error) {
	var result bool
	row := cr.db.Raw("SELECT EXISTS(SELECT 1 FROM carts WHERE userID = ?)", id).Row()
	row.Scan(&result)
	return &result, nil
}

func (cr *cartRepository) UpdateGrandTotalByID(id uuid.UUID, grandTotal float64) (bool, error) {
	err := cr.db.Model(&models.Cart{}).Where("id = ?", id).Update(&models.Cart{
		GrandTotal: grandTotal,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cr *cartRepository) DeleteByID(id uuid.UUID) (bool, error) {
	err := cr.db.Model(&models.Cart{}).Delete(&models.Cart{
		ID: id,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cr *cartRepository) UpdateQtyCartItemByID(id uuid.UUID, quantity int64) (bool, error) {
	err := cr.db.Model(&models.CartItem{}).Where("id = ?", id).Update(&models.CartItem{
		Quantity: quantity,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cr *cartRepository) DeleteCartItemByID(id uuid.UUID) (bool, error) {
	err := cr.db.Model(&models.CartItem{}).Delete(&models.CartItem{
		ID: id,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
