package database

import (
	"database/sql"
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"passwork-me-bot-go/config"
)

//go:embed migrations/*.sql
var fs embed.FS

func New() *sql.DB {
	sqliteDb, err := sql.Open(config.DB["drive"], config.DB["url"])
	if err != nil {
		panic("Failed to open sqlite DB")
	}

	return sqliteDb
}

// RunMigrateScripts
//
// TODO: Если Drive изменится, то изменить реализацию данной функции
func RunMigrateScripts(db *sql.DB) {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		panic(err)
	}

	d, err := iofs.New(fs, config.MigrationsPath)
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithInstance("iofs", d, "main", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil {
		// ...
	}

}
