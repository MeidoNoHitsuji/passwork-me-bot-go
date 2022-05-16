package main

import (
	"encoding/json"
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
	client := Client{}

	if err := client.init(); err != nil {
		panic(err)
	}

	resp := client.groupApi.Get()

	var b []byte

	b, _ = json.Marshal(resp)

	fmt.Println(string(b))
}
