package models

import (
	"strings"
	"todo/utils"

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
// make sure user does not exist
func (u *User) Validate() (map[string]interface{}, bool) {
	if !strings.Contains(u.Email, "@") {
		return utils.Message(false, "Invalid Email"), false
	}
	if len(u.Password) < 6 {
		return utils.Message(false, "Invalid Password"), false
	}

	temp := &User{}
	err := GetDB().Table("users").Where("email = ?", u.Email).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return utils.Message(false, "Something went wrong"), false
	}

	if temp.Email != "" {
		return utils.Message(false, "Email already taken"), false
	}
	return utils.Message(true, "Validated"), true
}

// CreateUser adds a new user to the db
func (u *User) CreateUser() map[string]interface{} {
	if message, ok := u.Validate(); !ok {
		return message
	}
}

// Login user, get jwt
func Login(email string, password string) {

}

// GetUser by ID
func GetUser(id uint) {

}
