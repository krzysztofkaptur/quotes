package main

import (
	"github.com/joho/godotenv"
)

func InitEnv() error {
	err := godotenv.Load()

	return err
}
