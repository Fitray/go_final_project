package main

import (
	"os"

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

	httpServer := core_server.NewHTTPServer(serverConfig)
	httpServer.RegisterRoutes()

	if err := httpServer.Run(); err != nil {
		panic(err)
	}
}
