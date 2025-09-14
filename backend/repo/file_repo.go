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

// FileRepo реализует TaskRepo интерфейс для файлового хранения
// Сохраняет задачи в JSON файл в пользовательской директории
type FileRepo struct {
	path string     // путь к файлу с данными
	mu   sync.Mutex // мьютекс для безопасного доступа к файлу
}

// NewFileRepo создает новый экземпляр FileRepo
// Создает директорию для данных приложения и файл tasks.json если его нет
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

// userDataDir возвращает путь к пользовательской директории данных
// для разных операционных систем
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

// readAll читает все задачи из JSON файла
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

// writeAll записывает все задачи в JSON файл
func (r *FileRepo) writeAll(tasks []models.Task) error {
	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.path, b, 0o644)
}

// List возвращает список всех задач из файла
func (r *FileRepo) List(ctx context.Context) ([]models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.readAll()
}
