package main

import (
	"net/http"
	"path/filepath"
)

func (app *application) routes(staticDir *string) *http.ServeMux {
	mux := http.NewServeMux()
	root, _ := findRootDir()
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(filepath.Join(root, *staticDir))})
	mux.Handle("/gameplatform/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/gameplatform/gameplatform", app.home)
	mux.HandleFunc("/gameplatform/game", app.showGamePreview)
	mux.HandleFunc("/gameplatform/play", app.showGame)
	return mux
}
