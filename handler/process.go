package handler

import (
	"database/sql"
	"fmt"
	"github.com/Tango992/perang-kode/entity"
)

// ShowAllGames menampilkan semua daftar game
func ShowAllGames(db *sql.DB) error {
	rows, err := db.Query("SELECT g.id, g.name, g.description, m.name, g.price, g.stock FROM games g JOIN maturity m ON m.id = g.maturity_id ORDER BY g.id ASC")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Daftar Game:")

	for rows.Next() {
		var game entity.Game
		if err := rows.Scan(&game.ID, &game.Name, &game.Description, &game.Maturity, &game.Price, &game.Stock); err != nil {
			fmt.Println("Error scanning game:", err)
			continue
		}
		fmt.Printf("\nID\t\t: %v\n", game.ID)
		fmt.Printf("Game\t\t: %v\n", game.Name)
		fmt.Printf("Description\t: %v\n", game.Description)
		fmt.Printf("Maturity\t: %v\n", game.Maturity)
		fmt.Printf("Price\t\t: %.2f\n", game.Price)
		fmt.Printf("Stock\t\t: %v\n", game.Stock)
	}
	return nil
}

// ShowCart menampilkan isi cart pengguna
func ShowCart(user entity.User, db *sql.DB) error {
	rows, err := db.Query("SELECT g.id, g.Name, g.Price FROM users_games u JOIN games g ON u.game_id = g.id WHERE u.user_id = ?", user.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Cart Anda:")
	fmt.Println("-------------------------------------------------")
	fmt.Println("| ID  | Nama                 | Harga  |")
	fmt.Println("-------------------------------------------------")

	var subtotal float64
	for rows.Next() {
		var item entity.CartItem
		if err := rows.Scan(&item.GameID, &item.Name, &item.Price); err != nil {
			return err
		}
		subtotal += item.Price
		if item.Price == 0 {
			fmt.Printf("| %-3d | %-20s | %-6s |\n", item.GameID, item.Name, "Free")
		} else {
			fmt.Printf("| %-3d | %-20s | %-6.2f |\n", item.GameID, item.Name, item.Price)
		}
	}
	fmt.Println("-------------------------------------------------")
	fmt.Printf("                    Subtotal   %.2f\n", subtotal)
	fmt.Printf("                     Voucher   %s\n", user.VoucherName)
	fmt.Printf("                    Discount   %.2f%%\n\n", user.VoucherNominee * 100)
	fmt.Printf("                 Grand Total   %.2f\n", subtotal * (1 - float64(user.VoucherNominee)))
	fmt.Println("-------------------------------------------------")
	return nil
}

// AddGameToCart menambahkan game ke cart pengguna
func AddGameToCart(user entity.User, gameID int, db *sql.DB) error {
	var exists bool
	row := db.QueryRow("SELECT EXISTS(SELECT * FROM users_games WHERE user_id = ? AND game_id = ?)", user.Id, gameID)
	if err := row.Scan(&exists); err != nil {
		return err
	} else if exists {
		return fmt.Errorf("game sudah ada di cart anda")
	}

	available, err1 := IsStockAvailable(gameID, db)
	if err1 != nil {
		return err1
	}

	if available {
		if _, err := db.Exec("INSERT INTO users_games (user_id, game_id) VALUES (?, ?)", user.Id, gameID); err != nil {
			return err
		}
		
		if _ , err := db.Exec("UPDATE games SET stock = stock - 1 WHERE id = ?", gameID); err != nil {
			return err
		}

		ClearTerminal()
		ShowCart(user, db)
		fmt.Printf("\nGame berhasil dimasukkan ke dalam cart.\n")
		return nil

	} else {
		return fmt.Errorf("stok game tidak mencukupi")
	}
}

// IsStockAvailable melakukan pengecekan ketersediaan stok game
func IsStockAvailable(gameID int, db *sql.DB) (bool, error) {
	var stock int
	err := db.QueryRow("SELECT stock FROM games WHERE id = ?", gameID).Scan(&stock)
	if err != nil {
		return false, fmt.Errorf("game id di luar jangkauan")
	}
	return stock >= 1, nil
}

// RemoveGameFromCart menghapus game dari cart pengguna.
func RemoveGameFromCart(user entity.User, gameID int, db *sql.DB) error {
	var exists bool
	row := db.QueryRow("SELECT EXISTS(SELECT * FROM users_games WHERE user_id = ? AND game_id = ?)", user.Id, gameID)
	if err := row.Scan(&exists); err != nil {
		return err
	} else if !exists {
		return fmt.Errorf("game tidak terdapat di dalam cart anda")
	}

	_, err := db.Exec("DELETE FROM users_games WHERE user_id = ? AND game_id = ?", user.Id, gameID)
	if err != nil {
		return err
	}
	ClearTerminal()
	ShowCart(user, db)
	fmt.Printf("\nGame berhasil dihapus dari cart\n")
	return nil
}
