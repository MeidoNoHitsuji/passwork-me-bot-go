package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Roles     []Role         `gorm:"many2many:user_roles;ForeignKey:id;References:id;" json:"roles"`
}
