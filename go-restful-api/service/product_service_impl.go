package service

import (
	"context"
	"database/sql"
	"go-restful-api/entity/domain"
	"go-restful-api/entity/dto"
	"go-restful-api/exception"
	"go-restful-api/helper"
	"go-restful-api/repository"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRespository repository.ProductRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, db *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRespository: productRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request dto.ProductCreateRequest) dto.ProductResponse {
	err := service.Validate.Struct(request)

	helper.PanicError(err)

	tx, err := service.DB.Begin()

	helper.PanicError(err)

	defer helper.Tx(tx)

	product := domain.Product{
		Name:  request.Name,
		Price: request.Price,
	}

	product = service.ProductRespository.Create(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []dto.ProductResponse {
	tx, err := service.DB.Begin()

	helper.PanicError(err)

	defer helper.Tx(tx)

	products := service.ProductRespository.FindAll(ctx, tx)

	productsResponse := []dto.ProductResponse{}

	for _, product := range products {
		productResponse := helper.ToProductResponse(product)

		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) dto.ProductResponse {
	tx, err := service.DB.Begin()

	helper.PanicError(err)

	defer helper.Tx(tx)

	product, err := service.ProductRespository.FindById(ctx, tx, productId)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request dto.ProductUpdateRequest) dto.ProductResponse {
	err := service.Validate.Struct(request)

	helper.PanicError(err)

	tx, err := service.DB.Begin()

	helper.PanicError(err)

	defer helper.Tx(tx)

	_, err = service.ProductRespository.FindById(ctx, tx, request.Id)

	helper.PanicError(err)

	product := domain.Product{
		Id:    request.Id,
		Name:  request.Name,
		Price: request.Price,
	}

	product = service.ProductRespository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) {
	tx, err := service.DB.Begin()

	helper.PanicError(err)

	defer helper.Tx(tx)

	_, err = service.ProductRespository.FindById(ctx, tx, productId)

	helper.PanicError(err)

	product := domain.Product{
		Id: productId,
	}

	service.ProductRespository.Delete(ctx, tx, product)
}
