package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type CartItem struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	CartID    uuid.UUID `gorm:"type:varbinary(255);column:cartID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProductID uuid.UUID `gorm:"type:varbinary(255);column:productID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Quantity  int64     `gorm:"column:quantity;type:int unsigned;default:0;not null"`
	CreatedAt time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
	Cart      Cart      `gorm:"foreignKey:CartID;references:ID"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ID"`
}

func (CartItem) TableName() string {
	return "cart_items"
}

func (cartItem *CartItem) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}
