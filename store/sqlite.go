package store

import (
	"database/sql"

	"github.com/Babur171/carZone-golang/config"
	_ "github.com/mattn/go-sqlite3"
)

type SQlite struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*SQlite, error) {
	db, err := sql.Open("sqlite3", cfg.DBURL)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
		)`)
	if err != nil {
		return nil, err
	}
	return &SQlite{
		DB: db,
	}, nil

}
func (s *SQlite) CreateStudent(name string, email string) (int64, error) {

	stmt, err := s.DB.Prepare("INSERT INTO students (name, email) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email)

	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}
