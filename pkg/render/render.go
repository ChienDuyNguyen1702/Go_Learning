package render

import (
	"bytes"
	"github.com/ChienDuyNguyen1702/Go_Learning/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate render a template to a html, show in url with port bla bla
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing  template to browser")
	}
}
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	// get all the file names .tmpl in ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// range through all file names *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

//
//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	//check to see if the template is already in the cache
//	_, inMap := tc[t]
//	if !inMap {
//		// need to create new template
//		log.Printf("creating new template and adding to cache")
//		err = createTemplates(t)
//		if err != nil {
//			log.Println(err)
//		}
//	} else {
//		log.Printf("Using cache template")
//	}
//
//	tmpl = tc[t]
//	err = tmpl.Execute(w, nil)
//	if err != nil {
//		log.Println(err)
//	}
//}
//
//func createTemplates(t string) error {
//	templates := []string{
//		fmt.Sprintf("./templates/%s", t), "./templates/base.layout.tmpl",
//	}
//	//parse the templates
//	tmpl, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//	// add template to cache
//	tc[t] = tmpl
//	return nil
//}
