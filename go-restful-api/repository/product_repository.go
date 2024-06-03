package repository

import (
	"context"
	"database/sql"
	"go-restful-api/entity/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
	FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error)
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
}
