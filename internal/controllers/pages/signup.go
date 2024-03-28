package pages

import (
	"html/template"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func PageSignup(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("../", "templates", "signup.html"),
	}
	templ, err := template.ParseFiles(htmlFiles...)

	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "signup", nil)
	return nil
}
