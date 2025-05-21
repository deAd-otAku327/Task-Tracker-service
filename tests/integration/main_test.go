package integration

import (
	"fmt"
	"log"
	"path/filepath"
	"task-tracker-service/internal/app"
	"task-tracker-service/internal/config"
	"task-tracker-service/tests/testenv"
	"testing"

	"github.com/ory/dockertest/v3"
)

var (
	env     = "testing"
	testApp *app.App
)

func TestMain(m *testing.M) {
	configDir := filepath.Join("..", "..", "configs")
	configPath := filepath.Join(configDir, fmt.Sprintf("%s.yaml", env))

	cfg, err := config.New(configPath)
	if err != nil {
		log.Panicf("Failed to load test config: %v", err)
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Panicf("Failed to create docker pool: %v", err)
	}

	testDB, dbConn, err := testenv.InitPostgresTestDB(pool)
	defer func(dbRes *dockertest.Resource) {
		if dbRes != nil {
			err = testenv.Teardown(pool, dbRes)
			if err != nil {
				log.Panicf("Failed to purge resource: %v", err)
			}
		} else {
			log.Println("No resource to purge")
		}
	}(testDB)
	if err != nil {
		log.Panicf("Failed to run docker container with postgres: %v", err)
	}

	cfg.DBConn.URL = dbConn

	testApp, err = app.New(cfg)
	if err != nil {
		log.Panicf("Failed to init test app: %v", err)
	}

	m.Run()
}
