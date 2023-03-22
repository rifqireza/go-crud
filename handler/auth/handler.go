package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jeypc/lat1/entity"
	"github.com/jeypc/lat1/models"
)

var userModel = models.NewUserModel()

func Login(w http.ResponseWriter, r *http.Request) {
	var creds entity.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.Response{Error: "Invalid request body"})
		return
	}

	ok, err := userModel.AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Response{Error: "Failed to authenticate user"})
		return
	}
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(entity.Response{Error: "Invalid credentials"})
		return
	}

	token, err := userModel.GenerateToken(creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.Response{Error: "Failed to generate token"})
		return
	}

	json.NewEncoder(w).Encode(entity.Response{Token: token})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}
