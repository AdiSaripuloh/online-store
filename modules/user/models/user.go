package models

import (
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID        uuid.UUID       `gorm:"column:id;primaryKey"`
	FullName  string          `gorm:"column:fullName;size:50"`
	Phone     string          `gorm:"column:phone;size:16;unique"`
	Email     string          `gorm:"column:email;size:100;unique"`
	CreatedAt time.Time       `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time       `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
	DeletedAt *time.Time      `gorm:"column:deletedAt"`
	Cart      *models.Cart    `gorm:"foreignKey:userID;references:id"`
	Orders    []*models.Order `gorm:"foreignKey:userID;references:id"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}
