package app

import (
	"github.com/prok05/spot-instrument-service/config"
	"github.com/prok05/spot-instrument-service/internal/controller/grpc"
	"github.com/prok05/spot-instrument-service/internal/controller/grpc/interceptor"
	"github.com/prok05/spot-instrument-service/internal/controller/grpc/server"
	"github.com/prok05/spot-instrument-service/internal/repo/in_memory"
	"github.com/prok05/spot-instrument-service/internal/usecase/market"
	"github.com/prok05/spot-instrument-service/pkg/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	// logger
	l, err := logger.New(cfg.Log.Level)
	if err != nil {
		log.Fatalf("logger error: %s", err)
	}

	// usecase
	marketUsecase := market.New(in_memory.New())

	// grpc server
	interceptors := interceptor.New(l)

	grpcServer := server.New(
		server.Port(cfg.GRPC.Port),
		server.WithUnaryInterceptor(interceptors.XRequestID()),
		server.WithUnaryInterceptor(interceptors.Log()),
		server.WithUnaryInterceptor(interceptors.Panic()),
		// TODO: prometheus interceptor
	)

	grpc.NewRouter(grpcServer.App, marketUsecase, l)

	// start server
	grpcServer.Start()

	l.Info("app started", "name", cfg.App.Name, "port", cfg.GRPC.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run",
			"signal", s.String(),
		)
	case err = <-grpcServer.Notify():
		l.Error("app - Run - httpServer.Notify",
			"err", err,
		)
	}

	if err := grpcServer.Shutdown(); err != nil {
		l.Error("app - Run",
			"httpServer.Shutdown", err,
		)
	}
}
