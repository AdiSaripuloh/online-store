package resolvers

import (
	repo "github.com/AdiSaripuloh/online-store/repositories"
	svc "github.com/AdiSaripuloh/online-store/services"
	"github.com/jinzhu/gorm"
)

type UserResolver struct {
	UserService svc.UserService
}

func NewUserResolver(db *gorm.DB) *UserResolver {
	userRepository := repo.NewUserRepository(db)
	userService := svc.NewUserService(userRepository)
	return &UserResolver{
		UserService: userService,
	}
}
