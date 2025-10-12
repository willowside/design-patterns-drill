package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" // prefix it with _ because it's not explicitly used here
)

const (
	maxOpenDBConn = 25
	maxIdleDBConn = 25
	maxDBLifeTime = 5 * time.Minute
)

func initMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// test our db by pinging it
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetMaxIdleConns(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDBLifeTime)

	return db, nil
}
