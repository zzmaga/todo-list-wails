package repo

import (
	"context"
	"todo-list-wails/backend/models"
)

// TaskRepo определяет интерфейс для работы с задачами
// Позволяет использовать разные реализации хранения (файл, БД)
type TaskRepo interface {
	List(ctx context.Context) ([]models.Task, error)
	Get(ctx context.Context, id string) (*models.Task, error)
	Add(ctx context.Context, task models.Task) error
	Update(ctx context.Context, task models.Task) error
	Delete(ctx context.Context, id string) error
}
