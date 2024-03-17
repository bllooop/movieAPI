package main

import (
	"context"
	"flag"
	"log"
	movieapi "movieapi"
	handler "movieapi/pkg/handlers"
	"movieapi/pkg/repository"
	"movieapi/pkg/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	port := "8000"
	addr := flag.String("addr", port, "web-server address")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	dbpool, err := repository.NewPostgresDB()
	if err != nil {
		errorLog.Fatal(err)
	}
	repos := repository.NewRepository(dbpool)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	infoLog.Printf("app is starting on %s port", *addr)
	srv := new(movieapi.Server)
	go func() {
		if err := srv.RunServer(port, handlers.InitRoutes()); err != nil {
			errorLog.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	infoLog.Printf("app is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error occured during server shutdown: %s", err.Error())
	}
	if err := dbpool.Close(); err != nil {
		log.Fatalf("error occured during closing db conn: %s", err.Error())
	}
}
