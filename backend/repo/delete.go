package repo

import (
	"context"
	"todo-list-wails/backend/models"
)

func (r *FileRepo) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	tasks, err := r.readAll()
	if err != nil {
		return err
	}
	out := make([]models.Task, 0, len(tasks))
	for _, t := range tasks {
		if t.ID != id {
			out = append(out, t)
		}
	}
	return r.writeAll(out)
}
