package main

import (
	"log"
	"reddit/pkg/handler"
	"reddit/pkg/repository"
	"reddit/pkg/service"
)

func main() {
	serverInstance := new(Server)
	postgresConfig := repository.Config{
		Host:     "",
		Port:     "",
		Username: "",
		Password: "",
		DBName:   "",
		SSLMode:  "",
	}
	database, err := repository.NewPostgresDB(postgresConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	repos := repository.NewRepository(database)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	runServer(serverInstance, handlers)
}

func runServer(serverInstance *Server, handlerLayer *handler.Handler) {
	port := "8080"
	router := handlerLayer.GetRouter()

	if err := serverInstance.Run(port, router); err != nil {
		log.Fatal(err.Error())
	}

	log.Print("server started successfully")
}
