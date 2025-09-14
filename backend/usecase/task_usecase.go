package usecase

import (
	"context"

	"todo-list-wails/backend/models"
	"todo-list-wails/backend/repo"
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

// Toggle переключает статус выполнения задачи
func (u *TaskUsecase) Toggle(ctx context.Context, id string) error {
	task, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	task.Done = !task.Done
	return u.repo.Update(ctx, *task)
}
