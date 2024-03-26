package server

import (
	"github.com/labstack/echo/v4"
)

func StartServer() {
	app := echo.New()
	app.Start(":8080")
}
