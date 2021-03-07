package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Status string

const (
	UNPAID     Status = "UNPAID"
	PAID       Status = "PAID"
	PROCESSING Status = "PROCESSING"
	DELIVERY   Status = "DELIVERY"
	DELIVERED  Status = "DELIVERED"
	FAILED     Status = "FAILED"
)

type Order struct {
	ID         uuid.UUID   `gorm:"column:id;primaryKey"`
	UserID     uuid.UUID   `gorm:"type:varbinary(255);column:userID;not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	GrandTotal float64     `gorm:"column:grandTotal;not null"`
	Status     Status      `gorm:"column:status;type:enum('UNPAID', 'PAID', 'PROCESSING', 'DELIVERY', 'DELIVERED', 'FAILED')"`
	CreatedAt  time.Time   `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt  time.Time   `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;references:ID"`
}

func (Order) TableName() string {
	return "orders"
}

func (order *Order) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}
