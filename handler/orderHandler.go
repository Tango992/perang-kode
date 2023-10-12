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
	rows, err := db.Query(`
		SELECT u.id, u.name, IFNULL(d.nominee, 0), g.price
		FROM users_games ug
		JOIN users u ON u.id = ug.user_id
		JOIN games g ON g.id = ug.game_id
		LEFT JOIN discounts d ON u.discount_id = d.id
		WHERE g.name = ?`, game)
	if err != nil {
		return err
	}
	defer rows.Close()

	var grossIncome float32
	counter := 1
	fmt.Printf("\n%v\n", strings.ToUpper(game))
	fmt.Println("--------------------------------------------------")
	fmt.Println("| No  | User ID |  Nama                 | Harga  |")
	fmt.Println("--------------------------------------------------")
	for rows.Next() {
		var user entity.User
		var price float32

		if err := rows.Scan(&user.Id, &user.Name, &user.VoucherNominee, &price); err != nil {
			return err
		}

		priceAfterDiscount := price * (1 - user.VoucherNominee)
		grossIncome += priceAfterDiscount

		if priceAfterDiscount == 0 {
			fmt.Printf("| %-3v | %-7v | %-21v | %-6s |\n", counter, user.Id, user.Name, "Free")
			continue
		}
		fmt.Printf("| %-3v | %-7v | %-21v | %-6.2f |\n", counter, user.Id, user.Name, priceAfterDiscount)
		counter++
	}
	fmt.Println("--------------------------------------------------")
	fmt.Printf("|                       Gross Income    | %-6.2f |\n", grossIncome)
	fmt.Println("--------------------------------------------------")
	return nil
}