package usecase

import (
	"context"
	"time"

	"todo-list-wails/backend/models"
	"todo-list-wails/backend/repo"

	"github.com/google/uuid"
)

// TaskUsecase содержит бизнес-логику для работы с задачами
// Реализует слой usecase в архитектуре Clean Architecture
type TaskUsecase struct {
	repo repo.TaskRepo // репозиторий для работы с данными
}

// NewTaskUsecase создает новый экземпляр TaskUsecase
func NewTaskUsecase(r repo.TaskRepo) *TaskUsecase {
	return &TaskUsecase{repo: r}
}

// List возвращает список всех задач
func (u *TaskUsecase) List(ctx context.Context) ([]models.Task, error) {
	return u.repo.List(ctx)
}

// Add создает новую задачу с указанными параметрами
// Генерирует уникальный ID и устанавливает время создания
func (u *TaskUsecase) Add(ctx context.Context, title string, due *time.Time, p models.Priority) (models.Task, error) {
	t := models.Task{
		ID:        uuid.NewString(), // генерируем уникальный ID
		Title:     title,
		Done:      false,      // новая задача всегда не выполнена
		CreatedAt: time.Now(), // устанавливаем текущее время
		DueAt:     due,
		Priority:  p,
	}
	return t, u.repo.Add(ctx, t)
}

// Toggle переключает статус выполнения задачи
// Получает задачу по ID, инвертирует статус и сохраняет изменения
func (u *TaskUsecase) Toggle(ctx context.Context, id string) error {
	t, err := u.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	t.Done = !t.Done // инвертируем статус выполнения
	return u.repo.Update(ctx, *t)
}

// Update обновляет существующую задачу
func (u *TaskUsecase) Update(ctx context.Context, t models.Task) error {
	return u.repo.Update(ctx, t)
}

// Delete удаляет задачу по ID
func (u *TaskUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.Delete(ctx, id)
}
