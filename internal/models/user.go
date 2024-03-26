package models

import (
	"github.com/Danila331/YAP-2/internal/store"
)

type User struct {
	ID       int
	Login    string
	Password string
}

type UserInterface interface {
	Create() error
	ReadByLogin() (User, error)
}

func (u *User) Create() error {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Query("INSERT INTO users(login, password) VALUES (?,?)", u.Login, u.Password)

	if err != nil {
		return err
	}

	return nil
}

func (u *User) ReadByLogin() (User, error) {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()
	var user User
	err = conn.QueryRow("SELECT * FROM users WHERE login = ?", u.Login).Scan(&user.Login, &user.Password)
	return user, nil
}
