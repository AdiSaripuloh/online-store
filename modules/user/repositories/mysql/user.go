package mysql

import (
	"github.com/AdiSaripuloh/online-store/database"
	model2 "github.com/AdiSaripuloh/online-store/modules/user/models"
	"github.com/AdiSaripuloh/online-store/modules/user/repositories"
	"github.com/jinzhu/gorm"
	"sync"
)

type userRepository struct {
	db *gorm.DB
}

var (
	userRepoLock sync.Once
	userRepo     repositories.UserRepository
)

func NewUserRepository(db *gorm.DB) repositories.UserRepository {
	userRepoLock.Do(func() {
		userRepo = &userRepository{
			db: db,
		}
	})

	return userRepo
}

func (u *userRepository) FindAll() ([]*model2.User, error) {
	var results []*model2.User
	err := database.Mysql.Select("id, fullName, phone, email").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
