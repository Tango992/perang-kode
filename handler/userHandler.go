package handler

import (
	"database/sql"
	"fmt"
	"github.com/Tango992/perang-kode/entity"
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
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("| ID | Nama                 | Email                | Birth      | Admin | Voucher    | Discount |")
	fmt.Println("-------------------------------------------------------------------------------------------------")

	for rows.Next() {
		var isAdmin int
		var user entity.User

		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birth, &isAdmin, &user.VoucherName, &user.VoucherNominee); err != nil {
			return err
		}

		if isAdmin == 1 {
			user.Admin = true
		} else {
			user.Admin = false
		}

		if user.VoucherNominee == 0 {
			fmt.Printf("| %-2v | %-20s | %-20s | %-10s | %-5v | %-10s | %-8s |\n", user.Id, user.Name, user.Email, user.Birth, user.Admin, "-", "-")
			continue
		}
		fmt.Printf("| %-2v | %-20s | %-20s | %-10s | %-5v | %-10s | %-7.2f%% |\n", user.Id, user.Name, user.Email, user.Birth, user.Admin, user.VoucherName, user.VoucherNominee * 100)
	}
	fmt.Println("-------------------------------------------------------------------------------------------------")
	return nil
}