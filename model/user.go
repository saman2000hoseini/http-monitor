package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null"`
	Password string `gorm:"not null"`
	URLs     []URL  `gorm:"null"`
}

func (u *User) HashPassword(pass string) error {
	if len(pass) < 8 {
		return errors.New("password must have at least 8 characters")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err == nil {
		u.Password = string(h)
	}
	return err
}

func (u *User) CheckPassword(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	return err == nil
}
