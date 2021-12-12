package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store interface {
	Connect() error
}

type store struct {
	db *sql.DB
}

func NewStore() Store {
	s := new(store)

	return s
}

func (s *store) Connect() error {
	db, err := sql.Open("postgres", "postgres://postgres:123@database/sbo?sslmode=disable")

	if err = db.Ping(); err != nil {
		db.Close()
		return fmt.Errorf("could not connecto to database: %v", err)
	}

	s.db = db
	return nil
}
