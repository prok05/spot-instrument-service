package server

import (
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	pbgrpc "google.golang.org/grpc"
	"net"
)

const (
	_defaultAddr = ":80"
)

type Server struct {
	App       *pbgrpc.Server
	Address   string
	unaryInts []pbgrpc.UnaryServerInterceptor
	notify    chan error
}

func New(opts ...Option) *Server {
	s := &Server{
		unaryInts: make([]pbgrpc.UnaryServerInterceptor, 0),
		notify:    make(chan error),
		Address:   _defaultAddr,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.App = pbgrpc.NewServer(
		pbgrpc.ChainUnaryInterceptor(s.unaryInts...),
		pbgrpc.StatsHandler(otelgrpc.NewServerHandler()))
	return s
}

func (s *Server) Start() {
	go func() {
		ln, err := net.Listen("tcp", s.Address)
		if err != nil {
			s.notify <- fmt.Errorf("failed to listen: %w", err)
			close(s.notify)

			return
		}

		s.notify <- s.App.Serve(ln)
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	s.App.GracefulStop()

	return nil
}
