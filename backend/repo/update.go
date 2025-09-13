package repo

import (
	"context"
	"os"
	"todo-list-wails/backend/models"
)

func (r *FileRepo) Update(ctx context.Context, t models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	tasks, err := r.readAll()
	if err != nil {
		return err
	}
	for i, it := range tasks {
		if it.ID == t.ID {
			tasks[i] = t
			return r.writeAll(tasks)
		}
	}
	return os.ErrNotExist
}
