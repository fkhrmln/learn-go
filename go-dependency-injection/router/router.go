package router

import (
	"go-dependency-injection/controller"
	"go-dependency-injection/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/products", httprouter.Handle(productController.Create))

	router.GET("/api/v1/products", httprouter.Handle(productController.FindAll))

	router.GET("/api/v1/products/:productId", httprouter.Handle(productController.FindById))

	router.PUT("/api/v1/products/:productId", httprouter.Handle(productController.Update))

	router.DELETE("/api/v1/products/:productId", httprouter.Handle(productController.Delete))

	router.PanicHandler = exception.ExceptionHandler

	return router
}
