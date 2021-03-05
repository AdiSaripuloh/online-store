package resolvers

import (
	repo "github.com/AdiSaripuloh/online-store/repositories"
	svc "github.com/AdiSaripuloh/online-store/services"
	"github.com/jinzhu/gorm"
)

type UserResolver struct {
	UserService svc.IUserService
}

func NewResolver(db *gorm.DB) *UserResolver {
	userRepository := repo.NewUserRepository(db)
	userService := svc.NewUserService(userRepository)
	return &UserResolver{
		UserService: userService,
	}
}
