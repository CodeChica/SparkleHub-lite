package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const NoSparklesMessage string = "No sparkles yet"

func TestWhenThereAreNoSparkles(t *testing.T) {
	sparkles := []Sparkle{}
	router := setupRouter(&sparkles)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/sparkles.html", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
	assert.Contains(t, response.Body.String(), NoSparklesMessage)
}

func TestWhenOneSparkleIsCreated(t *testing.T) {
	sparkles := []Sparkle{}
	router := setupRouter(&sparkles)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/sparkles", strings.NewReader("body=@monalisa+for+being+kind"))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(response, request)

	assert.Equal(t, 302, response.Code)

	response = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/sparkles.html", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
	assert.NotContains(t, response.Body.String(), NoSparklesMessage)
	assert.Contains(t, response.Body.String(), "for being kind")
}
