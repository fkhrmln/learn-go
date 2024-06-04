//go:build wireinject
// +build wireinject

package injector

import (
	"go-dependency-injection/app"
	"go-dependency-injection/config"
	"go-dependency-injection/controller"
	"go-dependency-injection/helper"
	"go-dependency-injection/middleware"
	"go-dependency-injection/repository"
	"go-dependency-injection/router"
	"go-dependency-injection/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

func InitializeServer() *http.Server {
	wire.Build(
		config.NewDB,
		helper.NewValidator,
		repository.NewProductRepository,
		service.NewProductService,
		controller.NewProductController,
		router.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		app.NewServer,
	)

	return nil
}
