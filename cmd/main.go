package main

import (
	"flag"
	"log"
	movieapi "movieAPI"
	handler "movieAPI/pkg/handlers"
	"movieAPI/pkg/repository"
	"movieAPI/pkg/service"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	password := os.Getenv("DB_PASSWORD")
	database := flag.String("databas", "postgres://admin:"+password+"@localhost:5432/movies?sslmode=disable", "Подключение к PSQL")
	dbpool, err := repository.NewPostgresDB(*database)
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
