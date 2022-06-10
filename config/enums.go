package config

const (
	NoAccess            = iota
	ListingAccess       = iota
	ReadAccess          = iota
	EditAccess          = iota
	FullAccess          = iota
	AdministratorAccess = iota
	InheritedAccess     = iota
)

var VaultPermission = map[int]string{
	NoAccess:            "denied",
	ListingAccess:       "listing",
	ReadAccess:          "read",
	EditAccess:          "write",
	FullAccess:          "delete",
	AdministratorAccess: "admin",
}

var FolderPermission = map[int]string{
	NoAccess:        "0",
	ListingAccess:   "0.5",
	ReadAccess:      "1",
	EditAccess:      "2",
	FullAccess:      "3",
	InheritedAccess: "-1",
}
