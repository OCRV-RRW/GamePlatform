package main

import (
	"net/http"
	"path/filepath"
)

func (app *application) routes(staticDir *string) *http.ServeMux {
	mux := http.NewServeMux()
	root, _ := findRootDir()
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(filepath.Join(root, *staticDir))})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/game", app.showGamePreview)
	mux.HandleFunc("/play", app.showGame)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	return mux
}
