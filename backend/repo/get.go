package repo

import (
	"context"
	"os"
	"todo-list-wails/backend/models"
)

// Get возвращает задачу по ID из файла
func (r *FileRepo) Get(ctx context.Context, id string) (*models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	tasks, err := r.readAll()
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		if task.ID == id {
			tt := task
			return &tt, nil
		}
	}
	return nil, os.ErrNotExist
}
