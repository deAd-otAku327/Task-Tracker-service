package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"task-tracker-service/internal/types/dto"
	"testing"

	"github.com/stretchr/testify/require"
)

func PrepareRegisterRequest(t *testing.T, input dto.PostUsersRegisterRequest) *http.Request {
	jsonData, err := json.Marshal(input)
	require.NoError(t, err, "Failed to marshal request body")

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")

	return req
}

func PrepareLoginRequest(t *testing.T, input dto.PostUsersLoginRequest) *http.Request {
	jsonData, err := json.Marshal(input)
	require.NoError(t, err, "Failed to marshal request body")

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")

	return req
}

func PrepareGetUsersRequest(t *testing.T, cookie *http.Cookie) *http.Request {
	req, err := http.NewRequest("GET", "/users", nil)
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}

func PrepareCreateTaskRequest(t *testing.T, input dto.PostTasksCreateRequest, cookie *http.Cookie) *http.Request {
	jsonData, err := json.Marshal(input)
	require.NoError(t, err, "Failed to marshal request body")

	req, err := http.NewRequest("POST", "/tasks/create", bytes.NewBuffer(jsonData))
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}

func PrepareUpdateTaskRequest(t *testing.T, input dto.PostTasksUpdateRequest, cookie *http.Cookie) *http.Request {
	jsonData, err := json.Marshal(input)
	require.NoError(t, err, "Failed to marshal request body")

	req, err := http.NewRequest("POST", "/tasks/update", bytes.NewBuffer(jsonData))
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}

func PrepareGetTasksRequest(t *testing.T, cookie *http.Cookie) *http.Request {
	req, err := http.NewRequest("GET", "/tasks", nil)
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}

func PrepareGetTaskSummaryRequest(t *testing.T, taskID int, cookie *http.Cookie) *http.Request {
	req, err := http.NewRequest("GET", fmt.Sprintf("/tasks/%d", taskID), nil)
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}

func PrepareCreateCommentRequest(t *testing.T, input dto.PostCommentRequest, cookie *http.Cookie) *http.Request {
	jsonData, err := json.Marshal(input)
	require.NoError(t, err, "Failed to marshal request body")

	req, err := http.NewRequest("POST", "/comment", bytes.NewBuffer(jsonData))
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}

func PrepareCreateDashboardRequest(t *testing.T, input dto.PostDashboardsCreateRequest, cookie *http.Cookie) *http.Request {
	jsonData, err := json.Marshal(input)
	require.NoError(t, err, "Failed to marshal request body")

	req, err := http.NewRequest("POST", "/dashboards/create", bytes.NewBuffer(jsonData))
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}

func PrepareUpdateDashboardRequest(t *testing.T, input dto.PostDashboardsUpdateRequest, cookie *http.Cookie) *http.Request {
	jsonData, err := json.Marshal(input)
	require.NoError(t, err, "Failed to marshal request body")

	req, err := http.NewRequest("POST", "/dashboards/update", bytes.NewBuffer(jsonData))
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}

func PrepareGetDashboardSummaryRequest(t *testing.T, boardID int, cookie *http.Cookie) *http.Request {
	req, err := http.NewRequest("GET", fmt.Sprintf("/dashboards/%d", boardID), nil)
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}

func PrepareAddDashboardAdminRequest(t *testing.T, input dto.PostDashboardsAdminRequest, cookie *http.Cookie) *http.Request {
	jsonData, err := json.Marshal(input)
	require.NoError(t, err, "Failed to marshal request body")

	req, err := http.NewRequest("POST", "/dashboards/addBoardAdmin", bytes.NewBuffer(jsonData))
	require.NoError(t, err, "Failed to create request")
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(cookie)

	return req
}
