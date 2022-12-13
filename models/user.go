package models

import (
	"thor/api/e-commerce/database"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"  validate:"required"`
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password"  validate:"required"`
}

type DetailUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func (User) TableName() string {
	return "user"
}

func (DetailUser) TableName() string {
	return "detail_user"
}

func PostRegister(user *User) (map[string]interface{}, error) {
	var dbConn database.Connection
	var response map[string]interface{}

	conn, err := dbConn.OpenConn()

	if err != nil {
		return response, err
	}
	generateID := uuid.Must(uuid.NewRandom()).String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	dataUser := map[string]interface{}{
		"id":       "usr~" + generateID,
		"username": user.Username,
		"email":    user.Email,
		"password": hashedPassword,
	}

	result := conn.Model(User{}).Create(dataUser)

	if result.RowsAffected > 0 {
		generateDetailID := uuid.Must(uuid.NewRandom()).String()
		dataDetail := map[string]interface{}{
			"id":       "detUsr~" + generateDetailID,
			"user_id":  "usr~" + generateID,
			"username": user.Username,
			"email":    user.Email,
		}

		result := conn.Model(DetailUser{}).Create(dataDetail)
		if result.RowsAffected > 0 {
			response = map[string]interface{}{
				"message": "user created successfully",
				"status": 1,
			}
			return response, nil
		}
	}

	return response, nil
}

// func LoginEmail(user *User) (map[string]interface{}, error) {

// }

// func LoginUsername(user *User) (map[string]interface{}, error) {

// }

// func UpdateDetailUser(detailUser *DetailUser) (map[string]interface{}, error) {

// }
