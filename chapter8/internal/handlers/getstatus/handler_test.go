package getstatus_test

import (
	"golang-a-to-z/chapter8/internal/api"
	"golang-a-to-z/chapter8/internal/handlers/getstatus"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/games/", nil)
	require.NoError(t, err)

	req.SetPathValue(api.GameId, "123456")

	recorder := httptest.NewRecorder()

	getstatus.Handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	// assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}
