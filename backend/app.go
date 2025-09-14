package backend

import (
	"context"
	"log"
	"time"

	"todo-list-wails/backend/models"
	"todo-list-wails/backend/repo"
	"todo-list-wails/backend/service"
	"todo-list-wails/backend/usecase"
)

// App представляет основное приложение с конфигурацией и сервисами
type App struct {
	Cfg  Config
	Task *service.TaskService
	repo repo.TaskRepo // репозиторий для работы с данными
}

// NewApp создает новый экземпляр приложения с заданной конфигурацией
func NewApp(cfg Config) *App {
	return &App{Cfg: cfg}
}

// Startup инициализирует приложение при запуске
// Настраивает репозиторий (файловый или PostgreSQL) и сервисы
func (a *App) Startup(ctx context.Context) {
	var err error

	// Выбираем тип репозитория на основе конфигурации
	switch a.Cfg.Database.Type {
	case "postgres":
		// PostgreSQL репозиторий
		connStr := a.Cfg.Database.Postgres.GetConnectionString()
		a.repo, err = repo.NewPostgresRepo(connStr)
		if err != nil {
			log.Printf("Failed to connect to PostgreSQL, falling back to file storage: %v", err)
			// Fallback на файловый репозиторий
			a.repo, _ = repo.NewFileRepo(a.Cfg.AppName)
		}
	default:
		// Файловый репозиторий (по умолчанию)
		a.repo, err = repo.NewFileRepo(a.Cfg.AppName)
		if err != nil {
			log.Printf("Failed to create file repo: %v", err)
			return
		}
	}

	// Создаем usecase и service
	uc := usecase.NewTaskUsecase(a.repo)
	a.Task = service.NewTaskService(uc)
}

// Экспортируемые в frontend методы

// List возвращает список всех задач
func (a *App) List() ([]models.Task, error) {
	return a.Task.List(context.Background())
}

// Add создает новую задачу с указанными параметрами
// title - заголовок задачи
// dueISO - дата и время выполнения в формате ISO (может быть nil)
// priority - приоритет задачи ("low", "medium", "high")
func (a *App) Add(title string, dueISO *string, priority string) (models.Task, error) {
	var due *time.Time
	if dueISO != nil && *dueISO != "" {
		if t, err := time.Parse(time.RFC3339, *dueISO); err == nil {
			due = &t
		}
	}
	p := models.Priority(priority)
	return a.Task.Add(context.Background(), title, due, p)
}

func (a *App) Toggle(id string) error {
	return a.Task.Toggle(context.Background(), id)
}

func (a *App) Update(t models.Task) error {
	return a.Task.Update(context.Background(), t)
}

func (a *App) Delete(id string) error {
	return a.Task.Delete(context.Background(), id)
}
