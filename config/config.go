package config

import (
	"os"
)

var URL = "https://passwork.me"

var DB = map[string]string{
	"drive":    "sqlite3",
	"url":      "database.sqlite",
	"user":     "",
	"password": "",
}

var LANG = map[string]string{
	NoAccess:        "No access",
	ListingAccess:   "Folder listing",
	ReadAccess:      "Read only",
	EditAccess:      "Read & Edit",
	DeleteAccess:    "Full access",
	FullAccess:      "Administrator",
	InheritedAccess: "Inherited from parent",
}

var (
	Email     = os.Getenv("EMAIL")
	Password  = os.Getenv("PASSWORD")
	MasterKey = os.Getenv("MASTER_KEY")
)
