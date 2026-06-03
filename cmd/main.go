package main

import (
	"os"
	"time"

	auth_handler "github.com/Fitray/go_final_project/internal/api/auth/handler"
	auth_repository "github.com/Fitray/go_final_project/internal/api/auth/repository"
	auth_service "github.com/Fitray/go_final_project/internal/api/auth/service"
	tasks_handler "github.com/Fitray/go_final_project/internal/api/tasks/handler"
	tasks_repository "github.com/Fitray/go_final_project/internal/api/tasks/repository"
	tasks_service "github.com/Fitray/go_final_project/internal/api/tasks/service"
	core_auth "github.com/Fitray/go_final_project/internal/core/auth"
	core_db "github.com/Fitray/go_final_project/internal/core/db"
	core_server "github.com/Fitray/go_final_project/internal/core/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	path := os.Getenv("TODO_DBFILE")
	if len(path) == 0 {
		path = "data/scheduler.db"
	}

	db, err := core_db.Init(path, 10*time.Second)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	serverConfig, err := core_server.NewConfig()
	if err != nil {
		panic(err)
	}

	auth, err := core_auth.NewAuth()
	if err != nil {
		panic(err)
	}

	var routes []core_server.Route

	tasksRepository := tasks_repository.NewTasksRepository(db)
	tasksService := tasks_service.NewTasksService(&tasksRepository)
	tasksHandler := tasks_handler.NewDateHandler(&tasksService)
	routes = tasksHandler.Routes(routes, auth)

	authRepository := auth_repository.NewAuthRepository(auth)
	authService := auth_service.NewAuthService(&authRepository)
	authHandler := auth_handler.NewAuthHandler(&authService)
	routes = authHandler.Routes(routes)

	httpServer := core_server.NewHTTPServer(serverConfig)
	httpServer.Init(routes)

	if err := httpServer.Run(); err != nil {
		panic(err)
	}
}
