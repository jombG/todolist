package health

import (
	"context"
	"todolist/internal/model"
)

type (
	Health interface {
		GetHealth(ctx context.Context) (*model.Health, error)
	}

	useCaseImpl struct {
	}
)

func New() Health {
	return &useCaseImpl{}
}

func (u *useCaseImpl) GetHealth(_ context.Context) (*model.Health, error) {
	return &model.Health{Ok: true}, nil
}
