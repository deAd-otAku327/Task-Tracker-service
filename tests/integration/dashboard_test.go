package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-tracker-service/internal/types/dto"
	req "task-tracker-service/tests/integration/requests"
	"task-tracker-service/tests/testenv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var DefaultCreateDashboardBody = dto.PostDashboardsCreateRequest{
	Title:       "testing",
	Description: nil,
}

func TestCreateDashboard_BasicSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, DefaultLoginBody))
	require.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	cookie := w.Result().Cookies()[0]

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareCreateDashboardRequest(t, DefaultCreateDashboardBody, cookie))
	assert.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var response dto.DashboardResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.NotZero(t, response.ID, "Expected dashboard ID to be non-zero")
	assert.Equal(t, "testing", response.Title, "Unexpected username")
	assert.Nil(t, response.Description, "Unexpected description")
	assert.NotZero(t, len(response.UpdatedAt), "Unexpected updatedAt time")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestUpdateDashboard_BasicSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, DefaultLoginBody))
	require.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	cookie := w.Result().Cookies()[0]

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareCreateDashboardRequest(t, DefaultCreateDashboardBody, cookie))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var response dto.DashboardResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	newTitle, newDesc := "newtitle", "newdesc"
	updateDashboardBody := dto.PostDashboardsUpdateRequest{
		BoardID:     response.ID,
		Title:       &newTitle,
		Description: &newDesc,
	}

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareUpdateDashboardRequest(t, updateDashboardBody, cookie))
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestGetDashboardSummary_BasicSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var user dto.UserResponse
	err := json.NewDecoder(w.Body).Decode(&user)
	require.NoError(t, err, "Failed to decode response")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, DefaultLoginBody))
	require.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	cookie := w.Result().Cookies()[0]

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareCreateDashboardRequest(t, DefaultCreateDashboardBody, cookie))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var dashboard dto.DashboardResponse
	err = json.NewDecoder(w.Body).Decode(&dashboard)
	require.NoError(t, err, "Failed to decode response")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareGetDashboardSummaryRequest(t, dashboard.ID, cookie))
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	var response dto.GetDashboardSummaryResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.Equal(t, response.Dashboard, &dashboard, "Unexpected dashboard")
	assert.Equal(t, 1, len(response.Admins), "Unexpected admin count")
	assert.Equal(t, response.Admins[0], &user, "Unexpected admin")

	assert.Equal(t, 0, len(response.Tasks), "Unexpected tasks count")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestAddBoardAdmin_Success(t *testing.T) {
	w := httptest.NewRecorder()

	anotherUserRegister := dto.PostUsersRegisterRequest{
		Email:    "another@gmail.com",
		Username: "givemeadmin",
		Password: "testpass123",
	}

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, anotherUserRegister))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var anotherUser dto.UserResponse
	err := json.NewDecoder(w.Body).Decode(&anotherUser)
	require.NoError(t, err, "Failed to decode response")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, DefaultLoginBody))
	require.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	cookie := w.Result().Cookies()[0]

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareCreateDashboardRequest(t, DefaultCreateDashboardBody, cookie))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var dashboard dto.DashboardResponse
	err = json.NewDecoder(w.Body).Decode(&dashboard)
	require.NoError(t, err, "Failed to decode response")

	w = httptest.NewRecorder()

	addBoardAdmin := dto.PostDashboardsAdminRequest{
		BoardID: dashboard.ID,
		UserID:  anotherUser.ID,
	}

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareAddDashboardAdminRequest(t, addBoardAdmin, cookie))
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	require.NoError(t, testenv.TestManager.CleanupDB())
}
