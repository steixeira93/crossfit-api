package main

import "net/http"

func main() {
	setUpDatabase()

	http.HandleFunc("/events", getAllEventsHandler)
	http.HandleFunc("/events/add", addEventHandler)
}
