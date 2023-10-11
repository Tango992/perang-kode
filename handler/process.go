package handler

import (
	"database/sql"
	"fmt"
	"perang-kode/entity"
)

// ShowAllGames menampilkan semua daftar game
func ShowAllGames(db *sql.DB) {
	rows, err := db.Query("SELECT g.id, g.name, g.description, m.name, g.price, g.stock FROM games g JOIN maturity m ON m.id = g.maturity_id ORDER BY g.id ASC")
	if err != nil {
		fmt.Println("Error querying games:", err)
		return
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
		fmt.Printf("Price\t\t: %v\n", game.Price)
		fmt.Printf("Stock\t\t: %v\n", game.Stock)
	}
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
		if _, err := db.Exec("INSERT INTO users_games (user_id, game_id) VALUES (?, ?)", userID, gameID); err != nil {
			return err
		}
		
		if _ , err := db.Exec("UPDATE games SET stock = stock - 1 WHERE id = ?", gameID); err != nil {
			return err
		}
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
