package routes

import (
	"thor/api/e-commerce/config"
	"thor/api/e-commerce/controller"
	"thor/api/e-commerce/helper"
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

func Routers(){
	e := echo.New()
	e.Validator = &helper.CustomValidator{Validator: validator.New()}
	e.POST("/user/register", controller.PostRegisterController)

	e.Logger.Fatal(e.Start(":" + config.ServerPort()))

	
}