package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password string
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0); err == nil {
		tx.Statement.SetColumn("Password", pw)
	}

	return
}

func (u *User) Validate(pwd string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
}
