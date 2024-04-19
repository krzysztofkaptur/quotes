package main

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/krzysztofkaptur/quotes/api/internal/database"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed db/migrations/*.sql
var embedMigrations embed.FS

func InitDB() (ApiConfig, error) {
	// DB_HOST := os.Getenv("DB_HOST")
	// DB_USER := os.Getenv("DB_USER")
	// DB_NAME := os.Getenv("DB_NAME")
	// DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_SSL_MODE := os.Getenv("DB_SSL_MODE")
	secrets := InitAWSEnv()

	var connStr string

	if DB_SSL_MODE != "" {
		connStr = fmt.Sprintf("host=%v user=%v dbname=%v password=%v sslmode=%v", secrets.Host, secrets.Username, secrets.Dbname, secrets.Password, DB_SSL_MODE)
	} else {
		connStr = fmt.Sprintf("host=%v user=%v dbname=%v password=%v", secrets.Host, secrets.Username, secrets.Dbname, secrets.Password)
	}

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return ApiConfig{}, err
	}

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		fmt.Println("SetDialect")
		panic(err)
	}

	if err := goose.Up(conn, "db/migrations"); err != nil {
		fmt.Println("db/migrations")
		panic(err)
	}

	fmt.Println("after db/migrations")

	return ApiConfig{
		DB: database.New(conn),
	}, nil
}
