package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func DBConnection() (*sql.DB, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "reza"
		password = "reza53//"
		dbname   = "reza"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	return db, err
}
