package controller

import (
	"net/http"
	"net/mail"
	"thor/api/e-commerce/models"

	"github.com/labstack/echo/v4"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func PostRegisterController(ctx echo.Context) error {
	bodyPayload := new(models.User)
	
	if err := ctx.Bind(bodyPayload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "missing data"})
	}

	if err := ctx.Validate(bodyPayload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid data", "error": err.Error()})
	}

	if !valid(bodyPayload.Email) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "email format not correct"})
	}

	response, err := models.PostRegister(bodyPayload)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, response)
}

// func PostLoginController(ctx echo.Context) error {
// 	bodyPayload := new(models.User)

// 	if err := ctx.Bind(bodyPayload); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "missing data"})
// 	}

// 	if err := ctx.Validate(bodyPayload); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid data"})
// 	}

// 	if bodyPayload.Email != "" && valid(bodyPayload.Email) {
// 		response, err := models.LoginEmail(bodyPayload)

// 		if err != nil {
// 			return ctx.JSON(http.StatusBadRequest, err.Error())
// 		}

// 		return ctx.JSON(http.StatusOK, response)
// 	}

// 	response, err := models.LoginUsername(bodyPayload)

// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	return ctx.JSON(http.StatusOK, response)

// }

// func PutDetailUserController(ctx echo.Context) error {
// 	bodyPayload := new(models.DetailUser)

// 	if err := ctx.Bind(bodyPayload); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "missing data"})
// 	}

// 	if err := ctx.Validate(bodyPayload); err != nil {
// 		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "invalid data"})
// 	}

// 	response, err := models.UpdateDetailUser(bodyPayload)

// 	if err != nil {
// 		return ctx.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	return ctx.JSON(http.StatusOK, response)

// }
