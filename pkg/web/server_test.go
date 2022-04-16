package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/google/jsonapi"
	"github.com/stretchr/testify/assert"
	"mokhan.ca/CodeChica/sparkleapi/pkg/domain"
)

func TestServer(t *testing.T) {
	t.Run("GET /sparkles.json", func(t *testing.T) {
		t.Run("with valid data", func(t *testing.T) {
			sparkles := []domain.Sparkle{{Sparklee: "@monalisa", Reason: "for helping me with my homework."}}

			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/sparkles.json", nil)
			NewServer(&sparkles).ServeHTTP(response, request)

			assert.Equal(t, http.StatusOK, response.Code)
			var got []domain.Sparkle
			assert.Nil(t, json.NewDecoder(response.Body).Decode(&got))
			assert.Equal(t, 1, len(got))
			assert.Equal(t, "@monalisa", got[0].Sparklee)
			assert.Equal(t, "for helping me with my homework.", got[0].Reason)
		})
	})

	t.Run("GET /v2/sparkles", func(t *testing.T) {
		t.Run("returns the list of sparkles", func(t *testing.T) {
			sparkle, _ := domain.NewSparkle("@mona for helping me")
			sparkles := []domain.Sparkle{*sparkle}

			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/v2/sparkles", nil)
			NewServer(&sparkles).ServeHTTP(response, request)

			assert.Equal(t, http.StatusOK, response.Code)

			got, err := jsonapi.UnmarshalManyPayload(response.Body, reflect.TypeOf(new(domain.Sparkle)))
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, 1, len(got))
			assert.Equal(t, "@mona", got[0].(*domain.Sparkle).Sparklee)
			assert.Equal(t, "for helping me", got[0].(*domain.Sparkle).Reason)
		})
	})

	t.Run("POST /sparkles.json", func(t *testing.T) {
		t.Run("with valid data", func(t *testing.T) {
			sparkles := []domain.Sparkle{}
			server := NewServer(&sparkles)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest(
				"POST",
				"/sparkles.json",
				strings.NewReader(`{"body":"@monalisa for being kind"}`),
			)
			request.Header.Set("Content-Type", "application/json")
			server.ServeHTTP(response, request)

			assert.Equal(t, 201, response.Code)
			assert.Equal(t, 1, len(sparkles))

			var data domain.Sparkle
			err := json.NewDecoder(response.Body).Decode(&data)
			t.Logf("Response: %v", data)

			assert.Equal(t, nil, err)
			assert.Equal(t, "@monalisa", data.Sparklee)
			assert.Equal(t, "for being kind", data.Reason)

			sparkle := sparkles[0]
			assert.Equal(t, "@monalisa", sparkle.Sparklee)
			assert.Equal(t, "for being kind", sparkle.Reason)
		})

		t.Run("with invalid data", func(t *testing.T) {
			sparkles := []domain.Sparkle{}
			server := NewServer(&sparkles)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest(
				"POST",
				"/sparkles.json",
				strings.NewReader(`{"body":"invalid"}`),
			)
			request.Header.Set("Content-Type", "application/json")
			server.ServeHTTP(response, request)

			assert.Equal(t, 422, response.Code)
			var got map[string]string
			assert.Nil(t, json.NewDecoder(response.Body).Decode(&got))
			assert.NotEmpty(t, got["error"])
		})
	})

	t.Run("POST /v2/sparkles", func(t *testing.T) {
		t.Run("with valid data", func(t *testing.T) {
			out := bytes.NewBuffer(nil)
			sparkle, _ := domain.NewSparkle("@mona for the things")
			jsonapi.MarshalOnePayloadEmbedded(out, sparkle)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodPost, "/v2/sparkles", out)
			request.Header.Set("Content-Type", "application/json")

			NewServer(nil).ServeHTTP(response, request)

			assert.Equal(t, http.StatusCreated, response.Code)

			data := new(domain.Sparkle)
			assert.Nil(t, jsonapi.UnmarshalPayload(response.Body, data))
			assert.Equal(t, "@mona", data.Sparklee)
			assert.Equal(t, "for the things", data.Reason)
		})
	})
}
