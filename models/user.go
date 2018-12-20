package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Token struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// User struct
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";gorm:"-"`
}

// Validate if all values are present
func (u *User) Validate() {

}

// CreateUser adds a new user to the db
func (u *User) CreateUser() {

}

// Login user, get jwt
func Login(email string, password string) {

}

// GetUser by ID
func GetUser(id uint) {

}
