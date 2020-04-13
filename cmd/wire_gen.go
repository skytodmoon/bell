// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cmd

import (
	"bell/app"
	"bell/app/web"
	"bell/config"
	"bell/controller"
	"bell/library/database"
	log "bell/library/logger"
	"bell/middleware"
	"bell/repository"
	api "bell/router"
	"bell/service"
)

// Injectors from wire.go:

func BuildApp(path2 string) (*app.App, error) {
	conf := config.NewConfig(path2)
	logger := log.New()

	db, err := database.NewDB(conf, logger)
	if err != nil {
		return nil, err
	}

	recoverMid := middleware.NewRecover(logger)

	userRepository := repository.NewUserRepository(logger, db)
	userService := service.NewUserService(userRepository, logger)
	userController := controller.NewUserController(logger, userService)

	engine := app.NewGinEngine()

	router := api.NewRouter(recoverMid, userController)
	server := web.NewServer(engine, router, logger, conf)
	app := app.NewApp(conf, server)

	return app, nil
}
