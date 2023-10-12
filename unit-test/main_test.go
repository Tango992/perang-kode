package main

import (
	"database/sql"
	"os"
	"perang-kode/config"
	"perang-kode/entity"
	"perang-kode/handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

var db, _ = config.ConnectDB()

func falseDB() (*sql.DB) {
	db, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/nonexistentialdb")
	return db
}

func TestUserReport(t *testing.T) {
	os.Stdout = nil
	assert.Nil(t, handler.UserReport(db))
	assert.NotNil(t, handler.UserReport(falseDB()))
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
	assert.NotNil(t, handler.UpdateStock(db, dataFalse1))

	dataFalse2 := entity.Stock{
		Id: -1,
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

