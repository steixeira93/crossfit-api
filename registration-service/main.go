package main

import (
	"log"
	"net/http"
)

func main() {
	setUpDatabase()

	http.HandleFunc("/registrations", getAllRegistrationsHandler)
	http.HandleFunc("/registrations/add", addRegistrationHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
