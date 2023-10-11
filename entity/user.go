package entity

type User struct {
	Id int
	Name string
	Email string
	Birth string
	Password string
	Admin bool
	VoucherId int
}