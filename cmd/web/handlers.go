package main

import (
	"database/sql"
	"errors"

	"net/http"

	"github.com/google/uuid"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/gameplatform" {
		http.NotFound(w, r)
		return
	}

	s, err := app.db.GetGames(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &HomeTemplateData{Games: s}
	app.infoLog.Println(data)

	app.render(w, "home.page.html", data)
}

func (app *application) showGamePreview(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.notFound(w)
		return
	}

	id_uuid, err := uuid.Parse(id)
	if err != nil {
		app.serverError(w, err)
	}

	game, err := app.db.GetGameByID(r.Context(), id_uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	preview, err := app.db.GetGamePreview(r.Context(), game.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	gamePreview := GamePreviewData{
		Preview: preview,
		Game:    game,
	}

	data := &ShowTemplateData{gamePreview}
	app.infoLog.Println(data)
	app.render(w, "show.page.html", data)
}

func (app *application) showGame(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		app.notFound(w)
		return
	}

	ctx := r.Context()
	id_uuid, err := uuid.Parse(id)

	if err != nil {
		app.clientError(w, 400)
	}

	game_data, err := app.db.GetGameByID(ctx, id_uuid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := &GameTemplateData{Game: game_data}
	app.render(w, "game.page.html", data)
}

// func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodPost)
// 		app.clientError(w, http.StatusMethodNotAllowed)
// 		return
// 	}
//
// 	if err := r.ParseForm(); err != nil {
// 		app.clientError(w, http.StatusMethodNotAllowed)
// 		return
// 	}
//
// 	title := "История про улитку"
// 	content := "Улитка выползла из раковины,\nвытянула рожки,\nи опять подобрала их."
// 	expires := "7"
//
// 	id, err := app.games.Insert(title, content, expires)
// 	if err != nil {
// 		app.serverError(w, err)
// 	}
//
// 	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
// }
