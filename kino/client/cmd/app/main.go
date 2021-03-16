package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type apis struct {
	films  string
}

type App struct {
	apis apis
}

var token string

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	filmsAPI := flag.String("filmsAPI", "http://localhost:8001/film/", "Films API")
	flag.Parse()

	app := &App{
		apis: apis{
			films: *filmsAPI,
		},
	}

	fmt.Println("token dai mne, a ya tebe site)")
	fmt.Scanln(&token)

	log.Printf("Server listening on %s", *addr)
	err := http.ListenAndServe(*addr, app.Routes())
	log.Fatal(err)
}
