package repo

import (
	"context"
	"os"
)

// Delete удаляет задачу по ID
func (r *FileRepo) Delete(ctx context.Context, id string) error {
	tasks, err := r.readAll()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return r.writeAll(tasks)
		}
	}

	return os.ErrNotExist
}
