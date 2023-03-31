package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// TemplateRenderer parses the template file and Executes it using the ResponseWriter
func TemplateRenderer(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// create a new buffer and run execute writing to the buffer
	// this allows us to catch errors within the template itself
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)

	if err != nil {
		log.Println(err)
	}
	// using the buffer write to the responseWriter as the target
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// createTemplateCache parses all pages and layouts and stores them by name in a map cache
func createTemplateCache() (map[string]*template.Template, error) {
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
