package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connectdb() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:sintara23@localhost:5432/pos")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	//defer dbpool.Close()
	return dbpool
}
