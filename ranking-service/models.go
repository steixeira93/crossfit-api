package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Ranking struct {
	AthleteID  int    `json:"athlete_id"`
	Category   string `json:"category"`
	TotalScore int    `json:"total_score"`
	Rank       int    `json:"rank"`
}

func calculateOverallRanking() ([]Ranking, error) {
	var overallRanking []Ranking

	db, err := sql.Open("sqlite3", "./ranking.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
        SELECT athlete_id, SUM(score) AS total_score
        FROM results
        GROUP BY athlete_id
        ORDER BY total_score DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rank := 1
	for rows.Next() {
		var ranking Ranking
		err = rows.Scan(&ranking.AthleteID, &ranking.TotalScore)
		if err != nil {
			return nil, err
		}
		ranking.Rank = rank
		rank++
		overallRanking = append(overallRanking, ranking)
	}

	return overallRanking, nil
}

func calculateCategoryRanking(category string) ([]Ranking, error) {
	var categoryRanking []Ranking

	db, err := sql.Open("sqlite3", "./ranking.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT athlete_id, SUM(score) AS total_score
		FROM results
		WHERE category = ?
		GROUP BY athlete_id
		ORDER BY total_score DESC
	`, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rank := 1
	for rows.Next() {
		var ranking Ranking
		err = rows.Scan(&ranking.AthleteID, &ranking.TotalScore)
		if err != nil {
			return nil, err
		}
		ranking.Category = category
		ranking.Rank = rank
		rank++
		categoryRanking = append(categoryRanking, ranking)
	}

	return categoryRanking, nil
}
