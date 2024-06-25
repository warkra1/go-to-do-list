package main

import (
	"to-do-list/app"
	"to-do-list/app/repository"
	"to-do-list/app/service"
)

func buildCommandProcessor() *app.CommandProcessor {
	userRepository := repository.NewUserRepository("data/user")
	toDoListRepository := repository.NewToDoListRepository("data/list")
	authService := service.NewAuthService(userRepository)
	toDoListService := service.NewToDoListService(toDoListRepository)
	return app.NewCommandProcessor(authService, toDoListService)
}

func main() {
	commandProcessor := buildCommandProcessor()
	commandProcessor.Run()
}
