package main

import (
	"fmt"

	core_server "github.com/Fitray/go_final_project/internal/core/server"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	serverConfig, err := core_server.NewConfig()
	if err != nil {
		panic(err)
	}
	httpServer := core_server.NewHTTPServer(serverConfig)
	httpServer.RegisterRoutes()

	if err := httpServer.Run(); err != nil {
		fmt.Println(err)
	}
}
