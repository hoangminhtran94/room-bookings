package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/hoangminhtran94/room-bookings/pkg/config"
	"github.com/hoangminhtran94/room-bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(res http.ResponseWriter, templ string, td *models.TemplateData) {
	var caches map[string]*template.Template
	if app.UseCache {
		caches = app.TemplateCache
	} else {
		caches, _ = CreateTemplateCache()
	}

	//getRequested template cache from caches
	cache, ok := caches[templ]
	if !ok {
		log.Fatal("Could not get template from caches")
	}
	td = AddDefaultData(td)
	err := cache.Execute(res, td)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		cache, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			cache, err = cache.ParseGlob("./templates/*.layout.html")
		}
		if err != nil {
			return myCache, err
		}
		myCache[name] = cache
	}
	return myCache, nil
}
