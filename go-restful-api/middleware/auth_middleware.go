package middleware

import (
	"go-restful-api/entity/dto"
	"go-restful-api/helper"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-KEY") != "SECRET" {
		response := dto.Response{
			Status: "Unauthorized",
			Code:   http.StatusUnauthorized,
			Data:   nil,
		}

		w.WriteHeader(http.StatusUnauthorized)

		w.Header().Add("Content-Type", "application/json")

		helper.ResponseWriter(w, response)
	} else {
		middleware.Handler.ServeHTTP(w, r)
	}
}
