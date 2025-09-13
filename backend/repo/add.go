package repo

import (
	"context"
	"todo-list-wails/backend/models"
)

func (r *FileRepo) Add(ctx context.Context, t models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	tasks, err := r.readAll()
	if err != nil {
		return err
	}
	tasks = append(tasks, t)
	return r.writeAll(tasks)
}
