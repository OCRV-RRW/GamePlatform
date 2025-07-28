package html

import (
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"text/template"
)

type Engine struct {
	dir       string
	cache     map[string]*template.Template
	hotReload bool
}

func New(dir string, hotReload bool) Engine {
	return Engine{
		dir:       dir,
		hotReload: hotReload,
	}
}

func (e *Engine) Load() error {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(e.dir, "*.page.html"))
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return err
		}

		ts, err = ts.ParseGlob(filepath.Join(e.dir, "*.layout.html"))
		if err != nil {
			return err
		}

		ts, err = ts.ParseGlob(filepath.Join(e.dir, "*.partial.html"))
		if err != nil {
			return err
		}

		cache[name] = ts
	}

	e.cache = cache

	return nil
}

func (e *Engine) Render(w io.Writer, name string, bind any, layouts ...string) error {
	if e.hotReload {
		err := e.Load()
		if err != nil {
			slog.Error("Couldn't load html cache")
		}
	}

	ts, ok := e.cache[name]
	if !ok {
		return fmt.Errorf("template %s does not exist", name)
	}

	err := ts.Execute(w, bind)
	return err
}
