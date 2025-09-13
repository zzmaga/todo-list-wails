package service

import "context"

func (s *TaskService) Delete(ctx context.Context, id string) error {
	return s.u.Delete(ctx, id)
}
