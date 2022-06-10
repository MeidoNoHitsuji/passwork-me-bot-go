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
	usersRouter.HandleFunc("/{id}", GetUserById).Methods("GET")

	groupsRouter := apiRouter.PathPrefix("/groups").Subrouter()
	groupsRouter.HandleFunc("", GetGroups).Methods("GET")

	rolesRouter := apiRouter.PathPrefix("/roles").Subrouter()
	rolesRouter.HandleFunc("", GetRoles).Methods("GET")
	rolesRouter.HandleFunc("/create", CreateRole).Methods("POST")
	rolesRouter.HandleFunc("/{id}", GetRoleById).Methods("GET")
	rolesRouter.HandleFunc("/{id}", UpdateRole).Methods("PATCH")

	permissionsRouter := apiRouter.PathPrefix("/permissions").Subrouter()
	permissionsRouter.HandleFunc("/vault", GetVaultPermissions).Methods("GET")
	permissionsRouter.HandleFunc("/folder", GetFolderPermissions).Methods("GET")

	router.HandleFunc("/", WebIndex).Methods("GET")

	return router
}
