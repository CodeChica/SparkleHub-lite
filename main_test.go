package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const NoSparklesMessage string = "No sparkles yet"

func TestServer(t *testing.T) {
	t.Run("GET /sparkles.html", func(t *testing.T) {
		t.Run("when there are no sparkles", func(t *testing.T) {
			sparkles := []Sparkle{}
			router := setupRouter(&sparkles)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/sparkles.html", nil)
			router.ServeHTTP(response, request)

			assert.Equal(t, 200, response.Code)
			assert.Contains(t, response.Body.String(), NoSparklesMessage)
		})

		t.Run("when there is a single sparkle", func(t *testing.T) {
			sparkles := []Sparkle{
				{
					Sparklee: "@monalisa",
					Reason:   "for being kind",
				},
			}
			router := setupRouter(&sparkles)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/sparkles.html", nil)
			router.ServeHTTP(response, request)

			assert.Equal(t, 200, response.Code)
			assert.NotContains(t, response.Body.String(), NoSparklesMessage)
			assert.Contains(t, response.Body.String(), "for being kind")
		})
	})

	t.Run("POST /sparkles", func(t *testing.T) {
		t.Run("with valid data", func(t *testing.T) {
			sparkles := []Sparkle{}
			router := setupRouter(&sparkles)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/sparkles", strings.NewReader("body=@monalisa+for+being+kind"))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(response, request)

			assert.Equal(t, 302, response.Code)
			assert.Equal(t, 1, len(sparkles))

			sparkle := sparkles[0]
			assert.Equal(t, "@monalisa", sparkle.Sparklee)
			assert.Equal(t, "for being kind", sparkle.Reason)
		})

		t.Run("with invalid data", func(t *testing.T) {
			sparkles := []Sparkle{}
			router := setupRouter(&sparkles)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest("POST", "/sparkles", strings.NewReader("invalid"))
			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			router.ServeHTTP(response, request)

			assert.Equal(t, 422, response.Code)
		})
	})
}
