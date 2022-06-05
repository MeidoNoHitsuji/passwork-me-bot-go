package models

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID          string `gorm:"uniqueIndex"`
	Name        string
	ParentId    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	ParentGroup *Group         `gorm:"foreignKey:ParentId;references:ID;"`
}

type RoleGroupPermissions struct {
	GroupID    string
	RoleID     int
	Permission string
	Group      Group `gorm:"foreignKey:GroupID;references:ID;constraint:OnDelete:CASCADE"`
	Role       Role  `gorm:"constraint:OnDelete:CASCADE"`
}

type UserGroupPermissions struct {
	GroupID    string
	Permission string
	UserID     string
	Group      Group `gorm:"foreignKey:GroupID;references:ID;constraint:OnDelete:CASCADE"`
	User       User  `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}
