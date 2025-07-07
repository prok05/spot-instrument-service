package app

import (
	"context"
	"github.com/prok05/spot-instrument-service/config"
	"github.com/prok05/spot-instrument-service/internal/controller/grpc"
	"github.com/prok05/spot-instrument-service/internal/controller/grpc/interceptor"
	"github.com/prok05/spot-instrument-service/internal/controller/grpc/server"
	"github.com/prok05/spot-instrument-service/internal/repo/in_memory"
	"github.com/prok05/spot-instrument-service/internal/usecase/market"
	"github.com/prok05/spot-instrument-service/pkg/logger"
	"github.com/prok05/spot-instrument-service/pkg/metric"
	"github.com/prok05/spot-instrument-service/pkg/tracer"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// logger
	l, err := logger.New(cfg.Log.Level)
	if err != nil {
		log.Fatalf("logger error: %s", err)
	}
	l.Info("logger initialized")

	// tracer
	shutdown, err := tracer.New(ctx, cfg)
	if err != nil {
		l.Fatal("app - Run - tracer.New", "failed to create tracer", err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			l.Fatal("app - Run - shutdown", "failed to shutdown tracer", err)
		}
	}()
	l.Info("tracer connected")

	// metric
	go func() {
		if err := metric.Start(cfg); err != nil {
			l.Fatal("app - Run - metric.Start", "prometheus start failed", err)
		}
	}()

	// usecase
	marketUsecase := market.New(in_memory.New())

	// grpc server
	interceptors := interceptor.New(l)

	grpcServer := server.New(
		server.Port(cfg.GRPC.Port),
		server.WithUnaryInterceptor(interceptors.XRequestID()),
		server.WithUnaryInterceptor(interceptors.Log()),
		server.WithUnaryInterceptor(interceptors.Panic()),
	)

	grpc.NewRouter(grpcServer.App, marketUsecase, l)

	// start server
	grpcServer.Start()

	l.Info("service started", "name", cfg.App.Name, "port", grpcServer.Address)

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
