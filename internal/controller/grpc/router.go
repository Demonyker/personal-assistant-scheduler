package grpc

import (
	v1 "github.com/Demonyker/personal-assistant-scheduler/internal/controller/grpc/v1"
	"github.com/Demonyker/personal-assistant-scheduler/internal/usecase"
	"github.com/Demonyker/personal-assistant-scheduler/pkg/logger"
	pbgrpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewRouter -.
func NewRouter(app *pbgrpc.Server, suc usecase.SchedulerUseCase, l logger.Interface) {
	{
		v1.NewSchedulerRoutes(app, suc, l)
	}

	reflection.Register(app)
}
