package handler

import (
	"database/sql"
	"fmt"
	"github.com/Tango992/perang-kode/entity"
	my "github.com/go-mysql/errors"
)

func DisplayStock(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, stock FROM games")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("-------------------------------------------")
	fmt.Println("| Game ID | Nama                 | Stock  |")
	fmt.Println("-------------------------------------------")
	for rows.Next() {
		var row entity.Stock

		if err := rows.Scan(&row.Id, &row.Name, &row.Stock); err != nil {
			return err
		}

		fmt.Printf("| %-7d | %-20s | %-6d |\n", row.Id, row.Name, row.Stock)
	}
	fmt.Println("-------------------------------------------")
	return nil
}

func UpdateStock(db *sql.DB, data entity.Stock) error {
	var exists bool
	row := db.QueryRow("SELECT EXISTS(SELECT * FROM games WHERE id = ?)", data.Id)
	if err := row.Scan(&exists); err != nil {
		return err
	} else if !exists {
		return fmt.Errorf("game id di luar jangkauan")
	}

	_, err := db.Exec("UPDATE games SET stock = ? WHERE id = ?", data.Stock, data.Id)
	if err != nil {
		if my.MySQLErrorCode(err) == 1264 {
			return fmt.Errorf("stock tidak boleh dibawah 0")
		}
		return err
	}

	ClearTerminal()
	DisplayStock(db)
	fmt.Printf("\nStock berhasil ter-update\n")
	return nil
}