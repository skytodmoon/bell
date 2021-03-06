package app

import (
	"bell/app/web"
	"bell/config"
)

type App struct {
	conf      *config.Config
	webServer *web.Server
}

func NewApp(conf *config.Config, server *web.Server) *App {
	return &App{
		conf:      conf,
		webServer: server,
	}
}

func (app *App) Start() {
	app.webServer.Start()
}
