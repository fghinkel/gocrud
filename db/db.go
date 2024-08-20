package db

// Go only let us import packages that are in use, sometimes, packages aren't always in use,
// in this cases we can pass the _ before the package name
import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectWithDatabase() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost port=5488 dbname=gocrud sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	return db
}
