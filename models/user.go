package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID        uuid.UUID  `gorm:"column:id;primaryKey"`
	FullName  string     `gorm:"column:fullName;size:100"`
	Phone     string     `gorm:"column:phone;size:100;unique"`
	Email     string     `gorm:"column:email;size:100;unique"`
	CreatedAt time.Time  `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time  `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
	DeletedAt *time.Time `gorm:"column:deletedAt"`
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

func DummyUser() *User {
	user := User{
		FullName: "Adi Saripuloh",
		Phone:    "1234567890",
		Email:    "adisaripuloh@gmail.com",
	}
	return &user
}
