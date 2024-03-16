package handlers

import (
	"encoding/json"
	"fmt"
	movieapi "movieAPI"
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
		servErr(w, err)
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
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

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		servErr(w, err)
		return
	}
	user, err := h.services.Authorization.SignUser(input.Username, input.Password)
	if err != nil {
		clientErr(w, http.StatusBadRequest)
	}

	token, err := createAndSetAuthCookie(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})
	r.Header.Set("Authorization", "Bearer "+token)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func createAndSetAuthCookie(user movieapi.User) (string, error) {
	token, err := CreateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
