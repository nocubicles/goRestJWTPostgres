package middleware

import (
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("superStrongSecret")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

// AuthMiddleware to authenticate users
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := verifyToken(tokenString)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		userId := claims.(jwt.MapClaims)["userId"].(string)

		r.Header.Set("userId", userId)

		next.ServeHTTP(w, r)
	})
}
