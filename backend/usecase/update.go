package usecase

import (
	"context"
	"todo-list-wails/backend/models"
)

func (u *TaskUsecase) Update(ctx context.Context, task models.Task) error {
	return u.repo.Update(ctx, task)
}
