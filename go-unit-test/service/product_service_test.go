package service

import (
	"go-unit-test/entity"
	"go-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var productRepository = &repository.ProductRepositoryMock{
	Mock: mock.Mock{},
}

var productService = ProductService{
	Repository: productRepository,
}

func TestProductServiceFindByIdSuccess(t *testing.T) {
	product := entity.Product{
		Id:    "12345",
		Name:  "Product One",
		Price: 1000000,
	}

	productRepository.Mock.On("FindById", "12345").Return(product)

	result, err := productService.FindById("12345")

	require.NotNil(t, product)
	require.Equal(t, product.Id, result.Id)
	require.Equal(t, product.Name, result.Name)
	require.Equal(t, product.Price, result.Price)

	require.Nil(t, err)
}

func TestProductServiceFindByIdNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "123").Return(nil)

	result, err := productService.FindById("123")

	require.Nil(t, result)

	require.NotNil(t, err)
}
