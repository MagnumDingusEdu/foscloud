package main

import (
	"database/sql"
	"fmt"
	"foscloud/api"
	db "foscloud/db/sqlc"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_BIND"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("Could not connect to db, err : ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(os.Getenv("SERVER_BIND"))
	if err != nil {
		log.Fatalln("Unable to start server, err : ", err)
	}
	fmt.Printf(os.Getenv("DB_NAME"))
}
