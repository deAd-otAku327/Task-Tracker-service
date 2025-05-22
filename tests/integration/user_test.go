package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-tracker-service/internal/middleware"
	"task-tracker-service/internal/service/_shared/serverrors"
	"task-tracker-service/internal/types/dto"
	req "task-tracker-service/tests/integration/requests"
	"task-tracker-service/tests/testenv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var DefaultRegisterBody = dto.PostUsersRegisterRequest{
	Email:    "test@example.com",
	Username: "testuser",
	Password: "securepassword123",
}

var DefaultLoginBody = dto.PostUsersLoginRequest{
	Username: "testuser",
	Password: "securepassword123",
}

func TestRegisterUser_Success(t *testing.T) {
	w := httptest.NewRecorder()
	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))

	assert.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var response dto.UserResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.NotZero(t, response.ID, "Expected user ID to be non-zero")
	assert.Equal(t, "testuser", response.Username, "Unexpected username")
	assert.Equal(t, "test@example.com", response.Email, "Unexpected email")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestRegisterUser_OccupiedEmail(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	userDuplicateEmail := DefaultRegisterBody
	userDuplicateEmail.Username = "testuser2"

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, userDuplicateEmail))
	assert.Equal(t, http.StatusBadRequest, w.Code, "Unexpected status code")

	var response dto.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.Equal(t, serverrors.ErrEmailOccupied.Error(), response.Message, "Unexpected error msg")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestRegisterUser_OccupiedUsername(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	userDuplicateUsername := DefaultRegisterBody
	userDuplicateUsername.Email = "test2@example.com"

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, userDuplicateUsername))
	assert.Equal(t, http.StatusBadRequest, w.Code, "Unexpected status code")

	var response dto.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.Equal(t, serverrors.ErrUsernameOccupied.Error(), response.Message, "Unexpected error msg")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestLogin_Success(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, DefaultLoginBody))
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")
	require.Equal(t, 1, len(w.Result().Cookies()), "Unexpected no cookie")

	cookie := w.Result().Cookies()[0]

	assert.Equal(t, middleware.CookieName, cookie.Name, "Unexpected no auth cookie")
	assert.NotZero(t, len(cookie.Value), "Unexpected empty token")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestLogin_NotRegistered(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	userNotRegistered := DefaultLoginBody
	userNotRegistered.Username = "testuser2"

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, userNotRegistered))
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Unexpected status code")

	var response dto.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.Equal(t, serverrors.ErrAccountIsNotRegistered.Error(), response.Message, "Unexpected error msg")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestLogin_InvalidPassword(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	userInvalidPassword := DefaultLoginBody
	userInvalidPassword.Password = "invalid"

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, userInvalidPassword))
	assert.Equal(t, http.StatusUnauthorized, w.Code, "Unexpected status code")

	var response dto.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.Equal(t, serverrors.ErrPasswordInvalid.Error(), response.Message, "Unexpected error msg")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestGetUsers_Success(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, DefaultLoginBody))
	require.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	cookie := w.Result().Cookies()[0]

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareGetUsersRequest(t, cookie))
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	var response dto.GetUsersResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")
	require.Equal(t, 1, len(response))

	assert.NotZero(t, response[0].ID, "Expected user ID to be non-zero")
	assert.Equal(t, "testuser", response[0].Username, "Unexpected username")
	assert.Equal(t, "test@example.com", response[0].Email, "Unexpected email")

	require.NoError(t, testenv.TestManager.CleanupDB())
}
