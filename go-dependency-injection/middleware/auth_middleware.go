package middleware

import (
	"go-dependency-injection/entity/dto"
	"go-dependency-injection/helper"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
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
