package service

import "context"

// Delete удаляет задачу по ID
func (s *TaskService) Delete(ctx context.Context, id string) error {
	return s.u.Delete(ctx, id)
}
