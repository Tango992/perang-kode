package handler

import (
	"database/sql"
	"perang-kode/entity"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password []byte, db *sql.DB) (entity.User, bool, error) {
	var user entity.User
	row := db.QueryRow(`
		SELECT u.id, u.name, u.email, u.birth, u.password, u.admin, IFNULL(discount_id, 0) AS CheckDiscount, IFNULL(d.voucher, 0), IFNULL(d.nominee, 0)
		FROM users u
		LEFT JOIN discounts d ON d.id = u.discount_id
		WHERE email = ?`, email)
  
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Birth, &user.Password, &user.Admin, &user.DiscountId, &user.VoucherName, &user.VoucherNominee)

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
