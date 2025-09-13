package repo

import (
	"context"
	"os"
	"todo-list-wails/backend/models"
)

func (r *FileRepo) Get(ctx context.Context, id string) (*models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	tasks, err := r.readAll()
	if err != nil {
		return nil, err
	}
	for _, t := range tasks {
		if t.ID == id {
			tt := t
			return &tt, nil
		}
	}
	return nil, os.ErrNotExist
}
