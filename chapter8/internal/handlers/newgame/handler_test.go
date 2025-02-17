package newgame_test

import (
	"golang-a-to-z/chapter8/internal/handlers/newgame"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/games", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	newgame.Handle(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	// assert.JSONEq(t, `{"id":"","attempts_left":0,"guesses":null,"WordLength":0,"status":""}`, recorder.Body.String())
}
