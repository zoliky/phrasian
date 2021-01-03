// Copyright 2020 Zoltán Király. All rights reserved.

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// routes matches incoming requests against a list of registered routes
// and calls a handler for the route that matches the URL or other conditions.
func (app *application) routes() *mux.Router {
	// Create a new router instance
	router := mux.NewRouter()

	// Create a new file server for serving the assets
	assetsHandler := http.FileServer(http.Dir("../../ui/static"))
	assetsHandler = http.StripPrefix("/static/", assetsHandler)
	router.PathPrefix("/static/").Handler(assetsHandler)

	// Routes
	router.HandleFunc("/", app.home).Methods("GET")
	router.HandleFunc("/login", app.login).Methods("GET")
	router.HandleFunc("/phrase/add", app.add).Methods("GET")
	router.HandleFunc("/phrase/edit", app.edit).Queries("id", "{id}").Methods("GET")
	router.HandleFunc("/trash", app.trash).Methods("GET")

	return router
}
