package getstatus

import (
	"encoding/json"
	"golang-a-to-z/chapter8/internal/api"
	"log"
	"net/http"
)

func Handle(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue(api.GameId)

	if id == "" {
		http.Error(writer, "missing id of the game", http.StatusBadRequest)
		return
	}
	log.Printf("retrieve status of game with id: %v", id)

	apiGame := api.GameResponse{
		ID: id,
	}
	err := json.NewEncoder(writer).Encode(apiGame)
	if err != nil {
		http.Error(writer, "failed to serialize value", http.StatusInternalServerError)
		return
	}

}
