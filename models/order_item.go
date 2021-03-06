package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type OrderItem struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	OrderID   uuid.UUID `gorm:"type:varbinary(255);column:orderID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ProductID uuid.UUID `gorm:"type:varbinary(255);column:productID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Quantity  int64     `gorm:"column:quantity;type:int unsigned;default:0;not null"`
	CreatedAt time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
	Order     Order     `gorm:"foreignKey:OrderID;references:ID"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ID"`
}

func (OrderItem) TableName() string {
	return "order_items"
}

func (user *OrderItem) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}
