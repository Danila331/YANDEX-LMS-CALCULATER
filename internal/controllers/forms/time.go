package forms

import (
	"strconv"

	"github.com/Danila331/YAP-2/internal/models"
	"github.com/labstack/echo/v4"
)

func formTimelistener(c echo.Context) error {
	timePulse, _ := strconv.Atoi(c.FormValue("timePulse"))
	timeMinus, _ := strconv.Atoi(c.FormValue("timeMinus"))
	timeProz, _ := strconv.Atoi(c.FormValue("timeProz"))
	timeDel, _ := strconv.Atoi(c.FormValue("timeDel"))
	var time models.Time
	time.TimePulse = timePulse
	time.TimeMinus = timeMinus
	time.TimeProz = timeProz
	time.TimeDel = timeDel
	err := time.Update()
	if err != nil {
		return err
	}
	return nil
}
