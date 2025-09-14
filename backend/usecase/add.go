package usecase

import (
	"context"
	"time"
	"todo-list-wails/backend/models"

	"github.com/google/uuid"
)

// Создаем новую задачу с указанными параметрами
func (u *TaskUsecase) Add(ctx context.Context, title string, due *time.Time, p models.Priority) (models.Task, error) {
	task := models.Task{
		ID:        uuid.NewString(),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
		DueAt:     due,
		Priority:  p,
	}
	return task, u.repo.Add(ctx, task)
}
