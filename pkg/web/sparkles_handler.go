package web

import (
	"net/http"

	"github.com/google/jsonapi"
	"mokhan.ca/CodeChica/sparkleapi/pkg/domain"
)

func (s Server) SparklesHTTPHandler(w http.ResponseWriter, r *http.Request) {
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
}
