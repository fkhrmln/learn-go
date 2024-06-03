package service

import (
	"context"
	"go-restful-api/entity/dto"
)

type ProductService interface {
	Create(ctx context.Context, request dto.ProductCreateRequest) dto.ProductResponse
	FindAll(ctx context.Context) []dto.ProductResponse
	FindById(ctx context.Context, productId string) dto.ProductResponse
	Update(ctx context.Context, request dto.ProductUpdateRequest) dto.ProductResponse
	Delete(ctx context.Context, productId string)
}
