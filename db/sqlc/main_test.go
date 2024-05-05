package db

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	// Import the postgres driver，但不會實際使用它，避免 format 時被移除，所以加上 _
	_ "github.com/lib/pq"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()

	if config, err := pgxpool.ParseConfig(dbSource); err != nil {
		os.Exit(1)
	} else {
		config.MaxConns = 10

		var err error

		testDB, err = pgxpool.NewWithConfig(ctx, config)
		if err != nil {
			os.Exit(1)
		}

		testQueries = New(testDB) // Pass conn as a pointer to DBTX
		os.Exit(m.Run())
	}
}
