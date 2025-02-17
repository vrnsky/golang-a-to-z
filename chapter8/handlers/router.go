package handlers

import (
	"golang-a-to-z/chapter8/api"
	"golang-a-to-z/chapter8/handlers/newgame"
	"net/http"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)
	return mux
}
