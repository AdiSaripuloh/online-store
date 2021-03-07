package services

import (
	"github.com/AdiSaripuloh/online-store/mappers"
	dto2 "github.com/AdiSaripuloh/online-store/modules/user/dto"
	"github.com/AdiSaripuloh/online-store/modules/user/repositories"
	"sync"
)

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

func (svc *userService) GetAll() ([]*dto2.User, error) {
	results, err := svc.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return mappers.UsersResponse(results), nil
}
