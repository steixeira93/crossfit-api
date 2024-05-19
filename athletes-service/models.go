package main

import (
	"database/sql"
	"log"
)

type Athlete struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Team      string `json:"team"`
	Box       string `json:"box"`
}

func setUpDatabase() {
	db, err := sql.Open("sqlite3", "./athletes.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS athletes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		firstName TEXT,
		lastName TEXT,
		age INTEGER,
		team TEXT,
		box TEXT
)`)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
}

func getAllAthletes() ([]Athlete, error) {

	db, err := sql.Open("sqlite3", "./athletes.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM athletes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	athletes := []Athlete{}
	for rows.Next() {
		var athlete Athlete
		err := rows.Scan(&athlete.ID, &athlete.FirstName, &athlete.LastName, &athlete.Age, &athlete.Team, &athlete.Box)
		if err != nil {
			return nil, err
		}
		athletes = append(athletes, athlete)
	}

	return athletes, nil
}

func addAthlete(a Athlete) error {
	db, err := sql.Open("sqlite3", "./athletes.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO athletes (firstName, lastName, age, team, box) VALUES (?, ?, ?, ?, ?)", a.FirstName, a.LastName, a.Age, a.Team, a.Box)
	if err != nil {
		return err
	}

	return nil
}
