package config

var URL = "https://passwork.me"

var (
	Email     string
	Password  string
	MasterKey string
)

var DB = map[string]string{
	"drive":    "sqlite3",
	"url":      "database.sqlite",
	"user":     "",
	"password": "",
}

var PermissionValue = map[string]string{
	NoAccess:            "denied",
	ListingAccess:       "listing",
	ReadAccess:          "read",
	EditAccess:          "write",
	FullAccess:          "delete",
	AdministratorAccess: "admin",
}
