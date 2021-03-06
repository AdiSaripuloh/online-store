package services

import (
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/repositories"
	"github.com/AdiSaripuloh/online-store/requests"
	"sync"
)

type IOrderService interface {
	GetOrderByID(id string) (*dto.Order, error)
	GetOrderByUserID(id string) ([]*dto.Order, error)
	Pay(userID string, req requests.PayOrder) (*dto.Order, error)
}

type orderService struct {
	orderRepo   repositories.IOrderRepository
	productRepo repositories.IProductRepository
}

var (
	orderSvcLock sync.Once
	orderSvc     IOrderService
)

func NewOrderService(orderRepo repositories.IOrderRepository, productRepo repositories.IProductRepository) IOrderService {
	orderSvcLock.Do(func() {
		orderSvc = &orderService{
			productRepo: productRepo,
			orderRepo:   orderRepo,
		}
	})

	return orderSvc
}

func (svc *orderService) GetOrderByID(id string) (*dto.Order, error) {
	result, err := svc.orderRepo.FindByIDWithItem(id)
	if err != nil {
		return nil, err
	}
	return mappers.OrderResponse(result), nil
}

func (svc *orderService) GetOrderByUserID(id string) ([]*dto.Order, error) {
	result, err := svc.orderRepo.FindByUserIDWithItem(id)
	if err != nil {
		return nil, err
	}
	return mappers.OrdersResponse(result), nil
}

func (svc *orderService) Pay(userID string, req requests.PayOrder) (*dto.Order, error) {
	return nil, nil
}
