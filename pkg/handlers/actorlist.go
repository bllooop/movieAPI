package handlers

import (
	"encoding/json"
	"fmt"
	movieapi "movieapi"
	"net/http"
	"strconv"
)

// @Summary Create actor list
// @Security ApiKeyAuth
// @Tags actorLists
// @Description create actor list
// @ID create-actor-list
// @Accept  json
// @Produce  json
// @Param input body movieapi.ActorList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/actors/add [post]
func (h *Handler) createActorlist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	var input movieapi.ActorList
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Name == "" || input.Gender == "" || input.Birthdate == "" {
		clientErr(w, http.StatusBadRequest, "invalid input body")
		return
	}
	// retrievedValue := "1" // when testing uncomment
	retrievedValue := r.Context().Value(roleCtx).(string) // when testing comment
	id, err := h.services.ActorList.CreateActor(retrievedValue, input)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	res, err := JSONStruct(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)

}

// @Summary Get list of actors
// @Security ApiKeyAuth
// @Tags actorLists
// @Description get actor list
// @ID get-actor-list
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/actors [get]
func (h *Handler) getAllActorList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/actors" {
		notFound(w)
		return
	}
	lists, err := h.services.ListActors()
	if err != nil {
		servErr(w, err, err.Error())
	}
	res, err := JSONStruct(lists)
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}

// @Summary Update actor in list
// @Security ApiKeyAuth
// @Tags actorLists
// @Description update movie in list by id
// @ID update-actor-list
// @Accept  json
// @Produce  json
// @Param input body movieapi.ActorList true "list info"
// @Param       id    query     int  false  "actor update by id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/actors/update [post]
func (h *Handler) updateActorList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	//retrievedValue := "1" // when testing uncomment
	retrievedValue := r.Context().Value(roleCtx).(string) // when testing comment
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		clientErr(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	var input movieapi.UpdateActorListInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	if err := h.services.ActorList.Update(retrievedValue, id, input); err != nil {
		servErr(w, err, err.Error())
	}
	res, err := JSONStruct(statusResponse{
		Status: "ok",
	})
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}

// @Summary Delete actor from list
// @Security ApiKeyAuth
// @Tags actorLists
// @Description delete actor from list by id
// @ID delete-actor-list
// @Produce  json
// @Param       id    query     int  false  "actor delete by id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/actors/delete [delete]
func (h *Handler) deleteActorList(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.Header().Set("Allow", http.MethodDelete)
		clientErr(w, http.StatusMethodNotAllowed, "only delete method allowed")
		return
	}
	retrievedValue := r.Context().Value(roleCtx).(string)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		clientErr(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	err = h.services.ActorList.Delete(retrievedValue, id)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	res, err := JSONStruct(statusResponse{
		Status: "ok",
	})
	if err != nil {
		servErr(w, err, err.Error())
	}
	fmt.Fprintf(w, "%v", res)
}
