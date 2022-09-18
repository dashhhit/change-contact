package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func DB() *sql.DB {
	connStr := "host=host.docker.internal port=5432 user=spuser password=SPuser96 dbname=project sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	return db
}
