package render

import (
	"bytes"
	"github.com/tnaucoin/go-web-app/pkg/config"
	"github.com/tnaucoin/go-web-app/pkg/models"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var app *config.AppConfig

// NewTemplates assigns the template cache to render
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds default data to the template data
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// TemplateRenderer parses the template file and Executes it using the ResponseWriter
func TemplateRenderer(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	// Enables or disables the use of the tmpl cache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		var err error
		// Recreate the cache on every call
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("could not create the template cache")
		}
	}
	// get template cache from app config
	// and requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from cache")
	}
	// Assign default data
	td = AddDefaultData(td)
	// create a new buffer and run execute writing to the buffer
	// this allows us to catch errors within the template itself
	buf := new(bytes.Buffer)
	err := t.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}
	// using the buffer write to the responseWriter as the target
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// CreateTemplateCache parses all pages and layouts and stores them by name in a map cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	// get all page files from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return cache, err
	}

	// range through all page files
	for _, page := range pages {
		// get filename from path
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		// find all the layout.tmpl files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}
		// For each page parse the found layout.tmpl and associate them with the current page
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		// store the final parsed tmpl file in the cache using the filename
		cache[name] = ts
	}

	return cache, nil
}
