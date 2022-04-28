package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	api := DefaultApi{}

	userApi := UserApi{
		d: &api,
	}

	result, err := userApi.Authorize(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
