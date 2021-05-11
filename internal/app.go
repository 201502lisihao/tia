package internal

import (
	"my-project/internal/app"
)

type App struct {
	Enter *app.Enter
}

func NewApp(
	Enter *app.Enter,
) *App {
	return &App{
		Enter: Enter,
	}
}
