package handlers

import (
	"encoding/json"
	"fmt"
	movieapi "movieapi"
	"net/http"
	"strconv"
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
func (h *Handler) updateActorList(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		clientErr(w, http.StatusBadRequest)
		return
	}
	var input movieapi.UpdateActorListInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		servErr(w, err)
		return
	}
	if err := h.services.ActorList.Update(1, id, input); err != nil {
		servErr(w, err)
	}
	res, err := JSONStruct(statusResponse{
		Status: "ok",
	})
	if err != nil {
		servErr(w, err)
	}
	fmt.Fprintf(w, "%v", res)
}
func (h *Handler) deleteActorList(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		clientErr(w, http.StatusBadRequest)
		return
	}
	err = h.services.ActorList.Delete(1, id)
	if err != nil {
		servErr(w, err)
		return
	}
	res, err := JSONStruct(statusResponse{
		Status: "ok",
	})
	if err != nil {
		servErr(w, err)
	}
	fmt.Fprintf(w, "%v", res)
}
