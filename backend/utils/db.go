package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	dbURL := os.Getenv("DATABASE_URL")
	var err error
	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}
	fmt.Println("Connected to the database!")
}
