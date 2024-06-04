package app

import (
	"go-dependency-injection/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: authMiddleware,
	}
}
