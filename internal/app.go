package internal

import (
	"my-project/internal/app"
)

type App struct {
	ServerRunner *app.ServerRunner
}

func NewApp(
	ServerRunner *app.ServerRunner,
) *App {
	return &App{
		ServerRunner: ServerRunner,
	}
}
