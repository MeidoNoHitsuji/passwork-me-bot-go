package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"passwork-me-bot-go/database"
	"passwork-me-bot-go/models"
	"strings"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db := database.Instant()
	db.Model(&models.User{}).Find(&users)

	data, _ := json.Marshal(&users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, string(data))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if vars["id"] == "" {
		data := map[string]interface{}{
			"error": http.StatusText(http.StatusNotFound),
		}
		errorJson, _ := json.Marshal(&data)
		http.Error(w, string(errorJson), http.StatusNotFound)
		return
	}

	db := database.Instant()
	db.Model(&models.User{}).Where("id = ?", vars["id"]).Find(&user)

	if user.ID == "" {
		data := map[string]interface{}{
			"error": http.StatusText(http.StatusNotFound),
		}
		errorJson, _ := json.Marshal(&data)
		http.Error(w, string(errorJson), http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(&user)

	w.WriteHeader(http.StatusOK)

	io.WriteString(w, string(data))
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	var groups []models.Group
	depth := 4

	var preloads []string
	for i := 0; i < depth; i++ {
		preloads = append(preloads, "ParentGroup")
	}

	db := database.Instant()
	db.Model(&models.Group{}).Preload(strings.Join(preloads, ".")).Find(&groups)

	data, _ := json.Marshal(&groups)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, string(data))
}

func GetRoles(w http.ResponseWriter, r *http.Request) {
	var groups []models.Role
	db := database.Instant()
	db.Model(&models.Role{}).Find(&groups)

	data, _ := json.Marshal(&groups)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, string(data))
}
