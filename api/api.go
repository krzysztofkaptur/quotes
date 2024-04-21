package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krzysztofkaptur/quotes/api/internal/database"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
	Code  int32  `json:"code"`
}

type ApiGenericRes struct {
	Message string `json:"string"`
}

func (server *ApiServer) Run() {
	router := http.NewServeMux()

	// test
	router.HandleFunc("GET /api/v1/hello", makeHTTPHandleFunc(server.hello))

	// authors
	router.HandleFunc("GET /api/v1/authors", makeHTTPHandleFunc(server.handleFetchAuthors))
	router.HandleFunc("POST /api/v1/authors", makeHTTPHandleFunc(server.handleCreateAuthor))

	// quotes
	router.HandleFunc("GET /api/v1/quotes", makeHTTPHandleFunc(server.handleFetchQuotes))
	router.HandleFunc("GET /api/v1/quotes/random", makeHTTPHandleFunc(server.handleFetchRandomQuote))
	router.HandleFunc("POST /api/v1/quotes", makeHTTPHandleFunc(server.handleCreateQuote))

	http.ListenAndServe(fmt.Sprintf(":%v", server.address), router)
}

func (server *ApiServer) hello(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, ApiGenericRes{Message: "herro with aws secrets"})
}

func (server *ApiServer) handleFetchAuthors(w http.ResponseWriter, r *http.Request) error {
	authors, err := server.store.DB.FetchAuthors(r.Context())
	if err != nil {
		fmt.Println(err)
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "something went wrong"})
	}

	return WriteJSON(w, http.StatusOK, authors)
}

func (server *ApiServer) handleCreateAuthor(w http.ResponseWriter, r *http.Request) error {
	type reqBody struct {
		Name string `json:"name"`
	}

	params := reqBody{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println(err)
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "something went wrong"})
	}

	server.store.DB.CreateAuthor(r.Context(), params.Name)

	return WriteJSON(w, http.StatusCreated, ApiGenericRes{Message: "new user created"})
}

func (server *ApiServer) handleFetchQuotes(w http.ResponseWriter, r *http.Request) error {
	quotes, err := server.store.DB.FetchQuotes(r.Context())
	if err != nil {
		fmt.Println(err)
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "something went wrong"})
	}

	return WriteJSON(w, http.StatusOK, quotes)
}

func (server *ApiServer) handleFetchRandomQuote(w http.ResponseWriter, r *http.Request) error {
	quote, err := server.store.DB.FetchRandomQuote(r.Context())
	if err != nil {
		fmt.Println(err)
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "something went wrong"})
	}

	return WriteJSON(w, http.StatusOK, quote)
}

func (server *ApiServer) handleCreateQuote(w http.ResponseWriter, r *http.Request) error {
	type parameters struct {
		AuthorID int32  `json:"author_id"`
		Text     string `json:"text"`
	}

	params := parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		fmt.Println(err)
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "something went wrong"})
	}

	err = server.store.DB.CreateQuote(r.Context(), database.CreateQuoteParams{AuthorID: params.AuthorID, Text: params.Text})
	if err != nil {
		fmt.Println(err)
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "something went wrong"})
	}

	return WriteJSON(w, http.StatusCreated, ApiGenericRes{Message: "new quote created"})
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
