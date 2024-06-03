package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-restful-api/entity/domain"
	"go-restful-api/helper"

	"github.com/google/uuid"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	query := "INSERT INTO product (id, name, price) VALUES (?, ?, ?);"

	uuid := uuid.NewString()

	_, err := tx.ExecContext(ctx, query, uuid, product.Name, product.Price)

	helper.PanicError(err)

	product.Id = uuid

	return product
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	query := "SELECT id, name, price FROM product;"

	rows, err := tx.QueryContext(ctx, query)

	helper.PanicError(err)

	defer rows.Close()

	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{}

		err := rows.Scan(&product.Id, &product.Name, &product.Price)

		helper.PanicError(err)

		products = append(products, product)
	}

	return products
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId string) (domain.Product, error) {
	query := "SELECT id, name, price FROM product WHERE id = ?;"

	rows, err := tx.QueryContext(ctx, query, productId)

	helper.PanicError(err)

	defer rows.Close()

	product := domain.Product{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price)

		helper.PanicError(err)

		return product, nil
	}

	return product, errors.New("Product Not Found")
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	query := "UPDATE product SET name = ?, price = ? WHERE id = ?;"

	_, err := tx.ExecContext(ctx, query, product.Name, product.Price, product.Id)

	helper.PanicError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	query := "DELETE FROM product WHERE id = ?;"

	_, err := tx.ExecContext(ctx, query, product.Id)

	helper.PanicError(err)
}
