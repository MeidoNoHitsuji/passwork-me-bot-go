package config

var URL = "https://passwork.me"

var DB = map[string]string{
	"drive":    "sqlite3",
	"url":      "database.sqlite",
	"user":     "",
	"password": "",
}

var MigrationsPath = "migrations"
