package repo

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"todo-list-wails/backend/models"
)

type FileRepo struct {
	path string
	mu   sync.Mutex
}

func NewFileRepo(appName string) (*FileRepo, error) {
	dir := userDataDir(appName)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	p := filepath.Join(dir, "tasks.json")
	if _, err := os.Stat(p); errors.Is(err, os.ErrNotExist) {
		_ = os.WriteFile(p, []byte("[]"), 0o644)
	}
	return &FileRepo{path: p}, nil
}

func userDataDir(appName string) string {
	switch runtime.GOOS {
	case "windows":
		base := os.Getenv("AppData")
		if base == "" {
			base = "."
		}
		return filepath.Join(base, appName)
	case "darwin":
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "Library", "Application Support", appName)
	default:
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".config", appName)
	}
}

func (r *FileRepo) readAll() ([]models.Task, error) {
	b, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	if err := json.Unmarshal(b, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *FileRepo) writeAll(tasks []models.Task) error {
	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.path, b, 0o644)
}

func (r *FileRepo) List(ctx context.Context) ([]models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.readAll()
}
