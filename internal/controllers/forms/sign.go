package forms

import (
	"github.com/Danila331/YAP-2/internal/models"
	"github.com/labstack/echo/v4"
)

// Обработчик формы регистрации
func SignFormListener(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")

	var user models.User

	user.Login = login
	user.Password = password

	err := user.Create()

	if err != nil {
		return err
	}

	return nil
}
