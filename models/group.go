package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string
	GroupId     string `gorm:"uniqueIndex"`
	ParentId    string
	ParentGroup *Group `gorm:"foreignKey:ParentId;references:GroupId;"`
	IsVault     bool
}

type RoleGroupPermissions struct {
	GroupID    int
	Group      Group `gorm:"constraint:OnDelete:CASCADE"`
	RoleID     int
	Role       Role `gorm:"constraint:OnDelete:CASCADE"`
	Permission string
}

type UserGroupPermissions struct {
	GroupID    int
	Group      Group `gorm:"constraint:OnDelete:CASCADE"`
	UserID     int
	User       User `gorm:"constraint:OnDelete:CASCADE"`
	Permission string
}
