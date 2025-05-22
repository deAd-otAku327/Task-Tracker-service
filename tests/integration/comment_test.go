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

var DefaultCreateCommentBody = dto.PostCommentRequest{
	Text: "testcomment",
}

func TestCreateComment_Success(t *testing.T) {
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

	commentCreateWithTaskID := DefaultCreateCommentBody
	commentCreateWithTaskID.TaskID = task.ID

	testenv.TestManager.TestApp.Server.Handler.ServeHTTP(w, req.PrepareCreateCommentRequest(t, commentCreateWithTaskID, cookie))
	assert.Equal(t, http.StatusCreated, w.Code, "Unexpected status code")

	var response dto.CommentResponse
	err = json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err, "Failed to decode response")

	assert.NotZero(t, response.ID, "Expected comment ID to be non-zero")
	assert.Equal(t, user.ID, response.AuthorID, "Unexpected author ID")
	assert.Equal(t, "testcomment", response.Text, "Unexpected username")
	assert.NotZero(t, len(response.DateTime), "Unexpected dateTime")

	require.NoError(t, testenv.TestManager.CleanupDB())
}
