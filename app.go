package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"passwork-me-bot-go/client"
	"passwork-me-bot-go/config"
	"passwork-me-bot-go/database"
	"passwork-me-bot-go/routes"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Email = os.Getenv("EMAIL")
	config.Password = os.Getenv("PASSWORD")
	config.MasterKey = os.Getenv("MASTER_KEY")
}

func main() {

	c := client.Instant()
	db := database.Instant()

	c.UpdatePermissions(db)

	srv := &http.Server{
		Handler: routes.New(),
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	//TODO:
	//1. Создать роуты создания/патча ролей, патча пользователей
	//2. Проверить работоспособность обновления прав относительно ролей
	//3. Либо дописать веб, либо написать отдельный маленький скрипт на дёрганье api.
	//4. Захостить боть
}
