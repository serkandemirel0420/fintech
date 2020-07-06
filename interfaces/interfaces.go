package interfaces

import "github.com/jinzhu/gorm"

//User usr
type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

//Account acnt
type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

//ResponseAccount ResponseAccount
type ResponseAccount struct {
	ID      uint
	Name    string
	Balance int
}

//ResponseUser ResponseUser
type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}

//Validation  interface
type Validation struct {
	Value string
	Valid string
}
