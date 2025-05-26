package v1

import (
	v1 "github.com/Demonyker/personal-assistant-contracts/contracts/scheduler/v1"
	"github.com/Demonyker/personal-assistant-scheduler/internal/usecase"
	"github.com/Demonyker/personal-assistant-scheduler/pkg/logger"
	"github.com/go-playground/validator/v10"
)

// V1 -.
type V1 struct {
	v1.SchedulerServer

	suc usecase.SchedulerUseCase
	l   logger.Interface
	v   *validator.Validate
}
