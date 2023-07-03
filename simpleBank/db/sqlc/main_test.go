package db

import (
	"database/sql"
	"testing"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQuesries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err!= nil {
		log.Fatal("connot connect to db")
	}
	testQuesries = New(conn)

	os.Exit(m.Run())
}