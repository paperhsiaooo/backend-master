package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"

	// Import the postgres driver，但不會實際使用它，避免 format 時被移除，所以加上 _
	_ "github.com/lib/pq"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgx.Conn

func TestMain(m *testing.M) {
	ctx := context.Background()
	var err error

	testDB, err = pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB) // Pass conn as a pointer to DBTX

	os.Exit(m.Run())
}
