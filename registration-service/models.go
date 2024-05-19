package main

import (
	"database/sql"
	"fmt"
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

	// TODO: Here I need to implement the logic to check if
	// the athlete is qualified based on criteria such as age, skill level, etc.

	// If the athlete is qualified, then we can proceed to add the registration
	_, err = db.Exec("INSERT INTO registrations (id, athlete_id, event_id) VALUES (?, ?, ?)", registration.ID, registration.AthleteID, registration.EventID)
	if err != nil {
		return err
	}
	return nil
}

func updateRegistration(registrationID, newEventID int) error {
	db, err := sql.Open("sqlite3", "./registrations.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE registrations SET event_id = ? WHERE id = ?", newEventID, registrationID)
	if err != nil {
		return err
	}
	return nil
}

func cancelRegistration(registrationID int) error {
	db, err := sql.Open("sqlite3", "./registrations.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM registrations WHERE id = ?", registrationID)
	if err != nil {
		return err
	}
	return nil
}

func generateConfirmation(athleteID, eventID int) string {
	confirmation := fmt.Sprintf("Confirmation of registration for the athlete %d in the event %d", athleteID, eventID)
	return confirmation
}

// TODO: create a function to payment processing
