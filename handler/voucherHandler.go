package handler

import (
	"database/sql"
	"fmt"
	"math/rand"
	"perang-kode/entity"
)

func GetVoucher(user *entity.User, db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	} else if user.DiscountId != 0 {
		fmt.Printf("\nAnda sudah memiliki voucher '%v' dengan nominal %.2f%%\n", user.VoucherName, user.VoucherNominee * 100)
		return nil
	}

	var voucherLength int
	row := db.QueryRow("SELECT id FROM discounts ORDER BY id DESC LIMIT 1")
	if err := row.Scan(&voucherLength); err != nil {
		return err
	}

	gacha := rand.Intn(voucherLength) + 1
	if _, err := db.Exec("UPDATE users SET discount_id = ? WHERE id = ?", gacha, user.Id); err != nil {
		return err
	}
	row1 := db.QueryRow("SELECT id, voucher, nominee FROM discounts WHERE id = ?", gacha)
	if err := row1.Scan(&user.DiscountId, &user.VoucherName, &user.VoucherNominee); err != nil {
		return err
	}
	fmt.Printf("\nSelamat, anda mendapatkan voucher '%v' dengan nominal %.2f%%!\n", user.VoucherName, user.VoucherNominee * 100)
	return nil
}