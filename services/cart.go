package services

import (
	"errors"
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/models"
	"github.com/AdiSaripuloh/online-store/repositories"
	"github.com/AdiSaripuloh/online-store/requests"
	uuid "github.com/satori/go.uuid"
	"log"
	"sync"
)

type ICartService interface {
	GetCartByUserID(id string) (*dto.Cart, error)
	Create(userID string, req requests.CreateCart) (*dto.Cart, error)
	Checkout(userID string, req requests.Checkout) (*dto.Order, error)
}

type cartService struct {
	cartRepo    repositories.ICartRepository
	productRepo repositories.IProductRepository
	orderRepo   repositories.IOrderRepository
}

var (
	cartSvcLock sync.Once
	cartSvc     ICartService
)

func NewCartService(cartRepo repositories.ICartRepository, productRepo repositories.IProductRepository, orderRepo repositories.IOrderRepository) ICartService {
	cartSvcLock.Do(func() {
		cartSvc = &cartService{
			cartRepo:    cartRepo,
			productRepo: productRepo,
			orderRepo:   orderRepo,
		}
	})

	return cartSvc
}

func (svc *cartService) GetCartByUserID(id string) (*dto.Cart, error) {
	result, err := svc.cartRepo.FindByUserIDWithItem(id)
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

	exists, err := svc.cartRepo.IsExists(userID)
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
			Quantity:  item.Quantity,
		})
	}

	cart := models.Cart{
		UserID:     userUUID,
		GrandTotal: grandTotal,
		CartItems:  items,
	}
	result, err := svc.cartRepo.Create(cart)
	if err != nil {
		return nil, errors.New("Failed create cart")
	}

	return mappers.CartResponse(result), nil
}

func (svc *cartService) Checkout(userID string, req requests.Checkout) (*dto.Order, error) {
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return nil, errors.New("Failed parsing UUID.")
	}

	cart, err := svc.cartRepo.FindByUserIDWithItem(userID)
	if err != nil {
		return nil, errors.New("Cart not found.")
	}

	var orderItems []models.OrderItem
	var grandTotalOrder float64
	var cartItem []models.CartItem

	for _, item := range req.Items {
		product, err := svc.productRepo.FindByID(item.ProductID.String())
		log.Println(product)
		if err != nil {
			return nil, errors.New("Some products not found. Please re-check your cart.")
		}
		if item.Quantity > product.Quantity {
			return nil, errors.New(product.Name + " out of stock.")
		}

		foundInCart := false
		for _, cItem := range cart.CartItems {
			if cItem.ProductID == item.ProductID {
				foundInCart = true
				restQuantity := cItem.Quantity - item.Quantity
				cartItem = append(cartItem, models.CartItem{
					ID:        cItem.ID,
					ProductID: cItem.ProductID,
					Quantity:  restQuantity,
				})
				break
			}
		}
		if !foundInCart {
			return nil, errors.New(product.Name + " not found in cart.")
		}

		grandTotalOrder += product.Price * float64(item.Quantity)
		orderItems = append(orderItems, models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
		})
	}

	order := models.Order{
		UserID:     userUUID,
		GrandTotal: grandTotalOrder,
		Status:     models.UNPAID,
		OrderItems: orderItems,
	}
	result, err := svc.orderRepo.Create(order)
	if err != nil {
		return nil, errors.New("Failed create order")
	}

	grandTotalCart := cart.GrandTotal - grandTotalOrder
	if grandTotalCart > 0 {
		cUpdate, err := svc.cartRepo.UpdateGrandTotalByID(cart.ID, grandTotalCart)
		if err != nil {
			return nil, errors.New("Failed update grand total cart " + cart.ID.String())
		}
		if cUpdate {
			for _, cItem := range cartItem {
				if cItem.Quantity > 0 {
					cIUpdate, err := svc.cartRepo.UpdateQtyCartItemByID(cItem.ID, cItem.Quantity)
					if err != nil || !cIUpdate {
						return nil, errors.New("Failed update quantity cart item " + cItem.ProductID.String())
					}
				} else {
					cDelete, err := svc.cartRepo.DeleteCartItemByID(cItem.ID)
					if err != nil || !cDelete {
						return nil, errors.New("Failed delete cart item " + cItem.ProductID.String())
					}
				}
			}
		}
	} else {
		cDelete, err := svc.cartRepo.DeleteByID(cart.ID)
		if err != nil || !cDelete {
			return nil, errors.New("Failed delete cart " + cart.ID.String())
		}
	}

	return mappers.OrderResponse(result), nil
}
