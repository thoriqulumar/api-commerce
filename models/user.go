package models

import (
	"thor/api/e-commerce/database"
	"thor/api/e-commerce/vendor/golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DetailUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func (User) TableName() string{
	return "user"
}

func PostRegister(user *User) (map[string]interface{}, error) {
	var dbConn database.Connection
	conn, err := dbConn.OpenConn()
	var response map[string]interface{}

	if err != nil {
		return response, err
	}
	out := uuid.Must(uuid.NewRandom()).String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	dataUser := map[string]interface{}{
		"id": "usr~"+out,
		"username": user.Username,
		"email":    user.Email,
		"password": hashedPassword,
	}
}

func LoginEmail(user *User) (map[string]interface{}, error) {

}

func LoginUsername(user *User) (map[string]interface{}, error) {

}

func UpdateDetailUser(detailUser *DetailUser) (map[string]interface{}, error) {

}
