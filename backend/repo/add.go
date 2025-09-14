package repo

import (
	"context"
	"todo-list-wails/backend/models"
)

// Add добавляет новую задачу в файл
func (r *FileRepo) Add(ctx context.Context, task models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	tasks, err := r.readAll()
	if err != nil {
		return err
	}
	tasks = append(tasks, task)
	return r.writeAll(tasks)
}
