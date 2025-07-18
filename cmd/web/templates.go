package main

import (
	"allcran_wsx/gameplatform/internal/database"
	"html/template"
	"path/filepath"
)

type GamePreview struct {
	preivew []database.Preview
	database.Game
}

type HomeTemplateData struct {
	Games []database.Game
}

type GamePreviewData struct {
	Preview []database.Preview
	database.Game
}

type ShowTemplateData struct {
	Game GamePreviewData
}

type GameTemplateData struct {
	Game database.Game
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.html"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
