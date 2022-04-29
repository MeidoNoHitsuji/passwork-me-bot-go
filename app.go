package main

import (
	"encoding/json"
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

	_, err := userApi.Authorize(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	if err != nil {
		fmt.Println(err)
	}

	if err := client.initCsrf(); err != nil {
		panic(err.Error())
	}

	resp := userApi.GetInfo()

	if err != nil {
		fmt.Println(err)
	}

	var b []byte

	b, err = json.Marshal(resp)
	
	fmt.Println(string(b))
}
