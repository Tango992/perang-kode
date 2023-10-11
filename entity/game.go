// entity/game.go

package entity

// Game adalah struktur yang merepresentasikan data game dalam database
type Game struct {
	ID          int
	Name        string
	Description string
	Maturity_id int
	Price       float64
	Stock       int
}

// UserGame adalah struktur yang merepresentasikan game dalam cart pengguna
type UserGame struct {
	ID       int
	UserID   int
	GameID   int
	Quantity int
}
