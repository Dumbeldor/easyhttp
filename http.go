package easyhttp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// swagger:response ErrorResponse
type ErrorResponse struct {
	// Error message
	// in: body
	// required: true
	Message string `json:"message,required"`
}

func ReadJsonRequest(payload io.ReadCloser, decodedPayload interface{}) bool {
	decoder := json.NewDecoder(payload)
	defer payload.Close()
	if err := decoder.Decode(&decodedPayload); err != nil {
		return false
	}

	return true
}
func WriteHTTPJsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != err {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
