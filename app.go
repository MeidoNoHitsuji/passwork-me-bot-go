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
	api := Requester{}

	client := Client{
		api: &api,
	}

	userApi := UserApi{
		api: &api,
	}

	result, err := userApi.Authorize(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	if err := client.initCsrf(); err != nil {
		fmt.Println(err)
	}
}
