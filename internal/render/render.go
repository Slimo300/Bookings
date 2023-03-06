package render

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/Slimo300/Bookings/internal/config"
	"github.com/Slimo300/Bookings/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewRenderer assigns app config to render engine
func NewRenderer(a *config.AppConfig) {
	app = a
}

// AddDefaultData is helper for passing data that will be necessary to
// every template
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	return td
}

// Template renders a template
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		return errors.New("Could not get template from cache")
	}

	td = AddDefaultData(td, r)

	if err := t.Execute(w, td); err != nil {
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templ, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		templ, err = templ.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return nil, err
		}

		myCache[name] = templ
	}

	return myCache, nil
}
