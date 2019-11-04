package main

import (
	"goRestJWTPostgres/api/v1/auth"
	"goRestJWTPostgres/api/v1/secret"
	"goRestJWTPostgres/api/v1/user"

	"goRestJWTPostgres/db"
	"goRestJWTPostgres/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := chi.NewRouter()

	var err error

	database, err := db.CreateDatabase()

	defer database.Close()

	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	userHandler := user.NewUserHandler(database)
	authHandler := auth.NewAuthHandler(database)
	secretHandler := secret.NewSecretHandler()

	router.Use(middleware.LoggingMiddleware)

	router.Route("/v1", func(rt chi.Router) {
		rt.Mount("/user", userRouter(userHandler))
		rt.Mount("/", authRouter(authHandler))
		rt.Mount("/secret", secretRouter(secretHandler))
	})

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":3000", nil))
}

func userRouter(userHandler *user.User) http.Handler {

	router := chi.NewRouter()
	router.Get("/", userHandler.Fetch)
	router.Get("/{id:[0-9]+}", userHandler.GetByID)
	router.Post("/", userHandler.Create)
	router.Put("/{id:[0-9]+}", userHandler.Update)
	router.Delete("/{id:[0-9]+}", userHandler.Delete)

	return router
}

func authRouter(authHandler *auth.Auth) http.Handler {
	router := chi.NewRouter()

	router.Post("/login", authHandler.Create)
	return router
}

func secretRouter(secretHandler *secret.Secret) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.AuthMiddleware)

	router.Get("/", secretHandler.Fetch)
	return router
}
