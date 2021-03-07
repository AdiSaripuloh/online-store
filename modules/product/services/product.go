package services

import (
	"errors"
	"github.com/AdiSaripuloh/online-store/mappers"
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

func (svc *productService) All() ([]*dto.Product, error) {
	results, err := svc.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return mappers.ProductsResponse(results), nil
}

func (svc *productService) IsAvailable(id string, quantity int64) (bool, error) {
	uuID, err := uuid.FromString(id)
	if err != nil {
		return false, errors.New("Failed parsing UUID.")
	}

	product, err := svc.repository.GetQuantityByID(uuID)
	if err != nil {
		return false, err
	}
	if product.Quantity < quantity {
		return false, nil
	}
	return true, nil
}
