package handler

import (
	"database/sql"
	"errors"
	"perang-kode/entity"

	my "github.com/go-mysql/errors"
	"golang.org/x/crypto/bcrypt"
)

func Register(user entity.User, db *sql.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (name, email, birth, password, admin) VALUES (?,?,?,?,?)", user.Name, user.Email, user.Birth, hashedPassword, user.Admin)
	if my.MySQLErrorCode(err) == 1062 {
		return errors.New("username tidak tersedia. silakan pilih username lain")
	}
	return err
}
