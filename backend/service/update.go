package service

import (
	"context"
	"todo-list-wails/backend/models"
)

// Update обновляет существующую задачу
func (s *TaskService) Update(ctx context.Context, t models.Task) error {
	return s.u.Update(ctx, t)
}
