package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) movieList(w http.ResponseWriter, r *http.Request) {
	a := "test"
	fmt.Fprintf(w, "%v", a)
}
