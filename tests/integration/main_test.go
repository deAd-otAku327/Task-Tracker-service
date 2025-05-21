package integration

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"task-tracker-service/internal/app"
	"task-tracker-service/internal/config"
	"task-tracker-service/tests/testenv"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	configDir := filepath.Join("..", "..", "configs")
	configPath := filepath.Join(configDir, fmt.Sprintf("%s.yaml", testenv.ENV))

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

	testenv.TestManager.TestApp, err = app.New(cfg)
	if err != nil {
		log.Panicf("Failed to init test app: %v", err)
	}

	m.Run()
}

func TestForbidden(t *testing.T) {
	endpoints := []string{
		"GET /users", "GET /tasks", "GET /tasks/1", "POST /tasks/create", "POST /tasks/update",
		"POST /comment", "GET /dashboards", "GET /dashboards/1", "POST /dashboards/create",
		"POST /dashboards/update", "POST /dashboards/delete", "POST /dashboards/addBoardAdmin",
		"POST /dashboards/deleteBoardAdmin",
	}

	w := httptest.NewRecorder()
	for _, endpoint := range endpoints {
		parsed := strings.Fields(endpoint)

		request := httptest.NewRequest(parsed[0], parsed[1], nil)

		testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, request)

		assert.Equal(t, http.StatusForbidden, w.Result().StatusCode)
	}
}
