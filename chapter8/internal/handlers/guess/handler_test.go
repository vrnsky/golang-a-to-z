package guess_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"golang-a-to-z/chapter8/internal/api"
	"golang-a-to-z/chapter8/internal/handlers/guess"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/games/", strings.NewReader(`{"guess":"pocket"}`))
	require.NoError(t, err)

	// add path parameters
	req.SetPathValue(api.GameId, "123456")

	recorder := httptest.NewRecorder()

	guess.Handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":null,"word_length":0,"status":""}`, recorder.Body.String())
}
