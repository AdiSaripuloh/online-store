package mysql

import (
	"github.com/AdiSaripuloh/online-store/modules/user/models"
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

func (ur *userRepository) FindAll() ([]*models.User, error) {
	var results []*models.User
	err := ur.db.Select("id, fullName, phone, email").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
