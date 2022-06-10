package models

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID          string         `gorm:"uniqueIndex" json:"id"`
	Name        string         `json:"name"`
	ParentId    string         `json:"parent_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ParentGroup *Group         `gorm:"foreignKey:ParentId;references:ID;" json:"parent_group"`
}

type RoleGroupPermissions struct {
	GroupID    string `json:"group_id"`
	RoleID     uint   `json:"role_id"`
	Permission string `json:"permission"`
	Group      Group  `gorm:"foreignKey:GroupID;references:ID;constraint:OnDelete:CASCADE" json:"group"`
	Role       Role   `gorm:"constraint:OnDelete:CASCADE" json:"role"`
}

func (s RoleGroupPermissions) ToResponse() map[string]interface{} {
	return map[string]interface{}{
		"group_id":       s.GroupID,
		"group_name":     s.Group.Name,
		"permission":     s.Permission,
		"group_is_vault": s.Group.ParentGroup == nil,
	}
}

type UserGroupPermissions struct {
	GroupID    string `json:"group_id"`
	Permission string `json:"permission"`
	UserID     string `json:"user_id"`
	Group      Group  `gorm:"foreignKey:GroupID;references:ID;constraint:OnDelete:CASCADE" json:"group"`
	User       User   `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user"`
}
