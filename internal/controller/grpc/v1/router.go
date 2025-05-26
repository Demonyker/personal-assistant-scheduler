package v1

import (
	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/scheduler/v1"
	"github.com/Demonyker/personal-assistant-scheduler/internal/usecase"
	"github.com/Demonyker/personal-assistant-scheduler/pkg/logger"
	"github.com/go-playground/validator/v10"
	pbgrpc "google.golang.org/grpc"
)

// NewSchedulerRoutes -.
func NewSchedulerRoutes(app *pbgrpc.Server, suc usecase.SchedulerUseCase, l logger.Interface) {
	r := &V1{suc: suc, l: l, v: validator.New(validator.WithRequiredStructEnabled())}

	{
		v1.RegisterSchedulerServer(app, r)
	}
}
