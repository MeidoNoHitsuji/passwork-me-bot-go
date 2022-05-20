package main

import (
	"github.com/joho/godotenv"
	"log"
	"passwork-me-bot-go/database"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	//c := client.New(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	//
	//resp := c.WorkspaceApi.GetUsers()
	//
	//var b []byte
	//
	//b, _ = json.Marshal(resp)
	//
	//fmt.Println(string(b))

	db := database.New()
	database.RunMigrateScripts(db)

	//database.New()
}
