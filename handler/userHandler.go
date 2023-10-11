package handler

import (
	"database/sql"
	"fmt"
	"perang-kode/entity"
)

func UserReport(db *sql.DB) error {
	rows, err := db.Query(`
		SELECT u.id, u.name, u.email, u.birth, u.admin, IFNULL(d.voucher, 0), IFNULL(d.nominee, 0)
		FROM users u
		LEFT JOIN discounts d ON d.id = u.discount_id`)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Printf("\nUSER REPORT\n")
	for rows.Next() {
		var adminbuffer int
		var user entity.User

		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birth, &adminbuffer, &user.VoucherName, &user.VoucherNominee); err != nil {
			return err
		}

		if adminbuffer == 1 {
			user.Admin = true
		} else {
			user.Admin = false
		}

		fmt.Printf("\nUser ID\t\t: %v\n", user.Id)
		fmt.Printf("Name\t\t: %v\n", user.Name)
		fmt.Printf("Email\t\t: %v\n", user.Email)
		fmt.Printf("Birthday\t: %v\n", user.Birth)
		fmt.Printf("Admin privilege\t: %v\n", user.Admin)
		if !user.Admin {
			if user.VoucherNominee == 0 {
				fmt.Printf("Voucher\t\t: -\n")
				return nil
			}
			fmt.Printf("Voucher\t\t: '%v' value %v%%\n", user.VoucherName, user.VoucherNominee * 100)
		}
	}
	return nil
}