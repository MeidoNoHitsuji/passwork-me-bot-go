package config

var URL = "https://passwork.me"

var DB = map[string]string{
	"drive":    "sqlite3",
	"url":      "database.sqlite",
	"user":     "",
	"password": "",
}

var LANG = map[string]string{
	noAccess:        "No access",
	listingAccess:   "Folder listing",
	readAccess:      "Read only",
	editAccess:      "Read & Edit",
	deleteAccess:    "Full access",
	fullAccess:      "Administrator",
	inheritedAccess: "Inherited from parent",
}
