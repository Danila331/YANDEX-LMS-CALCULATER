package forms

import (
	"time"

	"github.com/Danila331/YAP-2/internal/models"
	"github.com/Danila331/YAP-2/internal/pkg"
	"github.com/labstack/echo/v4"
)

func formExpressionListener(c echo.Context) error {
	expressionString := c.FormValue("expression")
	cookie, err := c.Cookie("jwtCookie")
	if err != nil {
		return err
	}
	jwtString := cookie.Value
	userLogin, err := pkg.GetUserLoginFromJWT(jwtString)
	if err != nil {
		return err
	}
	expressionTime := time.Now().Format("2 Jan 2006 15:04:05")
	timeEnd, timeProccesing, err := pkg.CalcEndTime(expressionString, int(time.Now().Unix()))
	if err != nil {
		return err
	}
	var expression models.Expression
	expression.Expression = expressionString
	expression.TimeStart = expressionTime
	expression.TimeEnd = time.Unix(int64(timeEnd), 0).Format("2 Jan 2006 15:04:05")
	expression.TimeProccesing = timeProccesing
	expression.UserLogin = userLogin
	err = expression.Create()
	if err != nil {
		return err
	}
	return nil
}
