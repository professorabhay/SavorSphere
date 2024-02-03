package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/professorabhay/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validateSignUpDetails(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// userService.CreateUser(user)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User signed up successfully"}`))
}

func validateSignUpDetails(user *models.User) error {
	return nil
}
