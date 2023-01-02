package utils

import (
	"e-learning/entity"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() error {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv("DATABASE_URL"),
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	err = conn.AutoMigrate(entity.User{}, entity.Task{}, entity.Question{}, entity.Admin{})
	if err != nil {
		return err
	}
	SetupDBConnection(conn)

	return nil
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}
