package main

import (
	"fmt"
	"net/http"
)

func (server *ApiServer) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/hello", makeHTTPHandleFunc(server.hello))
	router.HandleFunc("GET /api/v1/authors", makeHTTPHandleFunc(server.handleFetchAuthors))
	router.HandleFunc("GET /api/v1/quotes", makeHTTPHandleFunc(server.handleFetchQuotes))
	router.HandleFunc("GET /api/v1/quotes/random", makeHTTPHandleFunc(server.handleFetchRandomQuote))

	http.ListenAndServe(fmt.Sprintf(":%v", server.address), router)
}

func (server *ApiServer) hello(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, struct{ Message string }{Message: "herro with aws secrets"})
}

func (server *ApiServer) handleFetchAuthors(w http.ResponseWriter, r *http.Request) error {
	authors, err := server.store.DB.FetchAuthors(r.Context())
	if err != nil {
		fmt.Println(err)
		return WriteJSON(w, http.StatusBadRequest, ApiError{Error: "something went wrong"})
	}

	return WriteJSON(w, http.StatusOK, authors)
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
