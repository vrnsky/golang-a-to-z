package newgame

import (
	"encoding/json"
	"fmt"
	"golang-a-to-z/chapter8/internal/api"
	"golang-a-to-z/chapter8/internal/gordle"
	"golang-a-to-z/chapter8/internal/session"
	"log"
	"net/http"

	"github.com/oklog/ulid/v2"
)

type GameAdder interface {
	Add(game session.Game) error
}

// global variable that references each corpus
var corpora = map[string]string{
	"en": "./../../../corpus/english.txt",
	"he": "./../../../corpus/greek.txt",
	"cr": "./../../../corpus/cree.txt",
}

// Handler returns the handler for the game creation endpoint.
func Handler(adder GameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get(api.Lang)
		corpusPath, ok := corpora[lang]
		if !ok {
			// default to English
			corpusPath = corpora["en"]
		}

		game, err := CreateGame(adder, corpusPath)
		if err != nil {
			log.Printf("unable to create a new game: %s", err)
			http.Error(w, "failed to create a new game", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		apiGame := api.ToGameResponse(game)
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}
	}
}

const maxAttempts = 5

func CreateGame(adder GameAdder, corpusPath string) (session.Game, error) {
	corpus, err := gordle.ReadCorpus(corpusPath)
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to read corpus: %w", err)
	}

	game, err := gordle.New(corpus)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to create a new gordle game")
	}

	g := session.Game{
		ID:           session.GameId(ulid.Make().String()),
		Gordle:       *game,
		AttemptsLeft: maxAttempts,
		Guesses:      []session.Guess{},
		Status:       session.StatusPlaying,
	}

	err = adder.Add(g)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to save the new game")
	}

	return g, nil
}
