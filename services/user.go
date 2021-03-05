package services

import (
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/repositories"
	"sync"
)

type IUserService interface {
	GetAll() ([]dto.User, error)
}

type userService struct {
	repository repositories.IUserRepository
}

var (
	userSvcLock sync.Once
	userSvc     IUserService
)

func NewUserService(repository repositories.IUserRepository) IUserService {
	userSvcLock.Do(func() {
		userSvc = &userService{
			repository: repository,
		}
	})

	return userSvc
}

func (svc *userService) GetAll() ([]dto.User, error) {
	results, err := svc.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return mappers.UsersResponse(results), nil
}
