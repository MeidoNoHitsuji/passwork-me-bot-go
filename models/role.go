package models

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `json:"name"`
	Users     []User         `gorm:"many2many:user_roles;ForeignKey:id;References:id;" json:"users"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
