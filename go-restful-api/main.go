package main

import (
	"fmt"
	"go-restful-api/config"
	"go-restful-api/controller"
	"go-restful-api/exception"
	"go-restful-api/helper"
	"go-restful-api/middleware"
	"go-restful-api/repository"
	"go-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := config.GetConnection()

	validate := validator.New()

	router := httprouter.New()

	productRepository := repository.NewProductRepository()

	productService := service.NewProductService(productRepository, db, validate)

	productController := controller.NewProductController(productService)

	router.POST("/api/v1/products", httprouter.Handle(productController.Create))

	router.GET("/api/v1/products", httprouter.Handle(productController.FindAll))

	router.GET("/api/v1/products/:productId", httprouter.Handle(productController.FindById))

	router.PUT("/api/v1/products/:productId", httprouter.Handle(productController.Update))

	router.DELETE("/api/v1/products/:productId", httprouter.Handle(productController.Delete))

	router.PanicHandler = exception.ExceptionHandler

	authMiddleware := middleware.AuthMiddleware{
		Handler: router,
	}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &authMiddleware,
	}

	err := server.ListenAndServe()

	helper.PanicError(err)

	fmt.Println("Server is running on", server.Addr)
}
