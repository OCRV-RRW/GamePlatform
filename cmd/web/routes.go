package main

import (
	"net/http"
	"path/filepath"
)

func (app *application) routes(staticDir *string) *http.ServeMux {
	mux := http.NewServeMux()
	root, _ := findRootDir()
	staticPath := filepath.Join(root, *staticDir, "/")
	app.infoLog.Printf("Static path %v", staticPath)
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(staticPath)})
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/gameplatform", app.home)
	mux.HandleFunc("/gameplatform/game", app.showGamePreview)
	mux.HandleFunc("/gameplatform/play", app.showGame)
	return mux
}
