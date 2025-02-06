package main

import (
	"allcran_wsx/gameplatform/pkg/models/mysql"
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	games          *mysql.GameModel
	templatesCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP address for start server")
	dsn := flag.String("dsn", "web:qazXSW@!12wsxCDE£@23@tcp(db:3306)/gamebox?parseTime=True", "MySQL DB sourse config")
	staticDir := flag.String("static-dir", "ui/static", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	db, err := openDB(*dsn)
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		errorLog.Output(2, trace)
	}

	defer db.Close()

	root, err := findRootDir()
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		errorLog.Output(2, trace)
	}

	templateCache, err := newTemplateCache(filepath.Join(root, "ui/html"))
	if err != nil {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		errorLog.Output(2, trace)
	}

	app := &application{
		infoLog:        infoLog,
		errorLog:       errorLog,
		games:          &mysql.GameModel{DB: db},
		templatesCache: templateCache,
	}

	infoLog.Printf("Server started on %s", *addr)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(staticDir),
	}
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
