package usecase

import (
	"context"
	"time"
	"todo-list-wails/backend/models"
)

// TaskUsecaseInterface определяет интерфейс для бизнес-логики работы с задачами
type TaskUsecaseInterface interface {
	List(ctx context.Context) ([]models.Task, error)
	Add(ctx context.Context, title string, due *time.Time, p models.Priority) (models.Task, error)
	Toggle(ctx context.Context, id string) error
	Update(ctx context.Context, task models.Task) error
	Delete(ctx context.Context, id string) error
}
