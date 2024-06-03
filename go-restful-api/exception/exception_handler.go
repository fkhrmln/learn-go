package exception

import (
	"go-restful-api/entity/dto"
	"go-restful-api/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ExceptionHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(w, r, err) {
		return
	}

	if validationError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	response := dto.Response{
		Status: "Internal Server Error",
		Code:   http.StatusInternalServerError,
		Data:   err,
	}

	w.WriteHeader(http.StatusInternalServerError)

	w.Header().Add("Content-Type", "application/json")

	helper.ResponseWriter(w, response)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		response := dto.Response{
			Status: "Not Found",
			Code:   http.StatusNotFound,
			Data:   exception.Error,
		}

		w.WriteHeader(http.StatusNotFound)

		w.Header().Add("Content-Type", "application/json")

		helper.ResponseWriter(w, response)

		return true
	} else {
		return false
	}
}

func validationError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		response := dto.Response{
			Status: "Bad Request",
			Code:   http.StatusBadRequest,
			Data:   exception.Error(),
		}

		w.WriteHeader(http.StatusBadRequest)

		w.Header().Add("Content-Type", "application/json")

		helper.ResponseWriter(w, response)

		return true
	} else {
		return false
	}
}
