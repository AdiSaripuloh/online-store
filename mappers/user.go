package mappers

import (
	dto2 "github.com/AdiSaripuloh/online-store/modules/user/dto"
	model2 "github.com/AdiSaripuloh/online-store/modules/user/models"
)

func UsersResponse(users []*model2.User) []*dto2.User {
	var response []*dto2.User
	for _, user := range users {
		response = append(response, &dto2.User{
			ID:       user.ID,
			FullName: user.FullName,
			Phone:    user.Phone,
			Email:    user.Email,
		})
	}
	return response
}
