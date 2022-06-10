package config

var PermissionName = map[int]string{
	NoAccess:            "No access",
	ListingAccess:       "Folder listing",
	ReadAccess:          "Read only",
	EditAccess:          "Read & Edit",
	FullAccess:          "Full access",
	AdministratorAccess: "Administrator",
	InheritedAccess:     "Inherited from parent",
	//RemoveAccess:        "Remove",
}
