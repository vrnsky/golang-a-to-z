package getstatus_test

import (
	"golang-a-to-z/chapter8/internal/api"
	"golang-a-to-z/chapter8/internal/handlers/getstatus"
	"golang-a-to-z/chapter8/internal/session"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	handle := getstatus.Handler(gameFinderStub{session.Game{ID: "123456"}, nil})

	req, err := http.NewRequest(http.MethodPost, "/games/123456", nil)
	require.NoError(t, err)

	// add path parameters
	req.SetPathValue(api.GameId, "123456")

	recorder := httptest.NewRecorder()

	handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}

type gameFinderStub struct {
	game session.Game
	err  error
}

func (g gameFinderStub) Find(_ session.GameId) (session.Game, error) {
	return g.game, g.err
}
