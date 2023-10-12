
package main

import (
	"database/sql"
	"log"
	"os"
	"perang-kode/entity"
	"perang-kode/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db, dbfalse *sql.DB

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
	user := entity.User{
		Name: "Dummy",
		Email: "dummy@mail.com",
		Birth: "2001-01-01",
		Password: []byte("dummy"),
		Admin: true,
	}
	assert.Nil(t, handler.Register(user, db))
	assert.NotNil(t, handler.Register(user, dbfalse).Error())

	assert.NotNil(t, handler.Register(user, dbfalse).Error())
}


func TestLogin(t *testing.T) {
	user := entity.User{
		Email: "dummy@mail.com",
		Password: []byte("dummy"),
	}
	userFalse := entity.User{
		Email: "dummyfalse@mail.com",
		Password: []byte("dummyfalse"),
	}

	_, _, success := handler.Login(user.Email, user.Password, db)
	assert.Nil(t, success)

	_, _, failed := handler.Login(userFalse.Email, userFalse.Password, db)
	assert.NotNil(t, failed.Error())
}

func TestUserReport(t *testing.T) {
	os.Stdout = nil
	assert.Nil(t, handler.UserReport(db))
	assert.NotNil(t, handler.UserReport(dbfalse).Error())
}

func TestUpdateStock(t *testing.T) {
	dataTrue := entity.Stock{
		Id: 1,
		Stock: 100,
	}
	assert.Nil(t, handler.UpdateStock(db, dataTrue))

	dataFalse1 := entity.Stock{
		Id: 1,
		Stock: -100,
	}
	assert.NotNil(t, handler.UpdateStock(db, dataFalse1).Error())

	dataFalse2 := entity.Stock{
		Id: -1,
		Stock: -100,
	}
	assert.NotNil(t, handler.UpdateStock(db, dataFalse2).Error())
}

func TestStockReport(t *testing.T) {
	assert.Nil(t, handler.DisplayStock(db))
	assert.NotNil(t, handler.DisplayStock(dbfalse).Error())
}

func TestOrderReport(t *testing.T) {
	assert.Nil(t, handler.OrderReport(db))
	assert.NotNil(t, handler.OrderReport(dbfalse).Error())
}
