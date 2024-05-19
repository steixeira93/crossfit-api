package main

import (
	"database/sql"
	"log"
)

type Event struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Date        string `json:"date"`
}

func setUpDatabase() {
	db, err := sql.Open("sqlite3", "./events.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS events (
		id TEXT PRIMARY KEY,
		name TEXT,
		description TEXT,
		location TEXT,
		date TEXT
)`)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
}

func getAllEvents() ([]Event, error) {
	var events []Event

	db, err := sql.Open("sqlite3", "./events.db")
	if err != nil {
		return events, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event Event
		err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Date)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func addEvent(event Event) error {
	db, err := sql.Open("sqlite3", "./events.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO events (id, name, description, location, date) VALUES (?, ?, ?, ?, ?)", event.ID, event.Name, event.Description, event.Location, event.Date)
	if err != nil {
		return err
	}
	return nil
}
