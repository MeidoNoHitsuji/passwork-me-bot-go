package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		log.Fatal("Err")
	}

	resp.Cookies()
}
