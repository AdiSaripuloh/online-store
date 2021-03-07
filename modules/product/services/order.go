package services

import (
	"errors"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
	"github.com/AdiSaripuloh/online-store/modules/product/requests"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type orderService struct {
	orderRepo   repositories.OrderRepository
	productRepo repositories.ProductRepository
}

var (
	orderSvcLock sync.Once
	orderSvc     OrderService
)

func NewOrderService(orderRepo repositories.OrderRepository, productRepo repositories.ProductRepository) OrderService {
	orderSvcLock.Do(func() {
		orderSvc = &orderService{
			productRepo: productRepo,
			orderRepo:   orderRepo,
		}
	})

	return orderSvc
}

func (svc *orderService) GetOrderByID(id string) (*dto.Order, error) {
	uuID, err := uuid.FromString(id)
	if err != nil {
		return nil, errors.New("Failed parsing UUID.")
	}

	result, err := svc.orderRepo.FindByIDWithItems(uuID)
	if err != nil {
		return nil, err
	}
	return mappers.OrderResponse(result), nil
}

func (svc *orderService) GetOrderByUserID(id string) ([]*dto.Order, error) {
	uuID, err := uuid.FromString(id)
	if err != nil {
		return nil, errors.New("Failed parsing UUID.")
	}

	result, err := svc.orderRepo.FindByUserIDWithItems(uuID)
	if err != nil {
		return nil, err
	}
	return mappers.OrdersResponse(result), nil
}

func (svc *orderService) Pay(id string, userID string, req requests.PayOrder) (*dto.Order, error) {
	uuID, err := uuid.FromString(id)
	if err != nil {
		return nil, errors.New("Failed parsing UUID.")
	}

	order, err := svc.orderRepo.FindByIDWithItems(uuID)
	if err != nil {
		return nil, errors.New("Order not found.")
	}

	if order.Status != models.UNPAID {
		return nil, errors.New("Order status is " + string(order.Status) + ". Contact administrator for more information.")
	}

	if order.GrandTotal != req.Amount {
		return nil, errors.New("Amount doesn't match.")
	}

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return nil, errors.New("Failed parsing UUID.")
	}
	if order.UserID != userUUID {
		return nil, errors.New("Forbidden.")
	}

	for _, item := range order.OrderItems {
		product, err := svc.productRepo.GetQuantityByID(item.ProductID)
		if err != nil {
			return nil, errors.New("Failed get product quantity.")
		}
		if item.Quantity > product.Quantity {
			return nil, errors.New(product.Name + " out of stock.")
		}
		quantity := product.Quantity - item.Quantity
		_, _ = svc.productRepo.Update(&models.Product{
			ID:       item.ProductID,
			Quantity: quantity,
		})
	}

	update, err := svc.orderRepo.Update(&models.Order{
		ID:     order.ID,
		Status: models.PAID,
	})
	if err != nil || !update {
		return nil, errors.New("Failed update status")
	}

	order.Status = models.PAID

	return mappers.OrderResponse(order), nil
}
