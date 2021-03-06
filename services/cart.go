package services

import (
	"errors"
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/models"
	"github.com/AdiSaripuloh/online-store/repositories"
	"github.com/AdiSaripuloh/online-store/requests"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type ICartService interface {
	GetCartByUserID(id string) (*dto.Cart, error)
	Create(userID string, req requests.CreateCart) (*dto.Cart, error)
}

type cartService struct {
	repository  repositories.ICartRepository
	productRepo repositories.IProductRepository
}

var (
	cartSvcLock sync.Once
	cartSvc     ICartService
)

func NewCartService(repository repositories.ICartRepository, productRepo repositories.IProductRepository) ICartService {
	cartSvcLock.Do(func() {
		cartSvc = &cartService{
			repository:  repository,
			productRepo: productRepo,
		}
	})

	return cartSvc
}

func (svc *cartService) GetCartByUserID(id string) (*dto.Cart, error) {
	result, err := svc.repository.FindByUserIDWithItem(id)
	if err != nil {
		return nil, err
	}
	return mappers.CartResponse(result), nil
}

func (svc *cartService) Create(userID string, req requests.CreateCart) (*dto.Cart, error) {
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return nil, errors.New("Failed parsing UUID.")
	}
	exists, err := svc.repository.IsExists(userID)
	if err != nil {
		return nil, errors.New("Can't find current cart.")
	}
	if *exists {
		return nil, errors.New("Can't create new cart.")
	}

	var items []models.CartItem
	var grandTotal float64
	for _, item := range req.Items {
		product, err := svc.productRepo.FindByID(item.ProductID.String())
		if err != nil {
			return nil, errors.New("Some products not found. Please re-check your cart.")
		}
		if item.Quantity > product.Quantity {
			return nil, errors.New(product.Name + " out of stock.")
		}
		grandTotal += product.Price * float64(item.Quantity)
		items = append(items, models.CartItem{
			ProductID: product.ID,
			Quantity:  product.Quantity,
		})
	}

	cart := models.Cart{
		UserID:     userUUID,
		GrandTotal: grandTotal,
		CartItems:  items,
	}
	result, err := svc.repository.Create(cart)
	if err != nil {
		return nil, err
	}

	return mappers.CartResponse(result), nil
}
