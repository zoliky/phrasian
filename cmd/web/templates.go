// Copyright 2020 Zoltán Király. All rights reserved.

package main

import (
	"html/template"
	"path/filepath"

	"github.com/zoliky/phrasian/pkg/models"
)

// templateData acts as a holding structure for any dynamic data
// that is passed to the HTML templates
type templateData struct {
	Phrase  *models.Phrase
	Phrases *[]models.Phrase
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	// Get a slice of all filepaths with the extension '.page.gohtml'
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.gohtml"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract the file name (like 'home.page.gohtml') from the
		// full file path and assign it to the name variable
		name := filepath.Base(page)

		// Parse the page template file in to a template set
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add '.layout.gohtml' templates to the template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.gohtml"))
		if err != nil {
			return nil, err
		}

		// Add '.partial.gohtml' templates to the template set
		//ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.gohtml"))
		//if err != nil {
		//	return nil, err
		//}

		cache[name] = ts
	}

	// Return the map
	return cache, nil
}
