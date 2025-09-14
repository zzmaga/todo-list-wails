package service

import (
	"context"

	"todo-list-wails/backend/models"
	"todo-list-wails/backend/usecase"
)

// TaskService содержит сервисный слой для работы с задачами
type TaskService struct {
	u *usecase.TaskUsecase
}

// Constructor
func NewTaskService(u *usecase.TaskUsecase) *TaskService {
	return &TaskService{u: u}
}

// List возвращает список всех задач
func (s *TaskService) List(ctx context.Context) ([]models.Task, error) {
	return s.u.List(ctx)
}

// Toggle переключает статус выполнения задачи
func (s *TaskService) Toggle(ctx context.Context, id string) error {
	return s.u.Toggle(ctx, id)
}

// Update обновляет существующую задачу
func (s *TaskService) Update(ctx context.Context, t models.Task) error {
	return s.u.Update(ctx, t)
}
