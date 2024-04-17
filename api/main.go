package main

import (
	"fmt"
	"log"
	"net/http"
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
	fmt.Println("setup")
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

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json="error"`
	Code  int32  `json="code"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{
				Error: err.Error(),
			})
		}
	}
}
