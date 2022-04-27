package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	api := DefaultApi{
		Url: "test",
	}

	request, err := api.Request("GET", "", map[string]interface{}{})
	if err != nil {
		return
	}

	fmt.Println(request)

	request, err = api.Request("GET", "", map[string]interface{}{})
	if err != nil {
		return
	}

	fmt.Println(request)
}
