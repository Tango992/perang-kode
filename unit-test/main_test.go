package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"perang-kode/config"
	"perang-kode/entity"
	"perang-kode/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db, _ = config.ConnectDB()

func falseDB() *sql.DB {
	db, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/nonexistentialdb")
	return db
}

func TestShowAllGames(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	defer func() {
		os.Stdout = old
	}()

	handler.ShowAllGames(db)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	fmt.Println("Output from ShowAllGames:", string(out))
}

// func TestShowCart(t *testing.T) {
// 	os.Stdout = nil
// 	assert.Nil(t, handler.ShowCart(db))
// 	assert.NotNil(t, handler.ShowCart(falseDB()))
// }

// func TestAddGameToCart(t *testing.T) {
// 	os.Stdout = nil
// 	assert.Nil(t, handler.AddGameToCart(db))
// 	assert.NotNil(t, handler.AddGameToCart(falseDB()))
// }
// func TestIsStockAvailable(t *testing.T) {
// 	os.Stdout = nil
// 	assert.Nil(t, handler.IsStockAvailable(db))
// 	assert.NotNil(t, handler.IsStockAvailable(falseDB()))
// }
// func TestRemoveGameFromCart(t *testing.T) {
// 	os.Stdout = nil
// 	assert.Nil(t, handler.RemoveGameFromCart(db))
// 	assert.NotNil(t, handler.RemoveGameFromCart(falseDB()))
// }

// func TestGetVoucher(t *testing.T) {
// 	dataTrue := entity.Discount{
// 		DiscountId:     1,
// 		VoucherName:    "GAMERS",
// 		VoucherNominee: 0.1,
// 	}

// 	assert.Nil(t, handler.GetVoucher(db, dataTrue))

// 	dataFalse := entity.Discount{
// 		DiscountId:     -1,
// 		VoucherName:    "GAMERS",
// 		VoucherNominee: 0.1,
// 	}
// 	assert.NotNil(t, handler.GetVoucher(db, dataFalse))
// }

func TestUserReport(t *testing.T) {
	os.Stdout = nil
	assert.Nil(t, handler.UserReport(db))
	assert.Nil(t, handler.UserReport(falseDB()))
}

func TestUpdateStock(t *testing.T) {
	dataTrue := entity.Stock{
		Id:    1,
		Stock: 100,
	}
	assert.Nil(t, handler.UpdateStock(db, dataTrue))

	dataFalse1 := entity.Stock{
		Id:    1,
		Stock: -100,
	}
	assert.NotNil(t, handler.UpdateStock(db, dataFalse1))

	dataFalse2 := entity.Stock{
		Id:    -1,
		Stock: 100,
	}
	assert.NotNil(t, handler.UpdateStock(db, dataFalse2))
}

func TestStockReport(t *testing.T) {
	assert.Nil(t, handler.DisplayStock(db))
	assert.NotNil(t, handler.DisplayStock(falseDB()))
}

func TestOrderReport(t *testing.T) {
	assert.Nil(t, handler.OrderReport(db))
	assert.NotNil(t, handler.OrderReport(falseDB()))
}
