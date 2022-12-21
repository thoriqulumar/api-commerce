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

type UserLogin struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"  validate:"required"`
}

type DetailUser struct {
	IdUser   string `json:"user_id"`
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

func CheckExistingUsernameAndEmail(user User) (map[string]interface{}, bool) {
	var dbConn database.Connection
	existingUser := []User{}
	conn, err := dbConn.OpenConn()

	if err != nil {
		return nil, false
	}

	conn.Select("user.*").
		Where("username", user.Username).
		Find(&existingUser)

	if len(existingUser) > 0 {
		response := map[string]interface{}{
			"message": "username already taken",
			"status":  0,
		}
		return response, false
	}

	conn.Select("user.*").
		Where("email", user.Email).
		Find(&existingUser)

	if len(existingUser) > 0 {
		response := map[string]interface{}{
			"message": "email already taken",
			"status":  0,
		}
		return response, false
	}
	return nil, true
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

	response, isDuplicate := CheckExistingUsernameAndEmail(*user)

	if !isDuplicate {
		return response, nil
	}

	result := conn.Model(User{}).Create(dataUser)

	if result.RowsAffected > 0 {
		generateDetailID := uuid.Must(uuid.NewRandom()).String()
		dataDetail := map[string]interface{}{
			"id":       "detUsr~" + generateDetailID,
			"user_id":  dataUser["id"],
			"username": user.Username,
			"email":    user.Email,
		}

		result := conn.Model(DetailUser{}).Create(dataDetail)
		if result.RowsAffected > 0 {
			response = map[string]interface{}{
				"message": "user created successfully",
				"id":      dataUser["id"],
				"status":  1,
			}
			return response, nil
		}
	}

	return response, nil
}

func PostLogin(user *UserLogin, method string) (map[string]interface{}, error) {
	var dbConn database.Connection
	var response map[string]interface{}
	existingUser := User{}

	conn, err := dbConn.OpenConn()

	if err != nil {
		return response, err
	}
	if method == "email" {
		conn.Select("user.*").
			Where("email", user.Email).
			Find(&existingUser)
	} else {
		conn.Select("user.*").
			Where("username", user.Username).
			Find(&existingUser)
	}

	if existingUser.Password != "" {
		if CheckPasswordHash(user.Password, existingUser.Password) {
			response = map[string]interface{}{
				"message": "login successfully",
				"id":      existingUser.Id,
				"status":  1,
			}
			return response, nil
		} else {
			response = map[string]interface{}{
				"message": "password incorrect",
				"status":  0,
			}
			return response, nil
		}
	}

	response = map[string]interface{}{
		"message": "email or username incorrect",
		"status":  0,
	}
	return response, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UpdateDetailUser(detailUser *DetailUser) (map[string]interface{}, error) {
	var dbConn database.Connection
	var response map[string]interface{}

	conn, err := dbConn.OpenConn()

	if err != nil {
		return response, err
	}

	data := map[string]interface{}{
		"username":  detailUser.Username,
		"email":     detailUser.Email,
		"full_name": detailUser.FullName,
		"phone":     detailUser.Phone,
		"address":   detailUser.Address,
	}

	result := conn.Model(&detailUser).
		Select("detail_user").
		Where("user_id", detailUser.IdUser).
		Updates(data)

	if result.RowsAffected > 0 {
		response = map[string]interface{}{
			"message": "user info updated",
			"status":  1,
		}

		return response, nil
	}

	response = map[string]interface{}{
		"message": "user info failed to update",
		"status":  0,
	}

	return response, nil
}


func GetDetailUser(detailUser *DetailUser) (map[string]interface{}, error) {
	var dbConn database.Connection
	var response map[string]interface{}
	var user DetailUser
	
	conn, err := dbConn.OpenConn()

	if err != nil {
		return response, err
	}


	result := conn.Model(&detailUser).
		Select("detail_user.*").
		Where("user_id", detailUser.IdUser).
		Find(&user)

	if result.RowsAffected > 0 {
		response = map[string]interface{}{
			"user": user,
			"status":  1,
		}

		return response, nil
	}

	response = map[string]interface{}{
		"message": "user info not found",
		"status":  0,
	}

	return response, nil
}
