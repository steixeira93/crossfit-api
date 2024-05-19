package main

import (
	"log"
	"net/http"
)

func main() {
	setUpDatabase()

	http.HandleFunc("/athletes", getAllAthletesHandler)
	http.HandleFunc("/atheletes/add", addAthleteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
