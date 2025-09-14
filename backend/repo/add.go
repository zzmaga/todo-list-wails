package repo

import (
	"context"
	"todo-list-wails/backend/models"
)

// Add добавляет новую задачу
func (r *FileRepo) Add(ctx context.Context, task models.Task) error {
	tasks, err := r.readAll()
	if err != nil {
		return err
	}

	tasks = append(tasks, task)
	return r.writeAll(tasks)
}
