package main

import (
	"go-oauth/controller"
	"go-oauth/middleware"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/auth/signin/google", controller.SignIn)

	mux.HandleFunc("/api/v1/auth/signin/google/callback", controller.SignInCallback)

	protectedMux := http.NewServeMux()

	protectedMux.HandleFunc("/api/v1/users", controller.GetUsers)

	mux.Handle("/api/v1/", &middleware.VerifyToken{Handler: protectedMux})

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
