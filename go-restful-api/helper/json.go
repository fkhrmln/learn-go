package helper

import (
	"encoding/json"
	"net/http"
)

func RequestReader(r *http.Request, request interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(request)

	PanicError(err)
}

func ResponseWriter(w http.ResponseWriter, response interface{}) {
	encoder := json.NewEncoder(w)

	encoder.SetIndent("", "    ")

	err := encoder.Encode(response)

	PanicError(err)

	w.Header().Add("Content-Type", "application/json")
}
