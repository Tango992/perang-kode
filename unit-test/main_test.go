package perangkode

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Tango992/perang-kode/entity"
	"github.com/Tango992/perang-kode/handler"

	"github.com/stretchr/testify/assert"
)

var db, dbfalse *sql.DB

var user = entity.User{
	Id:       1,
	Name:     "Dummy",
	Email:    "dummy@mail.com",
	Birth:    "2005-01-01",
	Age:      18,
	Password: []byte("dummy"),
	Admin:    true,
}

var userFalse = entity.User{
	Id:       0,
	Name:     "Dummyfalse",
	Email:    "dummyfalse@mail.com",
	Birth:    "2005-01-01",
	Age:      18,
	Password: []byte("dummyfalse"),
	Admin:    true,
}

func TestMain(m *testing.M) {
	var err error
	// Establish Db Connection
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/perang_kode_test?multiStatements=true")
	if err != nil {
		log.Fatal(err)
	} else if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Establish False Db Connections
	dbfalse, _ = sql.Open("mysql", "root:@tcp(localhost:3306)/nonexistentialdb")

	_, err = db.Exec(`
		DELETE FROM perang_kode_test.users_games;
		ALTER TABLE users_games AUTO_INCREMENT = 1;
		DELETE FROM perang_kode_test.users;
		ALTER TABLE users AUTO_INCREMENT = 1;
	`)
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestRegister(t *testing.T) {
	assert.Nil(t, handler.Register(user, db))                 // Register
	assert.NotNil(t, handler.Register(user, db).Error())      // Register with the same account
	assert.NotNil(t, handler.Register(user, dbfalse).Error()) // Register with a false database
}

func TestLogin(t *testing.T) {
	_, _, success := handler.Login(user.Email, user.Password, db)
	assert.Nil(t, success) // Login with registered account

	_, _, failed := handler.Login(userFalse.Email, userFalse.Password, db)
	assert.NotNil(t, failed.Error()) // Login with false credentials
}

func TestShowAllgames(t *testing.T) {
	os.Stdout = nil
	assert.Nil(t, handler.ShowAllGames(db))                 // Show all games
	assert.NotNil(t, handler.ShowAllGames(dbfalse).Error()) // Show all games with false database
}

func TestGetVoucher(t *testing.T) {
	assert.Nil(t, handler.GetVoucher(&user, db))                 // Get voucher
	assert.NotNil(t, handler.GetVoucher(&user, dbfalse).Error()) // Get voucher with false database
}

func TestShowCart(t *testing.T) {
	assert.Nil(t, handler.ShowCart(user, db))                 // Show cart
	assert.NotNil(t, handler.ShowCart(user, dbfalse).Error()) // Show cart with false database
}

func TestAddGameToCart(t *testing.T) {
	assert.Nil(t, handler.AddGameToCart(user, 1, db))                 //Adding game to cart
	assert.NotNil(t, handler.AddGameToCart(user, 1, db))              //Adding duplicate game to cart
	assert.NotNil(t, handler.AddGameToCart(user, 2, db))              // Adding game exceeding age limit
	assert.NotNil(t, handler.AddGameToCart(user, -100, db).Error())   // Adding game to cart with non existent game
	assert.NotNil(t, handler.AddGameToCart(user, 1, dbfalse).Error()) // Adding game to cart with false database
	assert.NotNil(t, handler.AddGameToCart(userFalse, 1, db).Error()) // Adding game to cart with non existent user
}

func TestRemoveGameFromCart(t *testing.T) {
	assert.Nil(t, handler.RemoveGameFromCart(user, 1, db))                 // Remove game from cart
	assert.NotNil(t, handler.RemoveGameFromCart(user, 1, db).Error())      // Remove game from cart that has already been removed
	assert.NotNil(t, handler.RemoveGameFromCart(user, -100, db).Error())   // Remove game from cart with non existent game id
	assert.NotNil(t, handler.RemoveGameFromCart(userFalse, 1, db).Error()) // Remove game from cart from non existent user
	assert.NotNil(t, handler.RemoveGameFromCart(user, 1, dbfalse).Error()) // Remove game from cart with false database
}

func TestUserReport(t *testing.T) {
	os.Stdout = nil
	assert.Nil(t, handler.UserReport(db))                 // Display user report
	assert.NotNil(t, handler.UserReport(dbfalse).Error()) // Display user report from false database
}

func TestUpdateStock(t *testing.T) {
	dataTrue := entity.Stock{
		Id:    1,
		Stock: 100,
	}
	assert.Nil(t, handler.UpdateStock(db, dataTrue)) // Updating stock

	dataFalse1 := entity.Stock{
		Id:    1,
		Stock: -100,
	}
	assert.NotNil(t, handler.UpdateStock(db, dataFalse1).Error()) // Updating stock with invalid quantity

	dataFalse2 := entity.Stock{
		Id:    -1,
		Stock: 100,
	}
	assert.NotNil(t, handler.UpdateStock(db, dataFalse2).Error()) // Updating stock with invalid game id
}

func TestStockReport(t *testing.T) {
	assert.Nil(t, handler.DisplayStock(db))                 // Display stock
	assert.NotNil(t, handler.DisplayStock(dbfalse).Error()) // Display stock with false database
}

func TestOrderReport(t *testing.T) {
	assert.Nil(t, handler.OrderReport(db))                 // Display order report
	assert.NotNil(t, handler.OrderReport(dbfalse).Error()) // Display order report from non existent database
}
