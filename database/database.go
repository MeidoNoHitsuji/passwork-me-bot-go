package database

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"passwork-me-bot-go/config"
	"passwork-me-bot-go/models"
)

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DB["url"]), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ms := []interface{}{
		&models.User{},
		&models.Role{},
		&models.Group{},
		&models.RoleGroupPermissions{},
		&models.UserGroupPermissions{},
	}

	if err = db.AutoMigrate(ms...); err != nil {
		return nil
	}

	return db
}
