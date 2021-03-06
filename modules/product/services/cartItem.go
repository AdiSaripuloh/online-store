package services

import (
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type cartItemService struct {
	cartItemRepo repositories.CartItemRepository
}

var (
	cartItemSvcLock sync.Once
	cartItemSvc     CartItemService
)

func NewCartItemService(cartItemRepo repositories.CartItemRepository) CartItemService {
	cartItemSvcLock.Do(func() {
		cartItemSvc = &cartItemService{
			cartItemRepo: cartItemRepo,
		}
	})
	return cartItemSvc
}

func (svc *cartItemService) GetByID(id uuid.UUID) (*models.CartItem, *responses.HttpError) {
	result, err := svc.cartItemRepo.FindByID(id)
	if err != nil {
		return nil, responses.InternalServerError(err.Error())
	}
	return result, nil
}

func (svc *cartItemService) UpdateQuantityByID(id uuid.UUID, quantity int64) (bool, *responses.HttpError) {
	result, err := svc.cartItemRepo.Update(&models.CartItem{
		ID:       id,
		Quantity: quantity,
	})
	if err != nil {
		return false, responses.InternalServerError(err.Error())
	}
	return result, nil
}

func (svc *cartItemService) DeleteByID(id uuid.UUID) (bool, *responses.HttpError) {
	result, err := svc.cartItemRepo.Delete(&models.CartItem{
		ID: id,
	})
	if err != nil {
		return false, responses.InternalServerError(err.Error())
	}
	return result, nil
}
