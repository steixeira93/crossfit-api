package main

import (
	"encoding/json"
	"net/http"
)

func overallRankingHandler(w http.ResponseWriter, r *http.Request) {
	overallRanking, err := calculateOverallRanking()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(overallRanking)
}

func categoryRankingHandler(w http.ResponseWriter, r *http.Request) {
	category := r.URL.Query().Get("category")
	if category == "" {
		http.Error(w, "missing category query parameter", http.StatusBadRequest)
		return
	}

	categoryRanking, err := calculateCategoryRanking(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categoryRanking)
}
