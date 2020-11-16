package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	fmt.Printf(os.Getenv("DB_NAME"))
}
