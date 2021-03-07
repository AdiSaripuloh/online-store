package dto

import (
	"github.com/AdiSaripuloh/online-store/modules/user/models"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
}

func UsersResponse(users []*models.User) []*User {
	var response []*User
	for _, user := range users {
		response = append(response, &User{
			ID:       user.ID,
			FullName: user.FullName,
			Phone:    user.Phone,
			Email:    user.Email,
		})
	}
	return response
}
