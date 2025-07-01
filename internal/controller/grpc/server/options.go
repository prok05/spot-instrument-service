package server

import (
	pbgrpc "google.golang.org/grpc"
	"net"
)

type Option func(*Server)

func Port(port string) Option {
	return func(s *Server) {
		s.Address = net.JoinHostPort("", port)
	}
}

func WithUnaryInterceptor(i pbgrpc.UnaryServerInterceptor) Option {
	return func(s *Server) {
		s.unaryInts = append(s.unaryInts, i)
	}
}
