package services

import (
	"github.com/AdiSaripuloh/online-store/dto"
	"github.com/AdiSaripuloh/online-store/mappers"
	"github.com/AdiSaripuloh/online-store/repositories"
	"sync"
)

type IProductService interface {
	GetAll() ([]dto.Product, error)
	IsAvailable(id string, quantity int64) (bool, error)
}

type productService struct {
	repository repositories.IProductRepository
}

var (
	productSvcLock sync.Once
	productSvc     IProductService
)

func NewProductService(repository repositories.IProductRepository) IProductService {
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
