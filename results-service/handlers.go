package main

import (
	"encoding/json"
	"net/http"
)

func addResultHandler(w http.ResponseWriter, r *http.Request) {
	var result Result
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = addResultToDB(result)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Result added successfully"))
}

func getAllResultsHandler(w http.ResponseWriter, r *http.Request) {
	results, err := getAllResultsFromDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResults, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResults)
}
