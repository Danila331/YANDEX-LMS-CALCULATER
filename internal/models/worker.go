package models

import (
	"github.com/Danila331/YAP-2/internal/store"
)

type Worker struct {
	Id            int
	Host          string
	IdsExpression string
}

type WorkerInterface interface {
	Create() error
	ReadAll() ([]Worker, error)
	ReadById(int) (Worker, error)
	Delete() error
}

func (w *Worker) Create() error {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Query("INSERT INTO workers(host, idsexpression) VALUES (?,?)", w.Host, w.IdsExpression)
	if err != nil {
		return err
	}
	return nil
}

func (w *Worker) ReadAll() ([]Worker, error) {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return []Worker{}, err
	}
	defer conn.Close()
	var workers []Worker
	rows, err := conn.Query("SELECT * FROM workers")
	if err != nil {
		return []Worker{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var worker Worker
		err = rows.Scan(&worker.Id, &worker.Host, &worker.IdsExpression)
		if err != nil {
			return []Worker{}, err
		}
		workers = append(workers, worker)
	}
	return workers, nil
}

func (w *Worker) ReadById() (Worker, error) {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return Worker{}, err
	}
	defer conn.Close()
	var worker Worker
	err = conn.QueryRow("SELECT * FROM workers WHERE id = ?", w.Id).Scan(&worker.Host, &worker.IdsExpression)
	if err != nil {
		return Worker{}, err
	}
	return worker, nil
}

func (w *Worker) Delete() error {
	conn, err := store.ConnectToDatabase()
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.Query("DELETE * FROM workers WHERE id = ?", w.Id)
	if err != nil {
		return err
	}
	return nil
}
