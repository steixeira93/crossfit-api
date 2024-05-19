package main

import (
	"log"
	"net/http"
)

func main() {
	setUpDatabase()

	http.HandleFunc("/events", getAllEventsHandler)
	http.HandleFunc("/events/add", addEventHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
