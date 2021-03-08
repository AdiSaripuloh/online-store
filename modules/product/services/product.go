package services

import (
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
	uuid "github.com/satori/go.uuid"
	"sync"
)

type productService struct {
	repository repositories.ProductRepository
}

var (
	productSvcLock sync.Once
	productSvc     ProductService
)

func NewProductService(repository repositories.ProductRepository) ProductService {
	productSvcLock.Do(func() {
		productSvc = &productService{
			repository: repository,
		}
	})

	return productSvc
}

func (svc *productService) All() ([]*dto.Product, *responses.HttpError) {
	results, err := svc.repository.FindAll()
	if err != nil {
		return nil, responses.InternalServerError(nil)
	}
	return dto.ProductsResponse(results), nil
}

func (svc *productService) IsAvailable(id string, quantity int64) (bool, *responses.HttpError) {
	uuID, err := uuid.FromString(id)
	if err != nil {
		return false, responses.BadRequest("Failed parsing UUID.", nil)
	}

	product, err := svc.repository.GetQuantityByID(uuID)
	if err != nil {
		return false, responses.InternalServerError(nil)
	}
	if product.Quantity < quantity {
		return false, responses.BadRequest("Out of stock.", nil)
	}
	return true, nil
}
