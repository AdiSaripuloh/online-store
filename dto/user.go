package dto

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
}
