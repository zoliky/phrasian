// Copyright 2020 Zoltán Király. All rights reserved.

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/alexedwards/scs/postgresstore"

	"github.com/alexedwards/scs/v2"
	"github.com/zoliky/phrasian/pkg/models/postgresql"

	_ "github.com/lib/pq"
)

// Define an application struct to hold application-wide dependencies
type application struct {
	infoLog        *log.Logger
	errorLog       *log.Logger
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager
	phrases        *postgresql.PhraseModel
}

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "zoliky"
	password = ""
	dbname   = "testdb"
)

// openDB returns a sql.DB connection pool for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Define a new command-line flag for the HTTP network address
	addr := flag.String("addr", ":8080", "HTTP network address")

	// Define a new command-line flag for the PostgreSQL DSN string
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	dsn := flag.String("dsn", psqlInfo, "PostgreSQL data source name")

	flag.Parse()

	// Create custom loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Create a database connection
	infoLog.Println("Connecting to database")
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Initialize a new template cache
	templateCache, err := newTemplateCache("../../ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize a new session manager
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(db)

	// Initialize a new instance of application containing the dependencies
	app := &application{
		infoLog:        infoLog,
		errorLog:       errorLog,
		templateCache:  templateCache,
		sessionManager: sessionManager,
		phrases:        &postgresql.PhraseModel{DB: db},
	}

	// Define custom parameters for the HTTP server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  sessionManager.LoadAndSave(app.routes()),
	}

	// Start the HTTP server
	infoLog.Printf("Starting server on %s", *addr)
	errorLog.Fatal(srv.ListenAndServe())
}
