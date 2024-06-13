package helper

import (
	"go-oauth/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user entity.User) (string, error) {
	claims := jwt.MapClaims{
		"id":      user.Id,
		"email":   user.Email,
		"name":    user.Name,
		"picture": user.Picture,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return token, nil
}
