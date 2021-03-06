package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type OrderItem struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	UserID    string    `gorm:"type:varbinary(255);column:userID;not null"`
	ProductID string    `gorm:"type:varbinary(255);column:productID;not null"`
	Quantity  int64     `gorm:"column:quantity;type:int unsigned;default:0;not null"`
	CreatedAt time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
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
