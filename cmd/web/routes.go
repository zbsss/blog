package main

import (
	"net/http"

	"github.com/justinas/alice"
	ui "github.com/zbsss/blog/public"
)

func (a *app) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", a.health)

	fs := http.FileServer(http.FS(ui.Files))
	mux.Handle("/blog/", http.StripPrefix("/blog/", fs))

	standard := alice.New(a.recoverPanic, a.logRequest, secureHeaders)

	return standard.Then(mux)
}

func (a *app) health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
