package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Config struct {
	Port            string
	ShutdownTimeout time.Duration
}

type Server struct {
	cfg    Config
	logger *slog.Logger
	server *grpc.Server
}

func NewGrpcServer(cfg Config, logger *slog.Logger, opts ...grpc.ServerOption) *Server {
	srv := grpc.NewServer(opts...)
	return &Server{cfg: cfg, logger: logger, server: srv}
}

func (s *Server) RegisterService(desc *grpc.ServiceDesc, impl any) {
	s.server.RegisterService(desc, impl)
}

func (s *Server) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":"+s.cfg.Port)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", s.cfg.Port, err)
	}

	errCh := make(chan error, 1)
	go func() {
		s.logger.Info("gRPC server started", "port", s.cfg.Port)
		if err := s.server.Serve(lis); err != nil {
			errCh <- fmt.Errorf("gRPC serve error: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		s.GracefulStop()
		return nil
	case err := <-errCh:
		return err
	}
}

func (s *Server) GracefulStop() {
	s.logger.Info("GRPC server shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
	defer cancel()

	stopped := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		s.logger.Info("GRPC server stopped gracefully")
	case <-shutdownCtx.Done():
		s.logger.Warn("GRPC server shutdown timed out, forcing stop")
		s.server.Stop()
	}
}
