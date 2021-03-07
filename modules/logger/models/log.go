package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Logs struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	ClientIP  string    `gorm:"column:clientIP;size:20;"`
	Duration  float64   `gorm:"column:duration;"`
	Method    string    `gorm:"column:method;size:10"`
	Path      string    `gorm:"column:path;"`
	Status    int       `gorm:"column:status;"`
	UserID    string    `gorm:"column:userID;"`
	Referer   string    `gorm:"column:referer;"`
	RequestID string    `gorm:"column:requestID;"`
	CreatedAt time.Time `gorm:"column:createdAt;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"column:updatedAt;type:timestamp;default:current_timestamp ON update current_timestamp"`
}

func (Logs) TableName() string {
	return "logs"
}

func (l *Logs) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", id)
}
