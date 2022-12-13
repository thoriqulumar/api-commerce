package routes

import (
	"thor/api/e-commerce/config"
	"github.com/labstack/echo/v4"
)

func Routers(){
	e := echo.New()

	e.Logger.Fatal(e.Start(":" + config.ServerPort()))
}