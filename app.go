package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"passwork-me-bot-go/client"
	"passwork-me-bot-go/config"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	c := client.New(os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	groupId := "62a086ada6fe8016ee05fc54"

	users := c.GroupApi.GetWorkspaceUsersNotInGroup(groupId)

	if len(users) > 0 {
		if c.AddUsersInGroup(users, groupId) {
			fmt.Println("Пользователи добавлены")
		} else {
			fmt.Println("Пользователи не добавлены")
		}
	} else {
		fmt.Println("Не кого добавлять")
	}

	permissions := map[string]string{}

	for _, user := range users {
		permissions[user.Id] = config.FullAccess
	}

	if c.GroupApi.UpdatePermissionsGroup(permissions, groupId) { // У категорий передются параметры в цифровом формате, а у групп в текстовом. Это надо пофиксить.
		fmt.Println("Пермишны обновлены")
	} else {
		fmt.Println("Пермишны не обновлены")
	}

	//db := database.New()
	//c.UpdatePermissions(db)

	//fmt.Println(u.Encode())
	//database.RunMigrateScripts(db)

	//database.New()
}
