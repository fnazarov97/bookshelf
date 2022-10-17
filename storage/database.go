package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	user     = "postgres"
	password = "666997"
	database = "bookshelf"
	port     = 5432
)

func Connect() (*sqlx.DB, error) {
	connection := fmt.Sprintf(
		"host = %s user = %s password = %s database = %s port = %d",
		host, user, password, database, port,
	)
	db, err := sqlx.Open("postgres", connection)
	if err != nil {
		fmt.Println(err)
		return db, nil
	}
	return db, nil

}
