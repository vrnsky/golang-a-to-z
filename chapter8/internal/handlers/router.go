package handlers

import (
	"golang-a-to-z/chapter8/internal/api"
	"golang-a-to-z/chapter8/internal/handlers/getstatus"
	"golang-a-to-z/chapter8/internal/handlers/newgame"
	"net/http"
)

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)
	return mux
}

// NewRouter return a router that listesn for requests to the following endpoints:
// - Create a new game
//
// The provided router is ready to serve.
func NewRouter() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle)
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handle)
	return r
}
