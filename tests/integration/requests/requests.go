package requests

import (
	"bytes"
	"encoding/json"
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
