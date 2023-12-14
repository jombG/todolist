package server

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"todolist/config"
	"todolist/internal/server/generated"
	"todolist/internal/usecase"
)

var _ generated.ServerInterface = (*Server)(nil)

type (
	Server struct {
		log      *zap.Logger
		cfg      *config.Config
		useCases *usecase.UseCases
	}

	OpenAPI struct {
		*openapi3.T
		routers.Router
	}
)

func New(
	log *zap.Logger,
	cfg *config.Config,
	useCases *usecase.UseCases,
) *Server {
	return &Server{
		log:      log,
		cfg:      cfg,
		useCases: useCases,
	}
}

func (s *Server) NewServerOptions() generated.ChiServerOptions {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)

	router.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metrics.WritePrometheus(w, true)
	})

	return generated.ChiServerOptions{
		BaseRouter:  router,
		Middlewares: []generated.MiddlewareFunc{},
	}
}
