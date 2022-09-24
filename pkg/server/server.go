package server

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/taaaaakahiro/go_gorilla_grpc_sqlboiler/pkg/handler"
	"go.uber.org/zap"
)

type Config struct {
	Log *zap.Logger
}

type Server struct {
	Mux        *http.ServeMux
	Handler    http.Handler
	server     *http.Server
	handler    *handler.Handler
	log        *zap.Logger
	MuxGorilla *mux.Router
}

func NewServer(registry *handler.Handler, cfg *Config) *Server {
	s := &Server{
		// Mux:        http.NewServeMux(),
		handler:    registry,
		MuxGorilla: mux.NewRouter(),
	}
	if cfg != nil {
		if log := cfg.Log; log != nil {
			s.log = log
		}
	}
	s.registerHandler()
	return s
}

func (s *Server) Serve(listener net.Listener) error {
	server := &http.Server{
		Handler: cors.Default().Handler(s.Mux),
	}
	s.server = server
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) GracefulShutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) registerHandler() {
	// rest api
	// net/httpの場合
	// s.Mux.Handle("/user/list", s.handler.V1.GetUsers())
	// s.Mux.Handle("/message/list/", s.handler.V1.GetMessages())

	// common
	// s.Mux.Handle("/healthz", s.healthCheckHandler())
	// s.Mux.Handle("/version", s.handler.Version.GetVersion())
}

func (s *Server) healthCheckHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
