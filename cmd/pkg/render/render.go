package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"banachtech.github.com/weirdville/cmd/pkg/config"
	"banachtech.github.com/weirdville/cmd/pkg/models"
)

const templatePath = "templates"

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

// common data to be included in every page
func AddCommonData(td *models.TemplateData) *models.TemplateData {
	// add data here
	// ...
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	var err error

	// use cache only if this flag is set
	if app.UseCache {
		// read from template cache
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplate()
		if err != nil {
			log.Fatal("error building template cache")
		}
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("template does not exist")
	}

	// Add common data
	td = AddCommonData(td)

	// a buffer technique
	buf := new(bytes.Buffer)
	err = t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	// execute template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// more sophisticated set-up
func CreateTemplate() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Get a slice of all filepaths with extension '.page.html'
	// These are the template files we want to parse
	pages, err := filepath.Glob(filepath.Join(templatePath, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, p := range pages {
		// extract file name
		name := filepath.Base(p)
		// parse page template p
		ts, err := template.New(name).ParseFiles(p)
		if err != nil {
			return nil, err
		}
		// add layout templates to the parsed set
		ts, err = ts.ParseGlob(filepath.Join(templatePath, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}

/*
// Simpler implementation of cache
// template cache
var tmplCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	var t *template.Template
	if _, ok := tmplCache[tmpl]; !ok {
		// not in cache, create new
		err := createTemplate(tmpl)
		if err != nil {
			log.Println(err)
		}
	}
	t = tmplCache[tmpl]
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}


// cache parsed template files in a map for reuse
func createTemplate(t string) error {
	files := []string{templatePath + t, templatePath + "base.layout.tmpl"}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}
	tmplCache[t] = tmpl
	return nil
}
*/
