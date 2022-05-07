package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Migrate() error {
	m, err := migrate.New(
		"file://migrations",
		os.Getenv("POSTGRES_URL"))

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

var DB *gorm.DB

func InitDatabase() {
	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_URL")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db
}
