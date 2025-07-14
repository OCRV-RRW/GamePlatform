package main

import (
	"context"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"allcran_wsx/gameplatform/cmd/web/config"
	"allcran_wsx/gameplatform/internal/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v5"
)

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	db             *database.Queries
	templatesCache map[string]*template.Template
}

func main() {
	//addr := flag.String("addr", ":4000", "HTTP address for start server")
	// dsn := flag.String("dsn", "web:qazXSW@!12wsxCDE£@23@tcp(db:3306)/gamebox?parseTime=True", "MySQL DB sourse config")
	// staticDir := flag.String("static-dir", "ui/static", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	config, err := config.Load()
	if err != nil {
		errorLog.Fatal("Couln't load config: ", err)
	}

	// Conecting to database
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.DatabaseURL)
	if err != nil {
		errorLog.Fatal("Couln't load config: ", err)
	} else {
		infoLog.Println("Successfully connect to database")
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	root, err := findRootDir()

	if err != nil {
		errorLog.Fatalf(err.Error())
	}

	templateCache, err := newTemplateCache(filepath.Join(root, "ui/html"))
	if err != nil {
		errorLog.Fatal(err.Error())
	}

	app := &application{
		infoLog:        infoLog,
		errorLog:       errorLog,
		db:             queries,
		templatesCache: templateCache,
	}

	infoLog.Printf("Server started on %s", config.Host)

	srv := &http.Server{
		Addr:     config.Host,
		ErrorLog: errorLog,
		Handler:  app.routes(&config.StatidDir),
	}
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// TODO: Delete
// func openDB(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}
//
// 	return db, nil
// }
