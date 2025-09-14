package repo

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"todo-list-wails/backend/models"
)

// FileRepo реализует TaskRepo для файлового хранения
type FileRepo struct {
	filename string
}

func NewFileRepo(appName string) (*FileRepo, error) {
	dir, err := userDataDir(appName)
	if err != nil {
		return nil, err
	}

	filename := filepath.Join(dir, "tasks.json")
	return &FileRepo{filename: filename}, nil
}

// userDataDir возвращает путь к директории данных пользователя
func userDataDir(appName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(home, "Library", "Application Support", appName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return dir, nil
}

// readAll читает все задачи из файла
func (r *FileRepo) readAll() ([]models.Task, error) {
	data, err := os.ReadFile(r.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}

	var tasks []models.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// writeAll записывает все задачи в файл
func (r *FileRepo) writeAll(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filename, data, 0644)
}

// List возвращает список всех задач
func (r *FileRepo) List(ctx context.Context) ([]models.Task, error) {
	return r.readAll()
}
