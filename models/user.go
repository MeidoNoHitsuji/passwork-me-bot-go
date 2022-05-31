package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ServiceID string
	Name      string
	Email     string
	Roles     []Role `gorm:"many2many:user_roles;"`
}
