package models

import "github.com/Danila331/YAP-2/internal/store"

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
	_ = conn.QueryRow("UPDATE times SET timepulse = ?, timeminus = ?, timeproz = ?, timedel = ?",
		t.TimePulse,
		t.TimeMinus,
		t.TimeProz,
		t.TimeDel)

	return nil
}

func (t *Time) ReadAll() (Time, error) {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return Time{}, err
	}
	defer conn.Close()
	var time Time
	err = conn.QueryRow("SELECT * FROM times").Scan(&time.TimePulse,
		&time.TimeMinus,
		&time.TimeProz,
		&time.TimeDel)

	if err != nil {
		return Time{}, err
	}
	return time, nil
}
