package routes

import (
	"flag"
	"github.com/gorilla/mux"
	"net/http"
)

func New() *mux.Router {

	var dir string
	flag.StringVar(&dir, "dir", "./static/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	apiRouter := router.PathPrefix("/api").Subrouter()

	usersRouter := apiRouter.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", GetUsers).Methods("GET")
	usersRouter.HandleFunc("/", GetUsers).Methods("GET")
	usersRouter.HandleFunc("/{id}", GetUserById).Methods("GET")
	usersRouter.HandleFunc("/{id}/", GetUserById).Methods("GET")

	apiRouter.HandleFunc("/groups", GetGroups).Methods("GET")
	apiRouter.HandleFunc("/roles", GetRoles).Methods("GET")

	router.HandleFunc("/", WebIndex).Methods("GET")

	return router
}
