package entity

type User struct {
	Id int
	Name string
	Email string
	Birth string
	Age int
	Password []byte
	Admin bool
	Discount
}