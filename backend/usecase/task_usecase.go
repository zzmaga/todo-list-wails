package usecase

import (
	"context"
	"log"
	"time"

	"todo-list-wails/backend/models"
	"todo-list-wails/backend/repo"

	"github.com/google/uuid"
)

type TaskUsecase struct {
	repo repo.TaskRepo
}

// Const
func NewTaskUsecase(r repo.TaskRepo) *TaskUsecase {
	return &TaskUsecase{repo: r}
}

func (u *TaskUsecase) List(ctx context.Context) ([]models.Task, error) {
	return u.repo.List(ctx)
}

func (u *TaskUsecase) Add(ctx context.Context, title string, due *time.Time, p models.Priority) (models.Task, error) {
	t := models.Task{
		ID:        uuid.NewString(),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
		DueAt:     due,
		Priority:  p,
	}
	return t, u.repo.Add(ctx, t)
}

func (u *TaskUsecase) Toggle(ctx context.Context, id string) error {
	t, err := u.repo.Get(ctx, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	t.Done = !t.Done
	return u.repo.Update(ctx, *t)
}

func (u *TaskUsecase) Update(ctx context.Context, t models.Task) error {
	return u.repo.Update(ctx, t)
}

func (u *TaskUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
