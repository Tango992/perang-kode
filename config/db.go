package config

import (
	"database/sql"
	"errors"
	my "github.com/go-mysql/errors"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(connection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connection)
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
