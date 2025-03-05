package getstatus

import (
	"encoding/json"
	"errors"
	"golang-a-to-z/chapter8/internal/api"
	"golang-a-to-z/chapter8/internal/repository"
	"golang-a-to-z/chapter8/internal/session"
	"log"
	"net/http"
)

// gameFinder finds a game in the storage.
type gameFinder interface {
	Find(id session.GameId) (session.Game, error)
}

// Handler returns the handler for the game creation endpoint.
func Handler(finder gameFinder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		id := req.PathValue(api.GameId)
		if id == "" {
			http.Error(w, "missing the id of the game", http.StatusNotFound)
			return
		}

		game, err := finder.Find(session.GameId(id))
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				http.Error(w, "this game does not exist yet", http.StatusNotFound)
				return
			}

			log.Printf("cannot fetch game %s: %s", id, err)
			http.Error(w, "failed to fetch game", http.StatusInternalServerError)
			return
		}

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			http.Error(w, "failed to write response", http.StatusInternalServerError)
			return
		}
	}
}
