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
	db       *sql.DB
	host     string
	port     string
	database string
	username string
	password string
}

func NewStore(host, port, database, username, password string) Store {
	s := new(store)
	s.host = host
	s.port = port
	s.database = database
	s.username = username
	s.password = password

	return s
}

func (s *store) getConString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", s.username, s.password, s.host, s.port, s.database)

}

func (s *store) Connect() error {
	db, err := sql.Open(
		"postgres",
		s.getConString(),
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
