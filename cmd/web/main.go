package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ChienDuyNguyen1702/Go_Learning/pkg/config"
	"github.com/ChienDuyNguyen1702/Go_Learning/pkg/handlers"
	"github.com/ChienDuyNguyen1702/Go_Learning/pkg/render"
)

// var portNumber = ":8080"
const portNumber = ":8080"

// Home is the home page handler

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = true
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
