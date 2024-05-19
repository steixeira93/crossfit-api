package main

import (
	"encoding/json"
	"net/http"
)

func getAllRegistrationsHandler(w http.ResponseWriter, r *http.Request) {
	registrations, err := getAllRegistrations()
	if err != nil {
		http.Error(w, "Error getting registrations", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(registrations)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func addRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var registration Registration
	err := json.NewDecoder(r.Body).Decode(&registration)
	if err != nil {
		http.Error(w, "Error decoding registration", http.StatusBadRequest)
		return
	}

	err = addRegistration(registration)
	if err != nil {
		http.Error(w, "Error adding registration", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
