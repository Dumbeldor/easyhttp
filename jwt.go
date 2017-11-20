package easyhttp

import (
	"net/http"
	"strings"
)

// swagger:parameters jwtToken
type JWTRawHeader struct {
	// Authorization header for a users bearer token
	// in: header
	// required: true
	Authorization string
}

func GetJWTAuthHeader(r *http.Request) (*JWTRawHeader, *ErrorResponse, int) {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) == 0 {
		er := &ErrorResponse{}
		er.Body.Message = "Authorization header is missing."
		return nil, er, http.StatusBadRequest
	}

	ahSplited := strings.Split(authHeader, " ")
	if len(ahSplited) != 2 {
		er := &ErrorResponse{}
		er.Body.Message = "Malformed Authorization header."
		return nil, er, http.StatusBadRequest
	}

	if ahSplited[0] != "Bearer" {
		er := &ErrorResponse{}
		er.Body.Message = "Authorization header doesn't describe a JWT Authorization."
		return nil, er, http.StatusBadRequest
	}

	return &JWTRawHeader{Authorization: ahSplited[1]}, nil, 0
}

