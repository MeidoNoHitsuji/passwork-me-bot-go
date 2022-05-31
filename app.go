package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"passwork-me-bot-go/client"
	"passwork-me-bot-go/database"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	c := client.New(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	db := database.New()
	c.UpdatePermissions(db)

	fmt.Println("kekw")
	//database.RunMigrateScripts(db)

	//database.New()
}
