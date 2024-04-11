package main

import (
	"encoding/json"
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
	err := InitEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	server := ApiServer{
		address: os.Getenv("PORT"),
		store:   db,
	}

	server.Run()
}

func (server *ApiServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/hello", makeHTTPHandleFunc(server.hello))

	router.HandleFunc("GET /api/v1/authors", makeHTTPHandleFunc(server.handleFetchAuthors))

	http.ListenAndServe(fmt.Sprintf(":%v", server.address), router)
}

func (server *ApiServer) hello(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, struct{ Message string }{Message: "herro"})
}

func (server *ApiServer) handleFetchAuthors(w http.ResponseWriter, r *http.Request) error {
	authors, err := server.store.DB.FetchAuthors(r.Context())

	if err != nil {
		fmt.Println(err)
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "something went wrong"})
	}

	return WriteJSON(w, http.StatusOK, authors)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json="error"`
	Code  int32  `json="code"`
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
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
