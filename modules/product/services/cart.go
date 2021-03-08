package services

import (
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/models"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
	uuid "github.com/satori/go.uuid"
	"log"
	"sync"
)

type cartService struct {
	cartRepo     repositories.CartRepository
	cartItemRepo repositories.CartItemRepository
	productRepo  repositories.ProductRepository
	orderRepo    repositories.OrderRepository
}

var (
	cartSvcLock sync.Once
	cartSvc     CartService
)

func NewCartService(cartRepo repositories.CartRepository, cartItemRepo repositories.CartItemRepository, productRepo repositories.ProductRepository, orderRepo repositories.OrderRepository) CartService {
	cartSvcLock.Do(func() {
		cartSvc = &cartService{
			cartRepo:     cartRepo,
			cartItemRepo: cartItemRepo,
			productRepo:  productRepo,
			orderRepo:    orderRepo,
		}
	})
	return cartSvc
}

func (svc *cartService) GetByUserID(id string) (*dto.Cart, *responses.HttpError) {
	uuID, err := uuid.FromString(id)
	result, err := svc.cartRepo.FindByUserIDWithItems(uuID)
	if err != nil {
		return nil, responses.NotFound(err.Error(), nil)
	}
	return dto.CartResponse(result), nil
}

func (svc *cartService) Create(userID string, req dto.CreateCartRequest) (*dto.Cart, *responses.HttpError) {
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return nil, responses.BadRequest("Failed parsing UUID.", nil)
	}

	exists, err := svc.cartRepo.IsExists(userUUID)
	if err != nil {
		return nil, responses.BadRequest("Can't find current cart.", nil)
	}
	if *exists {
		return nil, responses.BadRequest("Can't create new cart.", nil)
	}

	var items []models.CartItem
	var grandTotal float64
	for _, item := range req.Items {
		product, err := svc.productRepo.FindByID(item.ProductID)
		if err != nil {
			return nil, responses.BadRequest("Some products not found. Please re-check your cart.", nil)
		}
		if item.Quantity > product.Quantity {
			return nil, responses.BadRequest(product.Name+" out of stock.", nil)
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
		return nil, responses.InternalServerError(nil)
	}

	return dto.CartResponse(result), nil
}

func (svc *cartService) Checkout(userID string, req dto.CheckoutCartRequest) (*dto.Order, *responses.HttpError) {
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return nil, responses.BadRequest("Failed parsing UUID.", nil)
	}

	cart, err := svc.cartRepo.FindByUserIDWithItems(userUUID)
	if err != nil {
		return nil, responses.NotFound(err.Error(), nil)
	}

	var orderItems []models.OrderItem
	var grandTotalOrder float64
	var cartItem []*models.CartItem

	for _, item := range req.Items {
		product, err := svc.productRepo.FindByID(item.ProductID)
		log.Println(product)
		if err != nil {
			return nil, responses.BadRequest("Some products not found. Please re-check your cart.", nil)
		}
		if item.Quantity > product.Quantity {
			return nil, responses.BadRequest(product.Name+" out of stock.", nil)
		}

		foundInCart := false
		for _, cItem := range cart.CartItems {
			if cItem.ProductID == item.ProductID {
				foundInCart = true
				restQuantity := cItem.Quantity - item.Quantity
				cartItem = append(cartItem, &models.CartItem{
					ID:        cItem.ID,
					ProductID: cItem.ProductID,
					Quantity:  restQuantity,
				})
				break
			}
		}
		if !foundInCart {
			return nil, responses.BadRequest(product.Name+" not found in cart.", nil)
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
		return nil, responses.InternalServerError(nil)
	}

	grandTotalCart := cart.GrandTotal - grandTotalOrder
	if grandTotalCart > 0 {
		cUpdate, err := svc.cartRepo.Update(&models.Cart{ID: cart.ID, GrandTotal: grandTotalCart})
		if err != nil {
			return nil, responses.BadRequest("Failed update grand total cart "+cart.ID.String(), nil)
		}
		if cUpdate {
			for _, cItem := range cartItem {
				if cItem.Quantity > 0 {
					cIUpdate, err := svc.cartItemRepo.Update(cItem)
					if err != nil || !cIUpdate {
						responses.BadRequest("Failed update quantity cart item "+cItem.ProductID.String(), nil)
					}
				} else {
					cDelete, err := svc.cartItemRepo.Delete(cItem)
					if err != nil || !cDelete {
						responses.BadRequest("Failed delete cart item "+cItem.ProductID.String(), nil)
					}
				}
			}
		}
	} else {
		cDelete, err := svc.cartRepo.Delete(&models.Cart{ID: cart.ID})
		if err != nil || !cDelete {
			responses.BadRequest("Failed delete cart "+cart.ID.String(), nil)
		}
	}

	return dto.OrderResponse(result), nil
}
