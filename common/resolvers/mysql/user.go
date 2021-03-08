package resolvers

import (
	"github.com/AdiSaripuloh/online-store/modules/user/repositories/mysql"
	"github.com/AdiSaripuloh/online-store/modules/user/services"
	"github.com/jinzhu/gorm"
)

type UserResolver struct {
	UserService services.UserService
}

func NewUserResolver(db *gorm.DB) *UserResolver {
	userRepository := mysql.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	return &UserResolver{
		UserService: userService,
	}
}
