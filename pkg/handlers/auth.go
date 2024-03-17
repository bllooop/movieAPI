package handlers

import (
	"encoding/json"
	"fmt"
	movieapi "movieapi"
	"net/http"
)

func JSONStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input movieapi.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		clientErr(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		clientErr(w, http.StatusBadRequest, err.Error())
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

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		clientErr(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.CreateToken(input.Username, input.Password)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})
	r.Header.Set("Authorization", "Bearer "+token)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
