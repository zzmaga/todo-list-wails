package service

import (
	"context"
	"time"

	"todo-list-wails/backend/models"
	"todo-list-wails/backend/usecase"
)

type TaskService struct {
	u *usecase.TaskUsecase
}

func NewTaskService(u *usecase.TaskUsecase) *TaskService {
	return &TaskService{u: u}
}

func (s *TaskService) List(ctx context.Context) ([]models.Task, error) {
	return s.u.List(ctx)
}
func (s *TaskService) Add(ctx context.Context, title string, due *time.Time, p models.Priority) (models.Task, error) {
	return s.u.Add(ctx, title, due, p)
}
func (s *TaskService) Toggle(ctx context.Context, id string) error {
	return s.u.Toggle(ctx, id)
}
func (s *TaskService) Update(ctx context.Context, t models.Task) error {
	return s.u.Update(ctx, t)
}
func (s *TaskService) Delete(ctx context.Context, id string) error {
	return s.u.Delete(ctx, id)
}
