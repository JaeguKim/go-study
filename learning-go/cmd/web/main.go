package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaegoo-kim/go-course/pkg/config"
	"github.com/jaegoo-kim/go-course/pkg/handlers"
	"github.com/jaegoo-kim/go-course/pkg/render"
)

const portNumber = ":8080" // can't be changed by other app

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
