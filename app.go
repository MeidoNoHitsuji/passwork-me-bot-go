package main

import (
	"fmt"
	"io"
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
	resp, err := http.Get("http://zaza.local/test")

	if err != nil {
		log.Fatal("Err")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Err Close")
		}
	}(resp.Body)

	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
}
