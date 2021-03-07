package mappers

import (
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/models"
)

func UsersResponse(users []*models.User) []*dto.User {
	var response []*dto.User
	for _, user := range users {
		response = append(response, &dto.User{
			ID:       user.ID,
			FullName: user.FullName,
			Phone:    user.Phone,
			Email:    user.Email,
		})
	}
	return response
}
