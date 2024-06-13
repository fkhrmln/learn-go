package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"go-oauth/config"
	"go-oauth/dto"
	"go-oauth/dto/response"
	"go-oauth/entity"
	"go-oauth/helper"
	"net/http"
	"os"

	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	oAuthConfig := config.OAuthConfig()

	url := oAuthConfig.AuthCodeURL(os.Getenv("STATE"))

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func SignInCallback(w http.ResponseWriter, r *http.Request) {
	db := config.GetConnection()

	oAuthConfig := config.OAuthConfig()

	if r.URL.Query().Get("state") != os.Getenv("STATE") {
		helper.ResponseError(w, "Bad Request", http.StatusBadRequest, "Invalid State")

		return
	}

	code := r.URL.Query().Get("code")

	token, err := oAuthConfig.Exchange(context.Background(), code)

	if err != nil {
		helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

		return
	}

	client := oAuthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

		return
	}

	defer resp.Body.Close()

	userInfo := dto.UserInfo{}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

		return
	}

	user := entity.User{}

	err = db.QueryRowContext(context.Background(), "SELECT id, email, name, picture FROM users WHERE email = ? LIMIT 1;", userInfo.Email).Scan(&user.Id, &user.Email, &user.Name, &user.Picture)

	if err != nil {
		if err == sql.ErrNoRows {
			id := uuid.NewString()

			_, err := db.ExecContext(context.Background(), "INSERT INTO users (id, email, name, picture) VALUES (?, ?, ?, ?);", id, userInfo.Email, userInfo.Name, userInfo.Picture)

			if err != nil {
				helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

				return
			}
		} else {
			helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

			return
		}
	}

	db.ExecContext(context.Background(), "UPDATE users SET name = ?, picture = ? WHERE email = ?;", userInfo.Name, userInfo.Picture, userInfo.Email)

	t, err := helper.GenerateToken(user)

	if err != nil {
		helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

		return
	}

	userInfo.AccessToken = t

	cookie := http.Cookie{}
	cookie.Name = "JWT"
	cookie.Value = t
	cookie.Path = "/"

	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	response := response.Response{
		Status:  "OK",
		Code:    http.StatusOK,
		Message: "Sign In Successfully",
		Data:    userInfo,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

		return
	}
}
