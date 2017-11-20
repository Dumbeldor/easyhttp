package easyhttp

import (
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CheckResponseCode(t *testing.T, want, actual int) {
	if want != actual {
		t.Errorf("Handler returned wrong status code: got %v want %v", actual, want)
	}
}

func ExecuteRequest(req *http.Request, handler echo.HandlerFunc, e *echo.Echo) (*httptest.ResponseRecorder, error) {
	rr := httptest.NewRecorder()
	c := e.NewContext(req, rr)
	h := handler(c)
	return rr, h
}
