package models

import (
	"fmt"

	"github.com/Danila331/YAP-2/internal/store"
)

type Time struct {
	TimePulse int
	TimeMinus int
	TimeProz  int
	TimeDel   int
}

type TimeInterface interface {
	Update() error
	ReadAll() (Time, error)
}

func (t *Time) Update() error {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer conn.Close()
	_ = conn.QueryRow("UPDATE time SET timepulse = ?, timeminus = ?, timeproz = ?, timedel = ?",
		t.TimePulse,
		t.TimeMinus,
		t.TimeProz,
		t.TimeDel)

	return nil
}

func (t *Time) ReadAll() (Time, error) {
	conn, err := store.ConnectToDatabase()
	fmt.Println(err)
	if err != nil {
		return Time{}, err
	}
	defer conn.Close()
	var time Time
	err = conn.QueryRow("SELECT * FROM time").Scan(&time.TimePulse,
		&time.TimeMinus,
		&time.TimeProz,
		&time.TimeDel)
	fmt.Println(err)
	if err != nil {
		return Time{}, err
	}
	return time, nil
}
