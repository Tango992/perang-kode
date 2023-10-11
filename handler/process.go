package handler

import (
	"database/sql"
	"fmt"
	"perang-kode/entity"
)

// ShowAllGames menampilkan semua daftar game
func ShowAllGames(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM games")
	if err != nil {
		fmt.Println("Error querying games:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Daftar Game:")
	fmt.Println("-------------------------------------------------")
	fmt.Println("| ID  | Nama             	      | Deskripsi  | Maturity ID | Harga  | Stok |")
	fmt.Println("-------------------------------------------------")

	for rows.Next() {
		var game entity.Game
		if err := rows.Scan(&game.ID, &game.Name, &game.Description, &game.Maturity_id, &game.Price, &game.Stock); err != nil {
			fmt.Println("Error scanning game:", err)
			continue
		}
		fmt.Printf("| %-3d | %-29s | %-94s | %-11d | %-6.2f | %-4d |\n", game.ID, game.Name, game.Description, game.Maturity_id, game.Price, game.Stock)
	}

	fmt.Println("-------------------------------------------------")
}

// ShowCart menampilkan isi cart pengguna
func ShowCart(userID int, db *sql.DB) {
	rows, err := db.Query("SELECT g.id, g.Name, g.Price FROM users_games u JOIN games g ON u.game_id = g.id WHERE u.user_id = ?", userID)
	if err != nil {
		fmt.Println("Error querying cart:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Cart Anda:")
	fmt.Println("-------------------------------------------------")
	fmt.Println("| ID  | Nama             | Harga  |")
	fmt.Println("-------------------------------------------------")

	for rows.Next() {
		var item entity.CartItem
		if err := rows.Scan(&item.GameID, &item.Name, &item.Price); err != nil {
			fmt.Println("Error scanning cart item:", err)
			continue
		}
		fmt.Printf("| %-3d | %-16s | %-6.2f |\n", item.GameID, item.Name, item.Price)
	}

	fmt.Println("-------------------------------------------------")
}

// AddGameToCart menambahkan game ke cart pengguna
func AddGameToCart(userID int, gameID int, db *sql.DB) error {
	available, err := IsStockAvailable(gameID, db)
	if err != nil {
		return err
	}

	if available {
		_, err := db.Exec("INSERT INTO users_games (user_id, game_id) VALUES (?, ?)", userID, gameID)
		if err != nil {
			return err
		}
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
		return false, err
	}
	return stock >= 1, nil
}

// RemoveGameFromCart menghapus game dari cart pengguna.
func RemoveGameFromCart(userID, gameID int, db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users_games WHERE user_id = ? AND game_id = ?", userID, gameID)
	if err != nil {
		return err
	}
	return nil
}
