package auth

import (
	"database/sql"
	"fmt"
	"goRestJWTPostgres/db"
	"goRestJWTPostgres/repository"
	"log"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type pgsqlAuthRepo struct {
	Conn *sql.DB
}

func NewSQLAuthRepo(Conn *sql.DB) repository.AuthRepo {
	return &pgsqlAuthRepo{
		Conn: db.Database,
	}
}

func createToken(email string, id int) (string, error) {

	signingKey := []byte("superStrongSecret")

	userId := strconv.Itoa(id)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
	})

	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

func comparePasswords(hashedPWD string, plainPWD []byte) bool {
	byteHash := []byte(hashedPWD)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPWD)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func (m *pgsqlAuthRepo) Create(email string, password string) (string, error) {

	var err error
	var token string

	getUserPasswordQuery := fmt.Sprintf(`SELECT passwordhash, id FROM users WHERE email='%s'`, email)

	var passwordhash string
	var id int
	err = m.Conn.QueryRow(getUserPasswordQuery).Scan(&passwordhash, &id)

	if comparePasswords(passwordhash, []byte(password)) {
		token, err = createToken(email, id)

		setTokenQuery := fmt.Sprintf(`UPDATE users SET bearer = '%s' WHERE email='%s' RETURNING id`,
			token, email)

		err = m.Conn.QueryRow(setTokenQuery).Scan(&id)

	}

	if err != nil {
		return "Something went wrong with sign-in", err
	}

	return token, nil
}
