package handlers

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

type statusResponse struct {
	Status string `json: status`
}

func servErr(w http.ResponseWriter, err error, message string) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.Println(trace)
	log.Println(message)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientErr(w http.ResponseWriter, status int, message string) {
	log.Fatal(message)
	http.Error(w, http.StatusText(status), status)
}

func notFound(w http.ResponseWriter) {
	clientErr(w, http.StatusNotFound, "page not found")
}
