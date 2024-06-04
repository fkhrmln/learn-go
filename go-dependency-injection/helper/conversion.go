package helper

import (
	"go-dependency-injection/entity/domain"
	"go-dependency-injection/entity/dto"
)

func ToProductResponse(product domain.Product) dto.ProductResponse {
	return dto.ProductResponse{
		Id:    product.Id,
		Name:  product.Name,
		Price: product.Price,
	}
}
