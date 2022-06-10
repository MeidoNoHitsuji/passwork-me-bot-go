package client

import (
	"passwork-me-bot-go/api"
	"passwork-me-bot-go/database"
	"passwork-me-bot-go/models"
)

func UpdateUserPermissionsByRoles(u models.User) {
	db := database.Instant()
	if len(u.Roles) == 0 {
		db.Model(&models.User{}).Association("Roles").Find(&u.Roles)
		if len(u.Roles) == 0 {
			return
		}
	}

	var permissions map[models.Group]string
	var groups []models.Group

	for _, role := range u.Roles {
		var perm []models.RoleGroupPermissions
		db.Model(&models.RoleGroupPermissions{}).Where("role_id = ?", role.ID).Preload("Group.ParentGroup").Find(&perm)
		for _, p := range perm {
			if permissions[p.Group] != "" {
				//TODO: Сортировка по возрастанию прав
			} else {
				permissions[p.Group] = p.Permission
				contains := false
				for _, a := range groups {
					if a == p.Group {
						contains = true
						break
					}
				}
				if !contains {
					groups = append(groups, p.Group)
				}
			}
		}
	}

	if len(permissions) > 0 {
		c := Instant()

		for _, group := range groups {
			users := c.GroupApi.GetWorkspaceUsersNotInGroup(group.ID)
			var contains []api.UserWithPublicKey
			for _, a := range users {
				if a.Id == u.ID {
					contains = append(contains, a)
				}
			}

			if len(contains) != 0 {
				g := c.GroupApi.GetFullData(group.ID)
				c.WorkspaceApi.AddRsaEncryptedGroupToManyUsers(contains, g.Group.GroupId, g.Group.DecryptPassword())
			}
		}

		for group, permission := range permissions {
			if group.ParentGroup != nil {
				c.GroupApi.UpdatePermissionsFolder(map[string]string{
					u.ID: permission,
				}, group.ParentGroup.ID, group.ID)
			} else {
				c.GroupApi.UpdatePermissionsGroup(map[string]string{
					u.ID: permission,
				}, group.ParentGroup.ID)
			}
		}
	}

}
