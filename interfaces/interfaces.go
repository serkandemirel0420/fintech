package interfaces

import "github.com/jinzhu/gorm"

//User user
type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

//Account account
type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

//ResponseAccount resp acc
type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

//ResponseUser resp user
type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}
