package pages

import (
	"path/filepath"
	"text/template"

	"github.com/Danila331/YAP-2/internal/models"
	"github.com/labstack/echo/v4"
)

func PageWorkers(c echo.Context) error {
	htmlFiles := []string{
		filepath.Join("../", "templates", "workers.html"),
	}

	templ, err := template.ParseFiles(htmlFiles...)
	if err != nil {
		return err
	}

	var worker models.Worker
	workers, err := worker.ReadAll()
	if err != nil {
		return err
	}
	templ.ExecuteTemplate(c.Response(), "workers", workers)
	return nil
}
