package auth

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"goRestJWTPostgres/repository"
	"goRestJWTPostgres/repository/auth"
)

type Auth struct {
	repo repository.AuthRepo
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(db *sql.DB) *Auth {
	return &Auth{
		repo: auth.NewSQLAuthRepo(db),
	}
}

// Create function to create and save token
func (auth *Auth) Create(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	json.NewDecoder(r.Body).Decode(&credentials)

	if len(credentials.Email) == 0 || len(credentials.Password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide email and password to obtain the token"))
		return
	}

	if len(credentials.Email) != 0 && len(credentials.Password) != 0 {
		token, err := auth.repo.Create(credentials.Email, credentials.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating JWT token: " + err.Error()))
		} else {
			w.Header().Set("Authorization", "Bearer "+token)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Token: " + token))
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Email and password do not match"))
		return
	}
}
