package main

import (
	"log"
	"myserver/app"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := app.NewRouter()

	http.ListenAndServe(":"+port, router)
}
