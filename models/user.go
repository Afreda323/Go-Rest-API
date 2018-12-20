package models

import (
	"os"
	"strings"
	"todo/utils"

	"golang.org/x/crypto/bcrypt"

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
	Token    string `json:"token";sql:"-"`
}

func genToken(u *User) string {
	tk := &Token{UserID: u.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	signed, _ := token.SignedString([]byte(os.Getenv("token_password")))
	return signed
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

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return utils.Message(false, "Something went wrong")
	}

	u.Password = string(hashed) // replace plain pw with hashed one
	GetDB().Create(u)           // save to db
	u.Password = ""             // dont send password back

	u.Token = genToken(u)

	resp := utils.Message(true, "User created")
	resp["user"] = u

	return resp
}

// Login user, get jwt
func Login(email string, password string) map[string]interface{} {
	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "User does not exist")
		}
		return utils.Message(false, "Something went wrong")
	}

	if user.ID <= 0 {
		return utils.Message(false, "Something went wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Incorrect Password")
	}

	user.Token = genToken(user)
	user.Password = ""
	resp := utils.Message(true, "Logged in")
	resp["user"] = user

	return resp
}

// GetUser by ID
func GetUser(id uint) *User {
	user := &User{}
	GetDB().Table("users").Where("id = ?", id).First(user)

	if user.Email == "" {
		return nil
	}

	user.Password = ""
	return user
}
