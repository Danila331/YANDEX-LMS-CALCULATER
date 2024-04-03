package forms

import (
	"net/http"

	"github.com/Danila331/YAP-2/internal/models"
	"github.com/Danila331/YAP-2/internal/pkg"
	"github.com/labstack/echo/v4"
)

func LoginFormListener(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")

	var user models.User
	user.Login = login
	user, err := user.ReadByLogin()
	if err != nil {
		return err
	}
	if password != user.Password {
		return echo.NewHTTPError(http.StatusNotFound, "Неверный пароль")
	}
	tokenString, err := pkg.GenerateJWT(login)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
