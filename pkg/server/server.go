package server

import (
	"context"
	"encoding/json"
	"log"
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

func NewServer(ctx context.Context, registry *handler.Handler, cfg *Config) *Server {
	s := &Server{
		Mux:        http.NewServeMux(),
		handler:    registry,
		MuxGorilla: mux.NewRouter(),
	}
	if cfg != nil {
		if log := cfg.Log; log != nil {
			s.log = log
		}
	}
	s.registerHandler(ctx)
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

func (s *Server) registerHandler(ctx context.Context) {
	// rest api × gorilla
	s.MuxGorilla.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	s.MuxGorilla.HandleFunc("/user/list", s.handler.V1.GetUsers).Methods("GET")   // GET Only Case1
	s.MuxGorilla.HandleFunc("/review/list", s.handler.V1.GetReviews).GetMethods() // GET Only Case2

	// common
	s.MuxGorilla.HandleFunc("/healthz", s.healthCheckHandler)
	s.MuxGorilla.HandleFunc("/version", s.handler.Version.GetVersion)
}

func (s *Server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	b, err := json.Marshal(http.StatusOK)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}
