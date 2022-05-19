package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"passwork-me-bot-go/client"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	c := client.New(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	resp := c.PasswordApi.GetRecentPasswords()

	var b []byte

	b, _ = json.Marshal(resp)

	fmt.Println(string(b))

	c.GroupApi.GetFullData("qwewq")
}
