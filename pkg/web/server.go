package web

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mokhan.ca/CodeChica/sparkleapi/pkg/domain"
)

type Server struct {
	Sparkles   *[]domain.Sparkle
	fileserver http.Handler
	logger     *log.Logger
}

func NewServer(sparkles *[]domain.Sparkle) Server {
	if sparkles == nil {
		sparkles = &[]domain.Sparkle{}
	}

	return Server{
		Sparkles:   sparkles,
		fileserver: http.FileServer(http.Dir("public")),
		logger:     log.New(os.Stderr, "", log.LstdFlags|log.LUTC|log.Lshortfile),
	}
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Println(r.Method, r.URL)
	switch r.URL.String() {
	case "/sparkles.json":
		s.LegacyHTTPHandler(w, r)
		break
	case "/v2/sparkles":
		s.SparklesHTTPHandler(w, r)
		break
	default:
		s.fileserver.ServeHTTP(w, r)
	}
}

func renderError(w http.ResponseWriter, c int, e error) {
	w.WriteHeader(c)
	w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, e)))
}
