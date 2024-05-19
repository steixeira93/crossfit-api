package main

import (
	"encoding/json"
	"net/http"
)

func getAllAthletesHandler(w http.ResponseWriter, r *http.Request) {
	athletes, err := getAllAthletes()
	if err != nil {
		http.Error(w, "Error getting athletes", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(athletes)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func addAthleteHandler(w http.ResponseWriter, r *http.Request) {
	var athlete Athlete
	err := json.NewDecoder(r.Body).Decode(&athlete)
	if err != nil {
		http.Error(w, "Error decoding athlete", http.StatusBadRequest)
		return
	}

	err = addAthlete(athlete)
	if err != nil {
		http.Error(w, "Error adding athlete", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
