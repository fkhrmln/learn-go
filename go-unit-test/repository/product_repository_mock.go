package repository

import (
	"go-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (productRepositoryMock *ProductRepositoryMock) FindById(id string) *entity.Product {
	arguments := productRepositoryMock.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	product := arguments.Get(0).(entity.Product)

	return &product
}
