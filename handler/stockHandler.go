package handler

import (
	"database/sql"
	"fmt"
	"perang-kode/entity"
)

func DisplayStock(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, stock FROM games")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Printf("\nCURRENT STOCK\n")
	for rows.Next() {
		var row entity.Stock

		if err := rows.Scan(&row.Id, &row.Name, &row.Stock); err != nil {
			return err
		}

		fmt.Printf("\nID\t: %v\n", row.Id)
		fmt.Printf("Game\t: %v\n", row.Name)
		fmt.Printf("Stock\t: %v\n", row.Stock)
	}
	return nil
}

func UpdateStock(db *sql.DB, data entity.Stock) error {
	_, err := db.Exec("UPDATE games SET stock = ? WHERE id = ?", data.Stock, data.Id)
	if err != nil {
		return err
	}
	fmt.Printf("\nStock berhasil ter-update\n")
	return nil
}