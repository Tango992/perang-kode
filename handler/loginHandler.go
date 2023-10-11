package handler

import (
	"database/sql"
	"perang-kode/entity"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password []byte, db *sql.DB) (entity.User, bool, error) {
	var user entity.User
	row := db.QueryRow("SELECT id, name, email, birth, password, admin, IFNULL(discount_id, 0) FROM users WHERE email = ?", email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Birth, &user.Password, &user.Admin, &user.DiscountId)
	if err != nil {
		return user, false, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return user, false, err
	}
	user.Password = []byte{}
	return user, true, nil
}
