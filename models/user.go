package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Roles     []Role         `gorm:"many2many:user_roles;ForeignKey:id;References:id;"`
}
