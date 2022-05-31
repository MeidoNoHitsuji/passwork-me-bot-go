package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Users []User `gorm:"many2many:user_roles;"`
}
