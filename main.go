package main

import (
	"embed"
	"log"

	"todo-list-wails/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	app := backend.NewApp(backend.Config{AppName: "Wails To-Do"})

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
