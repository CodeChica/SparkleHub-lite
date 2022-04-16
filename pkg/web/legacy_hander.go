package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/codechica/SparkleHub-lite/pkg/domain"
)

func (s Server) LegacyHTTPHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		data, err := json.Marshal(s.db.Sparkles)
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
				s.db.Save(sparkle)
				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(sparkle)
			} else {
				renderError(w, http.StatusUnprocessableEntity, err)
			}
		} else {
			renderError(w, http.StatusBadRequest, errors.New("Bad Request"))
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
