package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username         string    `json:"username" gorm:"unique;not null"`
	Email            string    `json:"email" gorm:"unique;not null"`
	Password         string    `json:"-" gorm:"not null"`
	PasswordUpdateAt time.Time `json:"password_updatedAt"`
	Task             []Task    `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) UpdatePassword(old, new string) error {
	if err := u.CheckPassword(old); err != nil {
		return errors.New("old password is incorrect")
	}
	if err := u.HashPassword(new); err != nil {
		return errors.New("failed to hash password")
	}
	u.PasswordUpdateAt = time.Now()
	return nil
}
