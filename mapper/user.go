package mapper

import (
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/models"
	"github.com/gin-gonic/gin"
)

func UsersResponse(users []models.User) gin.H {
	var response []dto.User
	for _, user := range users {
		response = append(response, dto.User{
			ID:       user.ID,
			FullName: user.FullName,
			Phone:    user.Phone,
			Email:    user.Email,
		})
	}
	return success(response)
}
