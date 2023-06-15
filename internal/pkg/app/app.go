package app

import "main/internal/app/start"

type App struct {
	main *start.Start
}

func (a *App) Run() *App {
	app := &App{}
	app.main = start.New()

	app.main.Start()

	return app
}
