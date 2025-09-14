package repo

import (
	"context"
	"todo-list-wails/backend/models"
)

// Delete удаляет задачу из файла
func (r *FileRepo) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	tasks, err := r.readAll()
	if err != nil {
		return err
	}
	out := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		if task.ID != id {
			out = append(out, task)
		}
	}
	return r.writeAll(out)
}
