package usecase

import (
	"go.uber.org/zap"
	"todolist/config"
	"todolist/internal/storage"
	"todolist/internal/usecase/health"
	"todolist/internal/usecase/task"
)

type (
	UseCases struct {
		health.Health
		task.Task
	}
)

func New(
	cfg *config.Config,
	log *zap.Logger,
	dbStorage storage.Storage,
) *UseCases {
	return &UseCases{
		Health: health.New(),
		Task:   task.New(cfg, log, dbStorage),
	}
}
