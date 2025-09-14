package usecase

import (
	"context"
	"time"

	"todo-list-wails/backend/models"
	"todo-list-wails/backend/repo"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	repo repo.TaskRepo
}

func NewTaskUsecase(r repo.TaskRepo) *TaskUsecase {
	return &TaskUsecase{repo: r}
}

// List - список всех задач
func (u *TaskUsecase) List(ctx context.Context) ([]models.Task, error) {
	return u.repo.List(ctx)
}

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

// Toggle переключает статус выполнения задачи
func (u *TaskUsecase) Toggle(ctx context.Context, id string) error {
	task, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	task.Done = !task.Done
	return u.repo.Update(ctx, *task)
}

func (u *TaskUsecase) Update(ctx context.Context, task models.Task) error {
	return u.repo.Update(ctx, task)
}

func (u *TaskUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
