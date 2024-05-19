package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ranking/overall", overallRankingHandler)
	http.HandleFunc("/ranking/category", categoryRankingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
