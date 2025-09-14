package service

import (
	"context"
	"time"
	"todo-list-wails/backend/models"
)

// Add создает новую задачу с указанными параметрами
func (s *TaskService) Add(ctx context.Context, title string, due *time.Time, p models.Priority) (models.Task, error) {
	return s.u.Add(ctx, title, due, p)
}
