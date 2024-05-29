package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (productService ProductService) FindById(id string) (*entity.Product, error) {
	product := productService.Repository.FindById(id)

	if product == nil {
		return nil, errors.New("Product Not Found")
	}

	return product, nil
}
