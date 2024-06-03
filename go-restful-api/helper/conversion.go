package helper

import (
	"go-restful-api/entity/domain"
	"go-restful-api/entity/dto"
)

func ToProductResponse(product domain.Product) dto.ProductResponse {
	return dto.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}
}
