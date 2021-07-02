package grpc

import (
	"context"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	root       *grpc.Server
	cfg        ServerConfig
	stopSignal chan struct{}
}

type ServerConfig struct {
	Network Network
	Host    string // e.g. `:8080`, `127.0.0.1:12345`
}

func NewServer(cfg ServerConfig, logger *zap.Logger, checkAccess grpc.UnaryServerInterceptor) *Server {
	interceptors := makeDefaultInterceptors(logger)
	interceptors = append(interceptors, checkAccess)

	// TODO: research additional server options
	joinedInterceptors := grpc_middleware.WithUnaryServerChain(interceptors...)

	return &Server{
		root: grpc.NewServer(joinedInterceptors),
		cfg:  cfg,
	}
}

func makeDefaultInterceptors(logger *zap.Logger) []grpc.UnaryServerInterceptor {
	checkIsMethodCallLoggable := func(context.Context, string, interface{}) bool {
		return true // enable every method call logging by default
	}

	return []grpc.UnaryServerInterceptor{
		// Recover after panics
		// grpc_recovery.UnaryServerInterceptor(), // TODO: enable
		// Enable rpc method calls logging
		grpc_zap.PayloadUnaryServerInterceptor(logger, checkIsMethodCallLoggable),
	}
}

func (s *Server) Serving(ctx context.Context) error {
	listener, err := net.Listen(s.cfg.Network.String(), s.cfg.Host)
	if err != nil {
		return errors.WithMessage(err, "running address listener on the network")
	}

	// Init new stop signal channel, 'cause current channels is closed
	s.stopSignal = make(chan struct{})
	go s.contextDoneServing(ctx)

	err = s.root.Serve(listener)
	if err != nil {
		// Have to stop in case serve process being failed by itself,
		// in order to cancel depended running goroutines
		s.Stop()

		return errors.WithMessage(err, "serving listener connections")
	}
	return nil
}

func (s *Server) contextDoneServing(ctx context.Context) {
	select {
	case <-ctx.Done():
		s.Stop()

	case <-s.stopSignal:
		return
	}
}

func (s *Server) Stop() {
	select {
	case <-s.stopSignal:
		return // already stopped

	default:
		close(s.stopSignal)
	}

	s.root.GracefulStop()
}

func (s *Server) Root() *grpc.Server {
	return s.root
}
