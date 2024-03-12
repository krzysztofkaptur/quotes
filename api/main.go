package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/hello", hello)

	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct{Message string `json:"message"`}{Message: "hello"})
}