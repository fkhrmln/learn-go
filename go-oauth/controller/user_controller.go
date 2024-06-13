package controller

import (
	"context"
	"encoding/json"
	"go-oauth/config"
	"go-oauth/dto/response"
	"go-oauth/entity"
	"go-oauth/helper"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := config.GetConnection()

	rows, err := db.QueryContext(context.Background(), "SELECT id, email, name, picture FROM users;")

	if err != nil {
		helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

		return
	}

	users := []entity.User{}

	for rows.Next() {
		user := entity.User{}

		if err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Picture); err != nil {
			helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

			return
		}

		users = append(users, user)
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	response := response.Response{
		Status:  "OK",
		Code:    http.StatusOK,
		Message: "",
		Data:    users,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		helper.ResponseError(w, "Internal Server Error", http.StatusInternalServerError, err.Error())

		return
	}
}
