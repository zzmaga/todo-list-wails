package main

import (
	"embed"
	"log"

	"todo-list-wails/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

var assets embed.FS

func main() {
	// Конфигурация приложения
	config := backend.Config{
		AppName: "Wails To-Do",
		Database: backend.DatabaseConfig{
			Type: "postgres",
			Postgres: backend.PostgresConfig{
				Host:     "localhost",
				Port:     5432,
				User:     "magjantastanov",
				Password: "", // Без пароля для локальной разработки
				DBName:   "todo_app",
				SSLMode:  "disable",
			},
		},
	}

	app := backend.NewApp(config)

	if err := wails.Run(&options.App{
		Title:     "Wails To-Do",
		Width:     1024,
		Height:    700,
		Assets:    assets,
		OnStartup: app.Startup,
		Bind:      []interface{}{app},
	}); err != nil {
		log.Fatal(err)
	}
}
