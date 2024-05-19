package main

import (
	"database/sql"
	"log"
)

type Registration struct {
	ID        string `json:"id"`
	AthleteID int    `json:"athlete_id"`
	EventID   int    `json:"event_id"`
}

func setUpDatabase() {
	db, err := sql.Open("sqlite3", "./registrations.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS registrations (
		id TEXT PRIMARY KEY,
		athlete_id INTEGER,
		event_id INTEGER
)`)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
}

func getAllRegistrations() ([]Registration, error) {
	var registrations []Registration

	db, err := sql.Open("sqlite3", "./registrations.db")
	if err != nil {
		return registrations, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM registrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var registration Registration
		err = rows.Scan(&registration.ID, &registration.AthleteID, &registration.EventID)
		if err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}
	return registrations, nil
}

func addRegistration(registration Registration) error {
	db, err := sql.Open("sqlite3", "./registrations.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO registrations (id, athlete_id, event_id) VALUES (?, ?, ?)", registration.ID, registration.AthleteID, registration.EventID)
	if err != nil {
		return err
	}
	return nil
}
