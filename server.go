package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	sparkles *[]Sparkle
}

func NewServer(sparkles *[]Sparkle) Server {
	return Server{
		sparkles: sparkles,
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
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
			sparkle, err := NewSparkle(params["body"])
			if err == nil {
				*s.sparkles = append(*s.sparkles, *sparkle)
				w.WriteHeader(http.StatusCreated)
				data, err := json.Marshal(sparkle)
				if err == nil {
					w.Write(data)
				} else {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
				}
			} else {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
			}
		} else {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
		}
	default:
		w.Write([]byte(`{"error":"unknown"}`))
	}
}
