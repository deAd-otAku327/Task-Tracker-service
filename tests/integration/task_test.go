package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-tracker-service/internal/types/dto"
	"task-tracker-service/internal/types/enum"
	req "task-tracker-service/tests/integration/requests"
	"task-tracker-service/tests/testenv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var DefaultCreateTaskBody = dto.PostTasksCreateRequest{
	Title:         "testing",
	Description:   nil,
	AssignieID:    nil,
	LinkedBoardID: nil,
}

func TestCreateTask_BasicSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, DefaultLoginBody))
	require.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	cookie := w.Result().Cookies()[0]

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareCreateTaskRequest(t, DefaultCreateTaskBody, cookie))
	assert.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var response dto.TaskResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.NotZero(t, response.ID, "Expected task ID to be non-zero")
	assert.Equal(t, "testing", response.Title, "Unexpected username")
	assert.Nil(t, response.Description, "Unexpected description")
	assert.Equal(t, enum.StatusCreated.String(), response.Status, "Unexpected status")
	assert.Nil(t, response.Assignie, "Unexpected assignie")
	assert.Nil(t, response.Board, "Unexpected board")
	assert.NotZero(t, len(response.UpdatedAt), "Unexpected updatedAt time")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestUpdateTask_BasicSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, DefaultLoginBody))
	require.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	cookie := w.Result().Cookies()[0]

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareCreateTaskRequest(t, DefaultCreateTaskBody, cookie))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var response dto.TaskResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	newTitle, newDesc := "newtitle", "newdesc"
	updateTaskBody := dto.PostTasksUpdateRequest{
		TaskID:      response.ID,
		Title:       &newTitle,
		Description: &newDesc,
	}

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareUpdateTaskRequest(t, updateTaskBody, cookie))
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestGetTasks_BasicSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareRegisterRequest(t, DefaultRegisterBody))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareLoginRequest(t, DefaultLoginBody))
	require.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	cookie := w.Result().Cookies()[0]

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareCreateTaskRequest(t, DefaultCreateTaskBody, cookie))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareGetTasksRequest(t, cookie))
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	var response dto.GetTasksResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	require.Zero(t, len(response), "Unexpected tasks count")

	require.NoError(t, testenv.TestManager.CleanupDB())
}

func TestGetTaskSummary_BasicSuccess(t *testing.T) {
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

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareCreateTaskRequest(t, DefaultCreateTaskBody, cookie))
	require.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var task dto.TaskResponse
	err = json.NewDecoder(w.Body).Decode(&task)
	require.NoError(t, err, "Failed to decode response")

	w = httptest.NewRecorder()

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareGetTaskSummaryRequest(t, task.ID, cookie))
	assert.Equal(t, http.StatusOK, w.Code, "Unexpected status code")

	var response dto.GetTaskSummaryResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.Equal(t, response.Task, &task, "Unexpected task")
	assert.Equal(t, response.Author, &user, "Unexpected author")

	assert.Nil(t, response.Comments, "Unexpected comments")
	assert.Nil(t, response.Assignie, "Unexpected assignie")
	assert.Nil(t, response.LinkedBoard, "Unexpected linked board")

	require.NoError(t, testenv.TestManager.CleanupDB())
}
