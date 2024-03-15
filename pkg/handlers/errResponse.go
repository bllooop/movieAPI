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

func servErr(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func clientErr(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
func notFound(w http.ResponseWriter) {
	clientErr(w, http.StatusNotFound)
}
