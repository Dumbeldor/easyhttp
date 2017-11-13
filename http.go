package easyhttp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	Message string `json:"message"`
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

func GetJWTAuthHeader(r *http.Request) (*string, *ErrorResponse, int) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) == 0 {
		er := &ErrorResponse{Message: "Authorization header is missing."}
		return nil, er, http.StatusBadRequest
	}

	ahSplited := strings.Split(authHeader, " ")
	if len(ahSplited) != 2 {
		er := &ErrorResponse{Message: "Malformed Authorization header."}
		return nil, er, http.StatusBadRequest
	}

	if ahSplited[0] != "JWT" {
		er := &ErrorResponse{Message: "Authorization header doesn't describe a JWT Authorization."}
		return nil, er, http.StatusBadRequest
	}

	return &ahSplited[1], nil, 0
}
