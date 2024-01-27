package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID	uint			`json:"id"`
	FirstName 	string	`json:"first_name"`
	LastName 	string	`json:"last_name`
	Email 		string	`json:"email" gorm:"unique"`
	Password 	[]byte	`json:"-"`
	IsAmbassdor bool	`json:"-"`
}

func (u *User) SetPassword(pwd string) {
	// Create hash password
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(pwd), 12)
	u.Password = hashPwd
}

func (u *User) ComparePassword(pwd string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(pwd))
}
