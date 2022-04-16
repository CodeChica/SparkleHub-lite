package web

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codechica/SparkleHub-lite/pkg/db"
)

type Server struct {
	db         *db.Storage
	fileserver http.Handler
	logger     *log.Logger
}

func NewServer(storage *db.Storage) Server {
	if storage == nil {
		storage = db.NewStorage()
	}

	return Server{
		db:         storage,
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
