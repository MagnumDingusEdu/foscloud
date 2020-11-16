package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	err := godotenv.Load("../../config.env")
	if err != nil {
		log.Fatalln("Unable to load .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_BIND"),
			os.Getenv("DB_PORT"),
			os.Getenv("TEST_DB_NAME"),
		)


	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	testQueries = New(testDb)

	os.Exit(m.Run())
}
