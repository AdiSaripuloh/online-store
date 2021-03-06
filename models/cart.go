package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Cart struct {
	ID         uuid.UUID `gorm:"column:id;primaryKey"`
	UserID     string    `gorm:"type:varbinary(255);column:userID;not null"`
	GrandTotal float64   `gorm:"column:grandTotal;not null"`
	CreatedAt  time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
}

func (Cart) TableName() string {
	return "carts"
}

func (user *Cart) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}
