package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/jsonapi"
	"mokhan.ca/CodeChica/sparkleapi/pkg/domain"
)

type Server struct {
	sparkles *[]domain.Sparkle
	address  string
}

func NewServer(sparkles *[]domain.Sparkle) Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Server{
		sparkles: sparkles,
		address:  ":" + port,
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	switch r.URL.String() {
	case "/sparkles.json":
		switch r.Method {
		case "GET":
			data, err := json.Marshal(s.sparkles)
			if err == nil {
				w.WriteHeader(http.StatusOK)
				w.Write(data)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
			}
		case "POST":
			var params map[string]string
			err := json.NewDecoder(r.Body).Decode(&params)
			if err == nil {
				sparkle, err := domain.NewSparkle(params["body"])
				if err == nil {
					*s.sparkles = append(*s.sparkles, *sparkle)
					w.WriteHeader(http.StatusCreated)
					json.NewEncoder(w).Encode(sparkle)
				} else {
					renderError(w, http.StatusUnprocessableEntity, err)
				}
			} else {
				renderError(w, http.StatusBadRequest, errors.New("Bad Request"))
			}
		default:
			renderError(w, http.StatusBadRequest, errors.New("Bad Request"))
		}
		break
	default:
		x := []*domain.Sparkle{}
		for _, item := range *s.sparkles {
			x = append(x, &item)
		}
		if err := jsonapi.MarshalPayload(w, x); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Oops",
				Detail: err.Error(),
				Status: http.StatusText(http.StatusInternalServerError),
			}})
		}
	}
}

func renderError(w http.ResponseWriter, c int, e error) {
	w.WriteHeader(c)
	w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, e)))
}
