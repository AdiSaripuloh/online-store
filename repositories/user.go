package repositories

import (
	"github.com/AdiSaripuloh/online-store/database"
	"github.com/AdiSaripuloh/online-store/models"
	"github.com/jinzhu/gorm"
	"sync"
)

type UserRepository interface {
	FindAll() ([]*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

var (
	userRepoLock sync.Once
	userRepo     UserRepository
)

func NewUserRepository(db *gorm.DB) UserRepository {
	userRepoLock.Do(func() {
		userRepo = &userRepository{
			db: db,
		}
	})

	return userRepo
}

func (u *userRepository) FindAll() ([]*models.User, error) {
	var results []*models.User
	err := database.Mysql.Select("id, fullName, phone, email").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
