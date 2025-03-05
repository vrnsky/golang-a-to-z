package handlers

import (
	"golang-a-to-z/chapter8/internal/api"
	"golang-a-to-z/chapter8/internal/handlers/getstatus"
	"golang-a-to-z/chapter8/internal/handlers/guess"
	"golang-a-to-z/chapter8/internal/handlers/newgame"
	"golang-a-to-z/chapter8/internal/repository"
	"net/http"
)

func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()

	// Register each endpoint.
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handler(db))    // curl -X POST -v http://localhost:9090/games
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handler(db)) // curl -X GET -v http://localhost:9090/games/1682279480
	r.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handler(db))         // curl -X PUT -v http://localhost:9090/games/1682279480 -d '{"guess":"faune"}'

	return r
}
