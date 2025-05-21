package testenv

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/ory/dockertest/v3"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var dbEnv = map[string]string{
	"PGUSER":            "postgres",
	"POSTGRES_PASSWORD": "testing123",
	"POSTGRES_DB":       "task-tracker-db",
}

func InitPostgresTestDB(pool *dockertest.Pool) (*dockertest.Resource, string, error) {
	log.Printf("Postgres starting...")
	res, err := pool.RunWithOptions(&dockertest.RunOptions{
		Name: "task-tracker-db-testing",
		Env: []string{
			fmt.Sprintf("PGUSER=%s", dbEnv["PGUSER"]),
			fmt.Sprintf("POSTGRES_PASSWORD=%s", dbEnv["POSTGRES_PASSWORD"]),
			fmt.Sprintf("POSTGRES_DB=%s", dbEnv["POSTGRES_DB"]),
		},
		Repository: "postgres",
		Tag:        "latest",
	})
	if err != nil {
		return nil, "", err
	}
	log.Printf("Postgres started")

	dbConn := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s?sslmode=disable",
		dbEnv["PGUSER"], dbEnv["POSTGRES_PASSWORD"], res.GetPort("5432/tcp"), dbEnv["POSTGRES_DB"],
	)

	var db *sql.DB
	if err := pool.Retry(func() error {
		log.Println("Checking postgres connection...")
		db, err = sql.Open("postgres", dbConn)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		return res, "", err
	}

	log.Println("Postgres connection established")

	migrationsPath := "file://../../internal/storage/migrations"

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return res, "", err
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		return res, "", err
	}

	if err := m.Up(); err != nil {
		return res, "", err
	}

	return res, dbConn, nil
}

func Teardown(pool *dockertest.Pool, resource *dockertest.Resource) error {
	if err := pool.Purge(resource); err != nil {
		return err
	}

	return nil
}
