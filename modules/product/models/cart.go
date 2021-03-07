package models

import (
	"github.com/AdiSaripuloh/online-store/modules/user/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Cart struct {
	ID         uuid.UUID   `gorm:"column:id;primaryKey"`
	UserID     uuid.UUID   `gorm:"type:varbinary(255);column:userID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	GrandTotal float64     `gorm:"column:grandTotal;not null"`
	CreatedAt  time.Time   `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt  time.Time   `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
	User       models.User `gorm:"foreignKey:UserID;references:ID"`
	CartItems  []CartItem  `gorm:"foreignKey:CartID;references:ID"`
}

func (Cart) TableName() string {
	return "carts"
}

func (cart *Cart) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}
