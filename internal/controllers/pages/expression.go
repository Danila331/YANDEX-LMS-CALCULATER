package pages

import (
	"path/filepath"
	"text/template"

	"github.com/Danila331/YAP-2/internal/models"
	"github.com/Danila331/YAP-2/internal/pkg"
	"github.com/labstack/echo/v4"
)

func PageExpression(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("../", "templates", "expression.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}
	cookie, err := c.Cookie("jwtCookie")
	if err != nil {
		return err
	}
	jwtString := cookie.Value
	userLogin, err := pkg.GetUserLoginFromJWT(jwtString)
	if err != nil {
		return err
	}
	var expression models.Expression
	expressions, err := expression.ReadById(userLogin)
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "expressions", expressions)
	return nil
}
