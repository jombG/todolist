package storage

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
	"todolist/config"
	"todolist/internal/model"
	"todolist/internal/storage/ent"
	"todolist/internal/storage/ent/task"

	"go.uber.org/zap"
)

type Storage interface {
	CreateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	GetTasks(ctx context.Context) ([]*model.Task, error)
	FinishTask(ctx context.Context, id uuid.UUID) error
	DeleteTask(ctx context.Context, id uuid.UUID) error

	Close() error
}

type (
	storageImpl struct {
		log    *zap.Logger
		cfg    *config.Config
		Client *ent.Client
	}
)

func New(ctx context.Context, log *zap.Logger, cfg *config.Config, drv *sql.Driver) (Storage, error) {
	options := []ent.Option{ent.Driver(drv)}
	client := ent.NewClient(options...)
	if err := client.Schema.Create(ctx); err != nil {
		return nil, errors.Wrap(err, "pg.New")
	}

	return &storageImpl{
		log:    log,
		cfg:    cfg,
		Client: client,
	}, nil
}

func (s *storageImpl) CreateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	taskEnt, err := s.Client.Task.Create().
		SetTitle(task.Title).
		SetNillableDescription(task.Description).
		SetCompleted(false).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "save task from ent")
	}

	task.ID = taskEnt.ID
	return task, nil
}

func (s *storageImpl) GetTasks(ctx context.Context) ([]*model.Task, error) {
	tasksEnt, err := s.Client.Task.Query().
		Where(task.And(task.CompletedEQ(false), task.DeleteAtIsNil())).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get tasks from ent")
	}

	tasks := make([]*model.Task, 0, len(tasksEnt))
	for _, elem := range tasksEnt {
		tasks = append(tasks, &model.Task{
			ID:          elem.ID,
			Title:       elem.Title,
			Description: &elem.Description,
		})
	}

	return tasks, nil
}

func (s *storageImpl) FinishTask(ctx context.Context, id uuid.UUID) error {
	if _, err := s.Client.Task.Update().SetCompleted(true).Where(task.IDEQ(id)).Save(ctx); err != nil {
		return errors.Wrap(err, "finish task from ent")
	}

	return nil
}

func (s *storageImpl) DeleteTask(ctx context.Context, id uuid.UUID) error {
	if _, err := s.Client.Task.Update().SetDeleteAt(time.Now()).Where(task.IDEQ(id)).Save(ctx); err != nil {
		return errors.Wrap(err, "delete task from ent")
	}

	return nil
}

func (s *storageImpl) Close() error {
	return s.Client.Close()
}
