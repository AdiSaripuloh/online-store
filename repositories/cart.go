package repositories

import (
	"github.com/AdiSaripuloh/online-store/database"
	"github.com/AdiSaripuloh/online-store/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type ICartRepository interface {
	Create(cart models.Cart) (*models.Cart, error)
	FindByUserIDWithItem(id string) (*models.Cart, error)
	IsExists(id string) (*bool, error)
	UpdateGrandTotalByID(id uuid.UUID, grandTotal float64) (bool, error)
	DeleteByID(id uuid.UUID) (bool, error)
	UpdateQtyCartItemByID(id uuid.UUID, quantity int64) (bool, error)
}

type cartRepository struct {
	db *gorm.DB
}

var (
	cartRepoLock sync.Once
	cartRepo     ICartRepository
)

func NewCartRepository(db *gorm.DB) ICartRepository {
	cartRepoLock.Do(func() {
		cartRepo = &cartRepository{
			db: db,
		}
	})

	return cartRepo
}

func (u *cartRepository) Create(cart models.Cart) (*models.Cart, error) {
	err := database.Mysql.Create(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}

func (u *cartRepository) FindByUserIDWithItem(id string) (*models.Cart, error) {
	var result models.Cart
	err := database.Mysql.Select("id, userID, grandTotal").Preload("CartItems").Where("userID = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (u *cartRepository) IsExists(id string) (*bool, error) {
	var result bool
	row := database.Mysql.Raw("SELECT EXISTS(SELECT 1 FROM carts WHERE userID = ?)", id).Row()
	row.Scan(&result)
	return &result, nil
}

func (u *cartRepository) UpdateGrandTotalByID(id uuid.UUID, grandTotal float64) (bool, error) {
	err := database.Mysql.Model(&models.Cart{}).Where("id = ?", id).Update(&models.Cart{
		GrandTotal: grandTotal,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *cartRepository) DeleteByID(id uuid.UUID) (bool, error) {
	err := database.Mysql.Model(&models.Cart{}).Delete(&models.Cart{
		ID: id,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *cartRepository) UpdateQtyCartItemByID(id uuid.UUID, quantity int64) (bool, error) {
	err := database.Mysql.Model(&models.CartItem{}).Where("id = ?", id).Update(&models.CartItem{
		Quantity: quantity,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
