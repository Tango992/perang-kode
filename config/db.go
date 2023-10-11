package config

import (
	"database/sql"
	"errors"
	my "github.com/go-mysql/errors"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/perang_kode")
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		if my.MySQLErrorCode(err) == 1049 {
			return db, errors.New("database tidak ditemukan")
		}
		return db, err
	}
	return db, nil
}
