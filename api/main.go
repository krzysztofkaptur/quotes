package main

import (
	"fmt"
	"log"
	"os"

	"github.com/krzysztofkaptur/quotes/api/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

type ApiServer struct {
	address string
	store   ApiConfig
}

func main() {
	err := InitEnv()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	db, err := InitDB()
	if err != nil {
		fmt.Println("before InitDB")
		log.Fatal(err)
		fmt.Println("after InitDB")
	}

	server := ApiServer{
		address: os.Getenv("PORT"),
		store:   db,
	}

	server.Run()
}
