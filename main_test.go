package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	router := setupRouter()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
	assert.Contains(t, response.Body.String(), "No sparkles yet")
}
