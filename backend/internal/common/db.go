package common

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL") // or build from env vars
	return sql.Open("postgres", connStr)
}
