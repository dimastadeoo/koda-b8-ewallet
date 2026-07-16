package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Conn() *pgx.Conn {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot read file ekstensions .env")
	}

	getCon := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), getCon)
	if err != nil {
		fmt.Println("Cannot connect database")
		os.Exit(1)
	}
	return conn
}
