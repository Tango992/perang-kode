package entity

type User struct {
	Id         int
	Name       string
	Email      string
	Birth      string
	Password   []byte
	Admin      bool
	DiscountId int
}
