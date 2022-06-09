package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"passwork-me-bot-go/client"
	"passwork-me-bot-go/database"
	"passwork-me-bot-go/routes"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

	var dir string
	flag.StringVar(&dir, "dir", "./static/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	c := client.Instant()
	db := database.Instant()

	c.UpdateUsers(db)

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	router.HandleFunc("/", routes.WebIndex).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

	//groupId := "62a086ada6fe8016ee05fc54"
	//
	//users := connect.GroupApi.GetWorkspaceUsersNotInGroup(groupId)
	//
	//if len(users) > 0 {
	//	if connect.AddUsersInGroup(users, groupId) {
	//		fmt.Println("Пользователи добавлены")
	//	} else {
	//		fmt.Println("Пользователи не добавлены")
	//	}
	//} else {
	//	fmt.Println("Не кого добавлять")
	//}
	//
	//permissions := map[string]string{}
	//
	//for _, user := range users {
	//	permissions[user.Id] = config.FullAccess
	//}
	//
	//if connect.GroupApi.UpdatePermissionsGroup(permissions, groupId) { // У категорий передются параметры в цифровом формате, а у групп в текстовом. Это надо пофиксить.
	//	fmt.Println("Пермишны обновлены")
	//} else {
	//	fmt.Println("Пермишны не обновлены")
	//}

	//c.UpdatePermissions(db)

	//fmt.Println(u.Encode())
	//database.RunMigrateScripts(db)

	//database.New()
}
