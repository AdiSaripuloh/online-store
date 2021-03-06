package models

import (
	"database/sql/driver"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type status string

const (
	UNPAID     status = "UNPAID"
	PAID       status = "PAID"
	PROCESSING status = "PROCESSING"
	DELIVERY   status = "DELIVERY"
	DELIVERED  status = "DELIVERED"
	FAILED     status = "FAILED"
)

type Order struct {
	ID         uuid.UUID `gorm:"column:id;primaryKey"`
	UserID     string    `gorm:"type:varbinary(255);column:userID;not null"`
	GrandTotal float64   `gorm:"column:grandTotal;not null"`
	Status     string    `gorm:"column:status;type:enum('UNPAID', 'PAID', 'PROCESSING', 'DELIVERY', 'DELIVERED', 'FAILED')"`
	CreatedAt  time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
}

func (Order) TableName() string {
	return "orders"
}

func (p *status) Scan(value interface{}) error {
	*p = status(value.([]byte))
	return nil
}

func (p status) Value() (driver.Value, error) {
	return string(p), nil
}

func (user *Order) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}
