package main

import (
	"encoding/json"
	"net/http"
)

func getAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := getAllEvents()
	if err != nil {
		http.Error(w, "Error getting events", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(events)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func addEventHandler(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Error decoding event", http.StatusBadRequest)
		return
	}

	err = addEvent(event)
	if err != nil {
		http.Error(w, "Error adding event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
