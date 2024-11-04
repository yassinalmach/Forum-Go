package api

import (
	"encoding/json"
	"forum/controller"
	"forum/models"
	"net/http"
)

type resp struct {
	Err string `json:"error,omitempty"`
}

// this func responsible for writing responses to me for debbuging
func responseJSON(w http.ResponseWriter, code int, resp resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

// this func handles user regestration requests
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

	responseJSON(w, http.StatusCreated, resp{Err: "user registered"})
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// this func handles login requests
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		responseJSON(w, 400, resp{Err: err.Error()})
		return
	}

	id, err := controller.GetUser(user.Username, user.Password)
	if err != nil {
		responseJSON(w, 400, resp{Err: err.Error()})
		return
	}

	if err := controller.AddSession(id, w); err != nil {
		responseJSON(w, 500, resp{Err: err.Error()})
		return
	}

	responseJSON(w, http.StatusOK, resp{Err: "user logged in"})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
