// Copyright 2020 Zoltán Király. All rights reserved.

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// home handler
func (app *application) home(w http.ResponseWriter, req *http.Request) {
	phrases, err := app.phrases.GetAll("published", "1")
	if err != nil {
		app.serverError(w, err)
	}

	d := &templateData{
		Phrases: phrases,
	}

	app.render(w, req, "home.page.gohtml", d)
}

// login handler
func (app *application) login(w http.ResponseWriter, req *http.Request) {
	app.render(w, req, "login.page.gohtml", &templateData{})
}

// add handler
func (app *application) add(w http.ResponseWriter, req *http.Request) {

}

// edit handler
func (app *application) edit(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	phrase, err := app.phrases.GetPhrase(params["id"])
	if err != nil {
		app.serverError(w, err)
	}

	d := &templateData{
		Phrase: phrase,
	}

	app.render(w, req, "edit.page.gohtml", d)
}

// trash handler
func (app *application) trash(w http.ResponseWriter, req *http.Request) {
	phrases, err := app.phrases.GetAll("trash", "1")
	if err != nil {
		app.serverError(w, err)
	}

	d := &templateData{
		Phrases: phrases,
	}

	app.render(w, req, "trash.page.gohtml", d)
}
