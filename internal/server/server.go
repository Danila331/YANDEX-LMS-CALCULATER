package server

import (
	"log"
	"os"

	"github.com/Danila331/YAP-2/internal/controllers/forms"
	"github.com/Danila331/YAP-2/internal/controllers/pages"
	"github.com/labstack/echo/v4"
)

func StartServer() {
	app := echo.New()
	log.SetOutput(os.Stdout)

	// Использование стандартного логгера Go для записи в терминал
	app.Logger.SetOutput(os.Stdout)
	app.GET("/", pages.PageMain)
	app.GET("/sign_up", pages.PageSignup)
	app.GET("/login", pages.PageLogin)
	app.GET("/time", pages.PageTime)

	app.POST("/login_login", forms.LoginFormListener)
	app.POST("/login_signup", forms.SignFormListener)

	app.Start(":8085")
}
