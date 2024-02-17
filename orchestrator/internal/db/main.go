package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func NewDB(path string) (*DB, error) {
	tasksScheme := `CREATE TABLE IF NOT EXISTS tasks (
		id text NOT NULL PRIMARY KEY,
		expression text,
		result int,
		agentId int,
		status text,
		FOREIGN KEY (agentId) REFERENCES agents(id)
	);`
	agentsScheme := `CREATE TABLE IF NOT EXISTS agents (
		id INTEGER NOT NULL PRIMARY KEY,
		ping int,
		maxThreads int,
		execThreads int,
		url text UNIQUE
	);`
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(tasksScheme)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(agentsScheme)
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}
