package main

import (
	"log"
	movieapi "movieAPI"
	handler "movieAPI/pkg/handlers"
	"movieAPI/pkg/repository"
	"movieAPI/pkg/service"

	_ "github.com/lib/pq"
)

func main() {
	//password := "54321"
	//database := flag.String("databas", "postgres://postgres:54321@localhost:5433/postgres?sslmode=disable", "Подключение к PSQL")
	dbpool, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	repos := repository.NewRepository(dbpool)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	port := "8000"
	srv := new(movieapi.Server)
	if err := srv.RunServer(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured during running server: %s", err.Error())
	}
	log.Printf("Запуск веб-сервиса на %s", port)
}
