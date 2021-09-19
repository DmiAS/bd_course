package http

import (
	"context"
	"net/http"

	"github.com/DmiAS/bd_course/internal/app/config"
)

type Server struct {
	srv *http.Server
}

func NewServer(handler http.Handler, cfg config.Config) *Server {
	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + cfg.HTTP.Port,
	}
	return &Server{srv: srv}
}

func (s *Server) Start() error {
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}
