package web

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
	Sparkles   *[]domain.Sparkle
	Address    string
	fileserver http.Handler
}

func NewServer(sparkles *[]domain.Sparkle) Server {
	if sparkles == nil {
		sparkles = &[]domain.Sparkle{}
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Server{
		Sparkles:   sparkles,
		Address:    ":" + port,
		fileserver: http.FileServer(http.Dir("public")),
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL)
	switch r.URL.String() {
	case "/sparkles.json":
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case "GET":
			data, err := json.Marshal(s.Sparkles)
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
					*s.Sparkles = append(*s.Sparkles, *sparkle)
					w.WriteHeader(http.StatusCreated)
					json.NewEncoder(w).Encode(sparkle)
				} else {
					renderError(w, http.StatusUnprocessableEntity, err)
				}
			} else {
				renderError(w, http.StatusBadRequest, errors.New("Bad Request"))
			}
		default:
			// renderError(w, http.StatusBadRequest, errors.New("Bad Request"))
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		break
	case "/v2/sparkles":
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", jsonapi.MediaType)

		switch r.Method {
		case "GET":
			x := []*domain.Sparkle{}
			for _, item := range *s.Sparkles {
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
		case "POST":
			w.WriteHeader(http.StatusCreated)
			sparkle := new(domain.Sparkle)
			if err := jsonapi.UnmarshalPayload(r.Body, sparkle); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
					Title:  "Invalid",
					Detail: err.Error(),
					Status: http.StatusText(http.StatusInternalServerError),
				}})
			} else {
				w.WriteHeader(http.StatusCreated)
				if err := jsonapi.MarshalPayload(w, sparkle); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
						Title:  "Oops",
						Detail: err.Error(),
						Status: http.StatusText(http.StatusInternalServerError),
					}})
				}
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	default:
		s.fileserver.ServeHTTP(w, r)
	}
}

func renderError(w http.ResponseWriter, c int, e error) {
	w.WriteHeader(c)
	w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, e)))
}
