package main

import (
	"os"

	tasks_handler "github.com/Fitray/go_final_project/internal/api/tasks/handler"
	tasks_repository "github.com/Fitray/go_final_project/internal/api/tasks/repository"
	tasks_service "github.com/Fitray/go_final_project/internal/api/tasks/service"
	core_db "github.com/Fitray/go_final_project/internal/core/db"
	core_server "github.com/Fitray/go_final_project/internal/core/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	path := os.Getenv("TODO_DBFILE")
	db, err := core_db.Init(path)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	serverConfig, err := core_server.NewConfig()
	if err != nil {
		panic(err)
	}

	tasksRepository := tasks_repository.NewTasksRepository(db)
	tasksService := tasks_service.NewTasksService(&tasksRepository)
	tasksHandler := tasks_handler.NewDateHandler(&tasksService)
	routes := tasksHandler.Routes()

	httpServer := core_server.NewHTTPServer(serverConfig)
	httpServer.Init(routes)

	if err := httpServer.Run(); err != nil {
		panic(err)
	}
}
