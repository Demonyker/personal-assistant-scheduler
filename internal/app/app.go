package app

import (
	"fmt"
	"github.com/Demonyker/personal-assistant-scheduler/internal/repo/persistence"
	"github.com/Demonyker/personal-assistant-scheduler/pkg/postgres"
	"os"
	"os/signal"
	"syscall"

	"github.com/Demonyker/personal-assistant-scheduler/config"
	"github.com/Demonyker/personal-assistant-scheduler/internal/controller/grpc"
	"github.com/Demonyker/personal-assistant-scheduler/internal/usecase/scheduler"
	"github.com/Demonyker/personal-assistant-scheduler/pkg/grpcserver"
	"github.com/Demonyker/personal-assistant-scheduler/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.DB.Url)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use-Case
	schedulerUseCase := scheduler.New(
		persistent.New(pg),
	)
	// gRPC Server
	grpcServer := grpcserver.New(grpcserver.Port(cfg.GRPC.Port))
	grpc.NewRouter(grpcServer.App, schedulerUseCase, l)
	grpcServer.Start()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: %s", s.String())
	case err = <-grpcServer.Notify():
		l.Error(fmt.Errorf("app - Run - grpcServer.Notify: %w", err))
	}

	err = grpcServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - grpcServer.Shutdown: %w", err))
	}
}
