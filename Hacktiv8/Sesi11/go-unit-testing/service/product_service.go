package service

import (
	"errors"
	"go-unit-testing/entity"
	"go-unit-testing/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id string) (*entity.Product, error) {
	product := service.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}
