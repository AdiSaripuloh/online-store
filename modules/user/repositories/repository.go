package repositories

import model2 "github.com/AdiSaripuloh/online-store/modules/user/models"

type UserRepository interface {
	FindAll() ([]*model2.User, error)
}
