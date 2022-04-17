package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/codechica/SparkleHub-lite/pkg/db"
	"github.com/codechica/SparkleHub-lite/pkg/domain"
	"github.com/google/jsonapi"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	t.Run("GET /sparkles.json", func(t *testing.T) {
		t.Run("with valid data", func(t *testing.T) {
			sparkle, _ := domain.NewSparkle("@monalisa for helping me with my homework.")
			store := db.NewStorage()
			store.Save(sparkle)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/sparkles.json", nil)
			NewServer(store).ServeHTTP(response, request)

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
			store := db.NewStorage()
			store.Save(sparkle)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/v2/sparkles", nil)
			NewServer(store).ServeHTTP(response, request)

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
			store := db.NewStorage()
			server := NewServer(store)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest(
				"POST",
				"/sparkles.json",
				strings.NewReader(`{"body":"@monalisa for being kind"}`),
			)
			request.Header.Set("Content-Type", "application/json")
			server.ServeHTTP(response, request)

			assert.Equal(t, 201, response.Code)
			assert.Equal(t, 1, len(store.Sparkles))

			var data domain.Sparkle
			err := json.NewDecoder(response.Body).Decode(&data)
			t.Logf("Response: %v", data)

			assert.Equal(t, nil, err)
			assert.Equal(t, "@monalisa", data.Sparklee)
			assert.Equal(t, "for being kind", data.Reason)

			sparkle := *store.Sparkles[0]
			assert.Equal(t, "@monalisa", sparkle.Sparklee)
			assert.Equal(t, "for being kind", sparkle.Reason)
		})

		t.Run("with invalid data", func(t *testing.T) {
			response := httptest.NewRecorder()
			request, _ := http.NewRequest(
				"POST",
				"/sparkles.json",
				strings.NewReader(`{"body":"invalid"}`),
			)
			request.Header.Set("Content-Type", "application/json")
			NewServer(nil).ServeHTTP(response, request)

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

		t.Run("without a Sparklee", func(t *testing.T) {
			out := bytes.NewBuffer(nil)
			sparkle := &domain.Sparkle{Reason: "because"}
			jsonapi.MarshalOnePayloadEmbedded(out, sparkle)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodPost, "/v2/sparkles", out)

			NewServer(nil).ServeHTTP(response, request)

			assert.Equal(t, http.StatusBadRequest, response.Code)
		})

		t.Run("without a Reason", func(t *testing.T) {
			out := bytes.NewBuffer(nil)
			sparkle := &domain.Sparkle{Sparklee: "@monalisa", Reason: ""}
			jsonapi.MarshalOnePayloadEmbedded(out, sparkle)

			response := httptest.NewRecorder()
			request, _ := http.NewRequest(http.MethodPost, "/v2/sparkles", out)

			NewServer(nil).ServeHTTP(response, request)

			assert.Equal(t, http.StatusBadRequest, response.Code)
		})
	})
}
