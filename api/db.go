package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/krzysztofkaptur/quotes/api/internal/database"
	_ "github.com/lib/pq"
)

func InitDB() (ApiConfig, error) {
	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_SSL_MODE := os.Getenv("DB_SSL_MODE")

	connStr := fmt.Sprintf("host=%v user=%v dbname=%v password=%v sslmode=%v", DB_HOST, DB_USER, DB_NAME, DB_PASSWORD, DB_SSL_MODE)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return ApiConfig{}, err
	}

	return ApiConfig{
		DB: database.New(conn),
	}, nil
}
