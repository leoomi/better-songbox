package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	_ "embed"

	"github.com/leoomi/better-songbox/internal/api"
	"github.com/leoomi/better-songbox/internal/ui"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	ui.SetupRoutes()
	app.RunWhenOnBrowser()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server listening on http://localhost:" + port)
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css",
		},
		Scripts: []string{
			"https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js",
		},
	})

	http.HandleFunc("/api/search", api.SearchHandler)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error starting server:", err)
	}
}
