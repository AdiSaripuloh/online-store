package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID    `gorm:"column:id;primaryKey"`
	Name        string       `gorm:"column:name;size:100;not null"`
	Price       float64      `gorm:"column:price;type:float;default:0;not null"`
	Quantity    int64        `gorm:"column:quantity;type:int unsigned;default:0;not null"`
	CreatedAt   time.Time    `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt   time.Time    `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
	DeletedAt   *time.Time   `gorm:"column:deletedAt"`
	CartItems   []*CartItem  `gorm:"foreignKey:ProductID;references:ID"`
	OrdersItems []*OrderItem `gorm:"foreignKey:ProductID;references:ID"`
}

func (Product) TableName() string {
	return "products"
}

func (product *Product) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}
