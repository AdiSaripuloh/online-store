package repositories

import (
	"github.com/AdiSaripuloh/online-store/database"
	"github.com/AdiSaripuloh/online-store/models"
	"github.com/jinzhu/gorm"
	"sync"
)

type IUserRepository interface {
	GetAll() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

var (
	userRepoLock sync.Once
	userRepo     IUserRepository
)

func NewUserRepository(db *gorm.DB) IUserRepository {
	userRepoLock.Do(func() {
		userRepo = &userRepository{
			db: db,
		}
	})

	return userRepo
}

func (u *userRepository) GetAll() ([]models.User, error) {
	var results []models.User
	err := database.Mysql.Select("id, fullName, phone, email").Where("deletedAt IS NULL").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}
