package web

import (
	"net/http"

	"github.com/codechica/SparkleHub-lite/pkg/domain"
	"github.com/google/jsonapi"
)

func (s Server) SparklesHTTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", jsonapi.MediaType)

	switch r.Method {
	case "GET":
		if err := jsonapi.MarshalPayload(w, s.db.All()); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Oops",
				Detail: err.Error(),
				Status: http.StatusText(http.StatusInternalServerError),
			}})
		}
	case "POST":
		sparkle := new(domain.Sparkle)
		if err := jsonapi.UnmarshalPayload(r.Body, sparkle); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Invalid",
				Detail: err.Error(),
				Status: http.StatusText(http.StatusInternalServerError),
			}})
		} else {
			if err := s.db.Save(sparkle); err != nil {
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
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
