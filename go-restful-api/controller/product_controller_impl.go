package controller

import (
	"go-restful-api/entity/dto"
	"go-restful-api/helper"
	"go-restful-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productCreateRequest := dto.ProductCreateRequest{}

	helper.RequestReader(r, &productCreateRequest)

	productResponse := controller.ProductService.Create(r.Context(), productCreateRequest)

	response := dto.Response{
		Status: "Created",
		Code:   http.StatusCreated,
		Data:   productResponse,
	}

	w.WriteHeader(http.StatusCreated)

	helper.ResponseWriter(w, response)
}

func (controller *ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productsResponse := controller.ProductService.FindAll(r.Context())

	response := dto.Response{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   productsResponse,
	}

	w.WriteHeader(http.StatusOK)

	helper.ResponseWriter(w, response)
}

func (controller *ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")

	productResponse := controller.ProductService.FindById(r.Context(), productId)

	response := dto.Response{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   productResponse,
	}

	w.WriteHeader(http.StatusOK)

	helper.ResponseWriter(w, response)
}

func (controller *ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")

	productUpdateRequest := dto.ProductUpdateRequest{}

	productUpdateRequest.Id = productId

	helper.RequestReader(r, &productUpdateRequest)

	productResponse := controller.ProductService.Update(r.Context(), productUpdateRequest)

	response := dto.Response{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   productResponse,
	}

	w.WriteHeader(http.StatusOK)

	helper.ResponseWriter(w, response)
}

func (controller *ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")

	controller.ProductService.Delete(r.Context(), productId)

	response := dto.Response{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   nil,
	}

	w.WriteHeader(http.StatusOK)

	helper.ResponseWriter(w, response)
}
