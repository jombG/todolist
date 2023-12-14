package task

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"todolist/config"
	"todolist/internal/model"
	"todolist/internal/storage"
)

type Task interface {
	CreateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	GetTasks(ctx context.Context) ([]*model.Task, error)
	FinishTask(ctx context.Context, id uuid.UUID) error
	DeleteTask(ctx context.Context, id uuid.UUID) error
}

type useCaseImpl struct {
	cfg     *config.Config
	log     *zap.Logger
	storage storage.Storage
}

func (u useCaseImpl) CreateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	nTask, err := u.storage.CreateTask(ctx, task)
	if err != nil {
		return nil, errors.Wrap(err, "create task from storage")
	}

	return nTask, nil
}

func (u useCaseImpl) GetTasks(ctx context.Context) ([]*model.Task, error) {
	tasks, err := u.storage.GetTasks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get tasks from storage")
	}

	return tasks, nil
}

func (u useCaseImpl) FinishTask(ctx context.Context, id uuid.UUID) error {
	if err := u.storage.FinishTask(ctx, id); err != nil {
		return errors.Wrap(err, "finish task from storage")
	}
	return nil
}

func (u useCaseImpl) DeleteTask(ctx context.Context, id uuid.UUID) error {
	if err := u.storage.FinishTask(ctx, id); err != nil {
		return errors.Wrap(err, "delete task from storage")
	}
	return nil
}

func New(
	cfg *config.Config,
	log *zap.Logger,
	storage storage.Storage,
) Task {
	return &useCaseImpl{
		cfg:     cfg,
		storage: storage,
		log:     log,
	}
}
