package service

import (
	"context"
	"time"
	"todo-list-wails/backend/models"
)

type TaskServiceInterface interface {
	List(ctx context.Context) ([]models.Task, error)
	Toggle(ctx context.Context, id string) error
	Update(ctx context.Context, task models.Task) error
	Delete(ctx context.Context, id string) error
	Add(ctx context.Context, title string, due *time.Time, p models.Priority) (models.Task, error)
}
