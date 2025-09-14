package repo

import (
	"context"
	"os"
	"todo-list-wails/backend/models"
)

// Update обновляет существующую задачу
func (r *FileRepo) Update(ctx context.Context, task models.Task) error {
	tasks, err := r.readAll()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == task.ID {
			tasks[i] = task
			return r.writeAll(tasks)
		}
	}

	return os.ErrNotExist
}
