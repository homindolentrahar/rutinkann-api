package local

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Storage interface {
	Connect() (*sql.DB, error)
}

type PostgresStorage struct {
	Connection string
	Driver     string
}

func NewPostgresStorage() *PostgresStorage {
	return &PostgresStorage{
		Connection: "postgres://postgres:root@localhost/rutinkann?sslmode=disable",
		Driver:     "postgres",
	}
}

func (ps *PostgresStorage) Connect() (*sql.DB, error) {
	db, err := sql.Open(ps.Driver, ps.Connection)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatal(err)

		return nil, err
	}

	return db, nil
}
