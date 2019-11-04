package secret

import (
	"fmt"
	"goRestJWTPostgres/utils"
	"net/http"
)

type Secret struct{}

func NewSecretHandler() *Secret {
	return &Secret{}
}

func (secret *Secret) Fetch(w http.ResponseWriter, r *http.Request) {

	userId := r.Header.Get("userId")

	welcomeMessage := fmt.Sprintf("Welcome friend number %s", userId)
	utils.RespondwithJSON(w, http.StatusOK, welcomeMessage)

}
