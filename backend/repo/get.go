package repo

import (
	"context"
	"os"
	"todo-list-wails/backend/models"
)

// Get возвращает задачу по ID
func (r *FileRepo) Get(ctx context.Context, id string) (*models.Task, error) {
	tasks, err := r.readAll()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}

	return nil, os.ErrNotExist
}
