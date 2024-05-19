package main

import "net/http"

func main() {
	setUpDatabase()

	http.HandleFunc("/results", getAllResultsHandler)
	http.HandleFunc("/results/add", addResultHandler)

	http.ListenAndServe(":8080", nil)
}
