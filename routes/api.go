package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"passwork-me-bot-go/config"
	"passwork-me-bot-go/database"
	"passwork-me-bot-go/helper"
	"passwork-me-bot-go/models"
	"strings"
)

func httpJsonException(w http.ResponseWriter, message string, status int) {
	data := map[string]interface{}{
		"error": message,
	}
	errorJson, _ := json.Marshal(&data)
	http.Error(w, string(errorJson), status)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db := database.Instant()
	db.Model(&models.User{}).Preload("Roles").Find(&users)

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
		httpJsonException(w, "Вы не передали параметр id", http.StatusNotFound)
		return
	}

	db := database.Instant()
	db.Model(&models.User{}).Where("id = ?", vars["id"]).Find(&user)

	if user.ID == "" {
		httpJsonException(w, "Пользователь с таким id не найден", http.StatusNotFound)
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
	var roles []models.Role

	type responseType struct {
		models.Role
		Groups []map[string]interface{} `json:"groups"`
	}

	var response []responseType

	db := database.Instant()
	db.Model(&models.Role{}).Find(&roles)

	for _, role := range roles {

		var rolePermissions []models.RoleGroupPermissions

		d := responseType{
			role,
			[]map[string]interface{}{},
		}

		db.Model(&models.RoleGroupPermissions{}).Preload("Group.ParentGroup").Where("role_id = ?", role.ID).Find(&rolePermissions)

		for _, p := range rolePermissions {
			d.Groups = append(d.Groups, p.ToResponse())
		}

		response = append(response, d)
	}

	data, _ := json.Marshal(&response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, string(data))
}

func GetRoleById(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if vars["id"] == "" {
		httpJsonException(w, "Вы не передали параметр id", http.StatusNotFound)
		return
	}

	db := database.Instant()
	db.Model(&models.Role{}).Where("id = ?", vars["id"]).Find(&role)

	if role.ID == 0 {
		httpJsonException(w, "Роль с таким id не найдена", http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(&role)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(data))
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role

	decoder := json.NewDecoder(r.Body)
	var vars map[string]interface{}

	err := decoder.Decode(&vars)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")

	if vars["name"] == nil || vars["name"] == "" {
		httpJsonException(w, "Вы не передали параметр name", http.StatusNotFound)
		return
	}

	db := database.Instant()
	db.Model(&models.Role{}).Where("name = ?", vars["name"].(string)).Find(&role)

	if role.ID != 0 {
		httpJsonException(w, "Роль с таким именем уже существует", http.StatusConflict)
		return
	}

	role = models.Role{
		Name: vars["name"].(string),
	}

	db.Create(&role)

	data, _ := json.Marshal(&role)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(data))
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	var group models.Group
	vars := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")

	if vars["role_id"] == "" {
		httpJsonException(w, "Вы не передали параметр role_id", http.StatusNotFound)
		return
	}

	if vars["group_id"] == "" {
		httpJsonException(w, "Вы не передали параметр group_id", http.StatusNotFound)
		return
	}

	if vars["permission"] == "" {
		httpJsonException(w, "Вы не передали параметр permission", http.StatusNotFound)
		return
	}

	db := database.Instant()

	db.Model(&models.Role{}).Where("id = ?", vars["role_id"]).Find(&role)

	if role.ID == 0 {
		httpJsonException(w, "Роль с таким id не найдена", http.StatusNotFound)
		return
	}

	db.Model(&models.Group{}).Preload("ParentGroup").Where("id = ?", vars["group_id"]).Find(&group)

	if group.ID == "" {
		httpJsonException(w, "Группа с таким id не найдена", http.StatusNotFound)
		return
	}

	parentGroup := group.ParentGroup.ID
	var permissions []string

	if parentGroup == "" {
		permissions = helper.GetPermissions(true)
	} else {
		permissions = helper.GetPermissions(false)
	}

	if !helper.Contains(permissions, vars["permission"]) {
		httpJsonException(w, "Переданное вами право не было найдено", http.StatusNotFound)
	}

	perm := models.RoleGroupPermissions{
		GroupID: group.ID,
		RoleID:  role.ID,
	}

	db.FirstOrInit(&perm, models.RoleGroupPermissions{
		Permission: vars["permission"],
	}).Save(perm)

	data, _ := json.Marshal(&perm)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(data))
}

func GetVaultPermissions(w http.ResponseWriter, r *http.Request) {

	response := map[string]interface{}{}

	permissions := helper.GetPermissionsKeys(true)

	for _, permission := range permissions {
		response[config.VaultPermission[permission]] = config.PermissionName[permission]
	}

	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(&response)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(data))
}

func GetFolderPermissions(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{}

	permissions := helper.GetPermissionsKeys(false)

	for _, permission := range permissions {
		response[config.FolderPermission[permission]] = config.PermissionName[permission]
	}

	w.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(&response)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(data))
}
