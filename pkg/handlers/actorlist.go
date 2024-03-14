package handlers

import (
	"encoding/json"
	"fmt"
	movieapi "movieAPI"
	"net/http"
)

func (h *Handler) createActorlist(w http.ResponseWriter, r *http.Request) {
	var input movieapi.ActorList
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		servErr(w, err)
		return
	}
	id, err := h.services.ActorList.CreateActor(1, input)
	if err != nil {
		servErr(w, err)
		return
	}
	res, err := JSONStruct(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		servErr(w, err)
	}
	fmt.Fprintf(w, "%v", res)

}

func (h *Handler) getAllActorList(w http.ResponseWriter, r *http.Request) {
	lists, err := h.services.ListActors()
	if err != nil {
		servErr(w, err)
	}

	res, err := JSONStruct(lists)
	if err != nil {
		servErr(w, err)
	}
	fmt.Fprintf(w, "%v", res)
}
