package models

import (
	"github.com/Danila331/YAP-2/internal/store"
)

type Expression struct {
	Id             int
	UserLogin      string
	Expression     string
	Value          int
	TimeStart      string
	TimeProccesing int
	TimeEnd        string
	IsProcessing   string
}

type ExpressionInterface interface {
	Create() error
	ReadById(string) (Expression, error)
	ReadAll(string) ([]Expression, error)
}

func (e *Expression) Create() error {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer conn.Close()
	_ = conn.QueryRow("INSERT INTO expressions(userlogin, expression, value, timestart, timeend, timeproccesing, isprocessing) VALUES (?,?,?,?,?,?,?)",
		e.UserLogin,
		e.Expression,
		0,
		e.TimeStart,
		e.TimeEnd,
		e.TimeProccesing,
		"process",
	)
	return nil
}

func (e *Expression) ReadAll(userLogin string) ([]Expression, error) {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return []Expression{}, err
	}
	defer conn.Close()
	rows, err := conn.Query("SELECT * FROM expressions WHERE userlogin = ?", userLogin)
	if err != nil {
		return []Expression{}, err
	}
	defer rows.Close()
	var expressions []Expression
	for rows.Next() {
		var expression Expression
		err = rows.Scan(
			&expression.Expression,
			&expression.Value,
			&expression.TimeStart,
			&expression.IsProcessing,
		)
		if err != nil {
			return []Expression{}, err
		}
		expressions = append(expressions, expression)
	}
	return expressions, nil
}

func (e *Expression) ReadById(userLogin string) (Expression, error) {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return Expression{}, err
	}
	defer conn.Close()
	var expression Expression
	err = conn.QueryRow("SELECT * FROM expressions WHERE userlogin = ? AND id = ?", userLogin, e.Id).Scan(expression.Expression, expression.Value, expression.TimeStart, expression.IsProcessing)
	if err != nil {
		return Expression{}, err
	}
	return expression, nil
}
