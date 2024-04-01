package server

import (
	"github.com/Danila331/YAP-2/internal/controllers/forms"
	"github.com/Danila331/YAP-2/internal/controllers/pages"
	"github.com/labstack/echo/v4"
)

func StartServer() {
	app := echo.New()

	app.GET("/", pages.PageMain)
	app.GET("/sign_up", pages.PageSignup)
	app.GET("/login", pages.PageLogin)
	app.POST("/submit_login", forms.LoginFormListener)
	app.POST("/submit_signup", forms.SignFormListener)
	app.Start(":8085")
}
