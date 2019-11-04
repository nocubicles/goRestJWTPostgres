package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "root"
	DB_PASSWORD = "teretere1"
	DB_NAME     = "postgresdb"
)

var Database *sql.DB

func CreateDatabase() (*sql.DB, error) {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	var err error
	Database, err = sql.Open("postgres", dbinfo)

	if err != nil {
		return nil, err
	}

	if err := migrateDatabase(Database); err != nil {
		return Database, err
	}

	return Database, nil

}

func migrateDatabase(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		return err
	}

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/db/migrations", dir),
		"postgres",
		driver,
	)

	if err != nil {
		return err
	}

	migration.Log = &MigrationLogger{}

	migration.Log.Printf("Applying Database migrations")

	err = migration.Up()

	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	version, _, err := migration.Version()

	if err != nil {
		return err
	}

	migration.Log.Printf("Active database version: %d", version)

	return nil
}
