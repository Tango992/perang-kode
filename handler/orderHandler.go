package handler

import (
	"database/sql"
	"fmt"
	"github.com/Tango992/perang-kode/entity"
	"strings"
)

func OrderReport(db *sql.DB) error {
	rows, err := db.Query("SELECT name FROM games")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Printf("\nORDER REPORT\n")

	for rows.Next() {
		var game_name string
		if err :=  rows.Scan(&game_name); err != nil {
			return err
		}

		if err := ReportPerGame(db, game_name); err != nil {
			return err
		}
	}
	return nil
}

func ReportPerGame(db *sql.DB, game string) error {
	rows, err := db.Query(`SELECT u.id, u.name
	FROM users_games ug
	JOIN users u ON u.id = ug.user_id
	JOIN games g ON g.id = ug.game_id
	WHERE g.name = ?`, game)
	if err != nil {
		return err
	}
	defer rows.Close()

	counter := 1
	fmt.Printf("\n%v:\n", strings.ToUpper(game))
	for rows.Next() {
		var user entity.User

		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return err
		}

		fmt.Printf("%v. User ID: %v, Name: %v\n", counter, user.Id, user.Name)
		counter++
	}
	return nil
}