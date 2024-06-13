package middleware

import (
	"go-oauth/helper"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type VerifyToken struct {
	Handler http.Handler
}

func (verifyToken *VerifyToken) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header.Get("Authorization")

	if authorization == "" {
		helper.ResponseError(w, "Unauthorized", http.StatusUnauthorized, "Token is Required")

		return
	}

	token := strings.Split(authorization, " ")[1]

	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		helper.ResponseError(w, "Unauthorized", http.StatusUnauthorized, "Invalid Token")

		return
	}

	verifyToken.Handler.ServeHTTP(w, r)
}
