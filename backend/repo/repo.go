package repo

import (
	"context"
	"todo-list-wails/backend/models"
)

type TaskRepo interface {
	List(ctx context.Context) ([]models.Task, error)
	Get(ctx context.Context, id string) (*models.Task, error)
	Add(ctx context.Context, t models.Task) error
	Update(ctx context.Context, t models.Task) error
	Delete(ctx context.Context, id string) error
}
