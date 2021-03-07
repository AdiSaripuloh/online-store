package services

import (
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/repositories"
	"sync"
)

type UserService interface {
	GetAll() ([]*dto.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

var (
	userSvcLock sync.Once
	userSvc     UserService
)

func NewUserService(repository repositories.UserRepository) UserService {
	userSvcLock.Do(func() {
		userSvc = &userService{
			repository: repository,
		}
	})

	return userSvc
}

func (svc *userService) GetAll() ([]*dto.User, error) {
	results, err := svc.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return mappers.UsersResponse(results), nil
}
