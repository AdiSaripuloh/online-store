package services

import "github.com/AdiSaripuloh/online-store/modules/user/dto"

type UserService interface {
	GetAll() ([]*dto.User, error)
}
