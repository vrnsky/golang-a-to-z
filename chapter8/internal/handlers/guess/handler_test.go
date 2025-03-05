package guess_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang-a-to-z/chapter8/internal/api"
	"golang-a-to-z/chapter8/internal/gordle"
	"golang-a-to-z/chapter8/internal/handlers/guess"
	"golang-a-to-z/chapter8/internal/session"
)

func TestHandle(t *testing.T) {
	game, _ := gordle.New([]string{"pocket"})
	handle := guess.Handler(successGameGuesserStub{session.Game{
		ID:           "123456",
		Gordle:       *game,
		AttemptsLeft: 5,
		Guesses:      nil,
		Status:       session.StatusPlaying,
	}})

	req, err := http.NewRequest(http.MethodPost, "/games/123456", strings.NewReader(`{"guess":"pocket"}`))
	require.NoError(t, err)

	// add path parameters
	req.SetPathValue(api.GameId, "123456")

	recorder := httptest.NewRecorder()

	handle(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":4,"guesses":[{"word":"pocket", "feedback":"++++++"}],"word_length":6,"status":"Won"}`, recorder.Body.String())
}

type successGameGuesserStub struct {
	game session.Game
}

func (g successGameGuesserStub) Find(id session.GameId) (session.Game, error) {
	return g.game, nil
}

func (g successGameGuesserStub) Update(game session.Game) error {
	return nil
}
