package api

import (
	"encoding/json"
	"forum/controller"
	"forum/models"
	"net/http"
)

type resp struct {
	Err  string `json:"error,omitempty"`
	Data any    `json:"data,omitempty"`
}

func responseJSON(w http.ResponseWriter, code int, resp resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responseJSON(w, 400, resp{Err: err.Error()})
		return
	}

	if err := controller.AddUser(user.Email, user.Username, user.Password); err != nil {
		responseJSON(w, 400, resp{Err: err.Error()})
		return
	}
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

}
