package services

import (
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
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

func (svc *orderService) GetOrderByID(id string) (*dto.Order, *responses.HttpError) {
	uuID, err := uuid.FromString(id)
	if err != nil {
		return nil, responses.BadRequest("Failed parsing UUID.", nil)
	}

	result, err := svc.orderRepo.FindByIDWithItems(uuID)
	if err != nil {
		return nil, responses.InternalServerError(nil)
	}
	return dto.OrderResponse(result), nil
}

func (svc *orderService) GetOrderByUserID(id string) ([]*dto.Order, *responses.HttpError) {
	uuID, err := uuid.FromString(id)
	if err != nil {
		return nil, responses.BadRequest("Failed parsing UUID.", nil)
	}

	result, err := svc.orderRepo.FindByUserIDWithItems(uuID)
	if err != nil {
		return nil, responses.NotFound(err.Error(), nil)
	}
	return dto.OrdersResponse(result), nil
}

func (svc *orderService) Pay(id string, userID string, req dto.PayOrderRequest) (*dto.Order, *responses.HttpError) {
	uuID, err := uuid.FromString(id)
	if err != nil {
		return nil, responses.BadRequest("Failed parsing UUID.", nil)
	}

	order, err := svc.orderRepo.FindByIDWithItems(uuID)
	if err != nil {
		return nil, responses.NotFound("Order not found.", nil)
	}

	if order.Status != models.UNPAID {
		return nil, responses.BadRequest("Order status is "+string(order.Status)+". Contact administrator for more information.", nil)
	}

	if order.GrandTotal != req.Amount {
		return nil, responses.BadRequest("Amount doesn't match.", nil)
	}

	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return nil, responses.BadRequest("Failed parsing UUID.", nil)
	}
	if order.UserID != userUUID {
		return nil, responses.Forbidden("User ID not match", nil)
	}

	for _, item := range order.OrderItems {
		product, err := svc.productRepo.GetQuantityByID(item.ProductID)
		if err != nil {
			return nil, responses.BadRequest("Failed get product quantity.", nil)
		}
		if item.Quantity > product.Quantity {
			return nil, responses.BadRequest(product.Name+" out of stock.", nil)
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
		return nil, responses.InternalServerError(nil)
	}

	order.Status = models.PAID

	return dto.OrderResponse(order), nil
}
