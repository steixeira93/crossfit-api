package main

import "database/sql"

type Result struct {
	ID        int `json:"id"`
	EventID   int `json:"event_id"`
	AthleteID int `json:"athlete_id"`
	Score     int `json:"score"`
}

func setUpDatabase() {
	db, err := sql.Open("sqlite3", "./results.db")
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS results (
		id INTEGER PRIMARY KEY,
		event_id INTEGER,
		athlete_id INTEGER,
		score INTEGER
)`)
	if err != nil {
		log.Fatal("Error creating table: ", err)
	}
}

func addResultToDB(result Result) error {
	db, err := sql.Open("sqlite3", "./results.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO results (
		event_id, athlete_id, score)
		VALUES (?, ?, ?)",
		result.EventID, result.AthleteID, result.Score)
	if err != nil {
		return err
	}

	return nil
}

func getAllResultsFromDB() ([]Result, error) {
	var results []Result

	db, err := sql.Open("sqlite3", "./results.db")
	if err != nil {
		return results, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM results")
	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var result Result
		err = rows.Scan(&result.ID, &result.EventID, &result.AthleteID, &result.Score)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}

	return results, nil
}
