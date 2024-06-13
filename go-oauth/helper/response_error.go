package helper

import (
	"encoding/json"
	"go-oauth/dto/response"
	"net/http"
)

func ResponseError(w http.ResponseWriter, status string, code int, message string) {
	response := response.Response{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    nil,
	}

	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	encoder.SetIndent("", "    ")

	encoder.Encode(response)
}
