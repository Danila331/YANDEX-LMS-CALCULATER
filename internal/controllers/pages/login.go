package pages

import (
	"html/template"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func PageLogin(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("../", "templates", "login.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "login", nil)
	return nil
}
