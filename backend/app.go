package backend

import (
	"context"
	"time"

	"todo-list-wails/backend/models"
	"todo-list-wails/backend/repo"
	"todo-list-wails/backend/service"
	"todo-list-wails/backend/usecase"
)

type App struct {
	Cfg  Config
	Task *service.TaskService
}

func NewApp(cfg Config) *App {
	return &App{Cfg: cfg}
}

func (a *App) Startup(ctx context.Context) {
	// файловый репозиторий (сохранение в JSON)
	repo, _ := repo.NewFileRepo(a.Cfg.AppName)
	uc := usecase.NewTaskUsecase(repo)
	a.Task = service.NewTaskService(uc)
}

// Экспортируемые в frontend методы

func (a *App) List() ([]models.Task, error) {
	return a.Task.List(context.Background())
}

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
