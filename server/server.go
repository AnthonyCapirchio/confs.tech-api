package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"squidward.confs.tech/conference"
	"squidward.confs.tech/handlers"
)

type Server struct {
	conferencesService *conference.Store
	httpServer         *http.Server
}

func (s *Server) Start() error {
	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.httpServer.Shutdown(ctx)
	return err
}

func NewServer(port string, conferencesService *conference.Store) *Server {
	return &Server{
		conferencesService: conferencesService,
		httpServer:         &http.Server{Addr: port, Handler: getHttpHandler(conferencesService)},
	}
}

func Liveness(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func Readiness(response http.ResponseWriter, _ *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func getHttpHandler(conferencesService *conference.Store) http.Handler {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	r.Path("/conferences/{year}").Handler(handlers.GetConferencesHandler(conferencesService))
	r.Path("/cities").Handler(handlers.GetCitiesHandler(conferencesService))
	r.Path("/countries").Handler(handlers.GetCountriesHandler(conferencesService))
	r.Path("/categories").Handler(handlers.GetCategoriesHandler(conferencesService))

	r.Path("/metrics").Handler(promhttp.Handler())
	r.Path("/liveness").HandlerFunc(Liveness)
	r.Path("/readiness").HandlerFunc(Readiness)

	return r
}
