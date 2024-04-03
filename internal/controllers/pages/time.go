package pages

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/Danila331/YAP-2/internal/models"
	"github.com/labstack/echo/v4"
)

func PageTime(c echo.Context) error {
	var time models.Time
	time, err := time.ReadAll()
	fmt.Println(err)
	if err != nil {
		return err
	}

	htmlFiles := []string{
		filepath.Join("../", "templates", "time.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	fmt.Println(err)
	if err != nil {
		return err
	}

	templ.ExecuteTemplate(c.Response(), "time", time)
	return nil
}
