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

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body movieapi.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} message
// @Failure 500 {string} message
// @Failure default {string} message
// @Router /api/auth/sign-up [post]
func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		clientErr(w, http.StatusMethodNotAllowed, "only post method allowed")
		return
	}
	var input movieapi.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Password == "" || input.UserName == "" || input.Role == "" {
		clientErr(w, http.StatusBadRequest, "invalid input body")
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		servErr(w, err, err.Error())
		return
	}
	res, err := JSONStruct(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		servErr(w, err, err.Error())
		return
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
