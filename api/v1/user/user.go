package user

import (
	"database/sql"
	"encoding/json"
	"goRestJWTPostgres/models"
	"goRestJWTPostgres/repository"
	"goRestJWTPostgres/repository/user"
	"goRestJWTPostgres/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type User struct {
	repo repository.UserRepo
}

const successMessage = "Success"

func NewUserHandler(db *sql.DB) *User {
	return &User{
		repo: user.NewSQLUserRepo(db),
	}
}

func (u *User) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := u.repo.Fetch(r.Context(), 5)

	utils.RespondwithJSON(w, http.StatusOK, payload)
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	_, err := u.repo.Create(r.Context(), &user)

	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, map[string]interface{}{"Server Error": err})
	} else {
		utils.RespondwithJSON(w, http.StatusCreated, map[string]interface{}{"message": successMessage})

	}

}

func (u *User) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.User{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := u.repo.Update(r.Context(), &data)

	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, map[string]interface{}{"Server Error": err})
	} else {
		utils.RespondwithJSON(w, http.StatusOK, payload)

	}

}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := u.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, map[string]interface{}{"Server Error": err})
	} else {
		utils.RespondwithJSON(w, http.StatusOK, payload)

	}

}

func (u *User) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := u.repo.Delete(r.Context(), int64(id))

	if err != nil {
		utils.RespondwithJSON(w, http.StatusInternalServerError, map[string]interface{}{"Server Error": err})
	} else {
		utils.RespondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": successMessage})

	}

}
