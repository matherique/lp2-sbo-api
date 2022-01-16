package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store interface {
	Connect() error
	Close() error
}

type store struct {
	db        *sql.DB
	conString string
}

func NewStore(conString string) Store {
	s := new(store)
	s.conString = conString

	return s
}

func (s *store) Connect() error {
	db, err := sql.Open(
		"postgres",
		s.conString,
	)

	if err != nil {
		return fmt.Errorf("could not open database connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return fmt.Errorf("could not ping to database: %v", err)
	}

	s.db = db
	return nil
}

func (s *store) Close() error {
	return s.db.Close()
}
