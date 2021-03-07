package services

import (
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/modules/product/dto"
	"github.com/AdiSaripuloh/online-store/modules/product/repositories"
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

func (svc *productService) GetAll() ([]dto.Product, error) {
	results, err := svc.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return mappers.ProductsResponse(results), nil
}

func (svc *productService) IsAvailable(id string, quantity int64) (bool, error) {
	product, err := svc.repository.GetQuantityByID(id)
	if err != nil {
		return false, err
	}
	if product.Quantity < quantity {
		return false, nil
	}
	return true, nil
}
