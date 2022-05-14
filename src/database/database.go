package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

var MinioClient *minio.Client

func InitMinio() {
	endpoint := os.Getenv("MINIO_URL")
	accessKeyID := os.Getenv("MINIO_KEY_ID")
	secretAccessKey := os.Getenv("MINIO_SECRET_ACCESS_KEY")

	// Initialize minio client object.
	minioclient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalln(err)
	}

	MinioClient = minioclient

	log.Printf("%#v\n", minioclient)
}
